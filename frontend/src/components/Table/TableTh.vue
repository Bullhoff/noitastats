<template>
	<th
		v-if="colState[columnKey]?.show && !colState[columnKey]?.splitted && !colState[columnKey]?.unwrap"
		:ref="(el) => (refs.columnKeys[columnKey] = el)"
		:title="'Column: ' + columnKey"
		:style="style.th"
		style="white-space: nowrap; overflow: hidden"
		@mouseenter="colState[columnKey].hover = true"
		@mouseleave="colState[columnKey].hover = false"
	>
		<span v-if="colState[columnKey]?.hover" @click="sort(columnKey)" class="prevent-select" style="cursor: pointer; padding: 0 2px 0 0px" :style="[table.sortBy.key == columnKey ? {color: 'yellow'} : {color: 'gray'}]">
			{{ table.sortBy.key != columnKey ? String.fromCodePoint('0x294D') : table.sortBy.ascending ? String.fromCodePoint('0x1F803') : String.fromCodePoint('0x1F811') }}
		</span>
		<span
			v-if="colState[columnKey]?.hover && table.showFilterButton"
			style="cursor: pointer; padding: 0 2px 0 0px; width: 10px"
			@click="filterDropDown($event, columnKey)"
			:ref="(el) => (refs.filter = el)"
			@focus="injected.onFocus({e: $event, path: [...injected.path, ...props.path, columnKey + '__filterButton']})"
			@blur="injected.onBlur({e: $event, path: [...injected.path, ...props.path, columnKey + '__filterButton']})"
		>
			<svg viewBox="0 0 80 90" focusable="false"><path fill="aquamarine" d="m 0,0 30,45 0,30 10,15 0,-45 30,-45 Z"></path></svg>
		</span>

		<!-- <span 
							v-if="colState[columnKey]?.type == 'object' || colState[columnKey]?.type == 'array'"
							style="cursor:pointer; padding: 0 2px 0 0px; width:10px; "  
							@click="toggleObjectWrap(columnKey, obj, colState)" 
							:ref="(el)=>refs.filter=el"
							@focus="injected.onFocus({e:$event, path:[...injected.path, ...props.path, columnKey+'__filterButton']})"
							@blur="injected.onBlur({e:$event, path:[...injected.path, ...props.path, columnKey+'__filterButton'] })"
						>
						{{ String.fromCodePoint('0x1F5F2') }}
						</span> -->

		<div @click="sort(columnKey)" style="cursor: pointer; width: 100%" :style="[style.item, {'text-overflow': 'ellipsis', overflow: 'hidden'}]">
			{{ columnKey }}
		</div>
		<ResizeHandle :columnKey="columnKey" />
	</th>
	<!-- <template v-if="colState[columnKey]?.unwrap && (colState[columnKey]?.type == 'object' || colState[columnKey]?.type == 'array')">
						<template v-for="(value, key) in colState[columnKey].sub" :key="JSON.stringify(value)">
							<TableTh :style="style" :path="[...path, columnKey, key]" :columnKey="key" :obj="obj" :colState="colState[columnKey].sub"/>
						</template>
					</template> -->
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../../stores/store.js';
import _ from 'lodash';
import ResizeHandle from './ResizeHandle.vue';
import TableTh from './TableTh.vue';

const emit = defineEmits(['emit-handler', 'onChange', 'click']);

const props = defineProps(
	reactive({
		columnKey: {default: ''},
		colState: {default: {}},
		style: {
			type: Object,
			required: false,
			default: {},
		},
		path: {
			type: Array,
			required: false,
			default: ['Table'],
		},
		containerRef: {default: null},
		prop: {
			type: Object,
			required: false,
			default: {
				state: {},
				path: [],
			},
		},

		obj: {
			type: [Boolean, Object],
			required: false,
			default: false,
		},
	})
);

const p = reactive(props);
var state = reactive({});
const sortBy = reactive({});
var stateObj = inject('stateObj', {});
var onBlur = inject('onBlur', () => {});
var onFocus = inject('onFocus', () => {});

//var {table, uniqueColumnValues, colState, refs} = reactive(props)
var {
	onClick,
	sortArguments,
	sortObject,
	getObjectLength,
	toggleObjectWrap,
	rows,
	style,
	refs,
	table,
	uniqueColumnValues,
	// colState,
} = inject('tableInject', {
	onClick: () => {},
	sortObject: () => {},
	sortArguments: () => {},
	getObjectLength: () => {},
	toggleObjectWrap: () => {},
	rows: {},
	style: {},
	refs: {},
	table: {},
	uniqueColumnValues: {},
});
var injected = inject('injected', {
	open: true,
	path: [],
	treeObj: {},
	onBlur: () => {},
	onFocus: () => {},
});

function sort(col = 'id') {
	if (table.sortBy.key != col) table.sortBy.ascending = true;
	table.sortBy.key = col;
	table.sortBy.ascending = !table.sortBy.ascending;
}
function filterDropDown(e, columnKey) {
	const check = (val = true) => {
		for (const [key, value] of Object.entries(uniqueColumnValues[columnKey])) {
			uniqueColumnValues[columnKey][key] = val;
		}
	};
	let filterInput = {
		type: 'checkbox',
		text: 'Filter',
		value: table.showFilter,
		placeholder: '',
		onChange: ({checked}) => {
			table.showFilter = checked;
		},
	};
	let buttons = {
		type: 'buttons',
		array: [
			{text: 'Check all', tooltip: 'Open config.json', onClick: () => check(true)},
			{text: 'Uncheck all', tooltip: 'Open config.json', onClick: () => check(false)},
		],
	};
	let checkboxes = {};

	for (const [key, value] of Object.entries(uniqueColumnValues[columnKey])) {
		checkboxes[key] = {
			text: key,
			type: 'checkbox',
			onChange: ({checked}) => {
				console.log('ONCHANGE', checked);
				uniqueColumnValues[columnKey][key] = checked;
			},
			value: uniqueColumnValues[columnKey][key],
		};
	}
	let items = [
		filterInput,
		buttons,
		...Object.entries(checkboxes)
			.sort((a, b) => sortArguments(a[1].text, b[1].text))
			.reduce((acc, c) => {
				acc.push(c[1]);
				return acc;
			}, []),
	];
	store.contextMenuStore().onContextMenu(e, items, null, injected);
}

function onContextMenu(e) {
	let items = reactive({});
	store.contextMenuStore().onContextMenu(e, items, null, injected);
}

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});
//watch(()=>props.obj, (obj) => {})
</script>

<style scoped></style>
