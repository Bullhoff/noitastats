<template style="">
	<template v-if="true">
		<tr :style="[style.tr]" @contextmenu.prevent="onContextMenu">
			<template v-for="(columnDisplay, columnKey) in colState">
				<template v-if="colState[columnKey].show && !colState[columnKey]?.splitted">
					<TableTh :style="style" :path="path" :columnKey="columnKey" :obj="obj" :colState="colState" />
				</template>
			</template>
		</tr>

		<tr :style="style.tr" v-if="table.showFilter">
			<template v-for="(columnDisplay, columnKey) in colState">
				<td v-if="colState[columnKey].show" :style="style.th">
					<input
						@focus="injected.onFocus({e: $event, element: 'p', fromSub: true, path: [...injected.path, ...props.path, columnKey + '_filter']})"
						@blur="injected.onBlur({e: $event, element: 'p', fromSub: true, path: [...injected.path, ...props.path, columnKey + '_filter']})"
						v-model="table.filterInput[columnKey]"
						@change="filterInputChange($event.target.value, columnKey)"
						style="background-color: rgba(0, 0, 0, 0); color: inherit; width: 100%"
						placeholder="filter"
					/>
				</td>
			</template>
		</tr>
	</template>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../../stores/store.js';
import _ from 'lodash';
import TableTh from './TableTh.vue';
import ResizeHandle from './ResizeHandle.vue';

const emit = defineEmits(['emit-handler', 'onChange', 'click']);

const props = defineProps(
	reactive({
		colState: {default: {}},
		obj: {
			type: [Boolean, Object],
			required: false,
			default: false,
		},

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
	})
);

const p = reactive(props);
var state = reactive({});
const sortBy = reactive({});
var stateObj = inject('stateObj', {});
var onBlur = inject('onBlur', () => {});
var onFocus = inject('onFocus', () => {});

//var {table, uniqueColumnValues, colState, refs} = reactive(props)
var {dblclick, resizeMouseDown, resizeMouseUp, resizeMouseMove, rows, style, refs, colState, table, uniqueColumnValues, sortArguments, sortObject} = inject('tableInject', {
	dblclick: () => {},
	resizeMouseDown: () => {},
	resizeMouseUp: () => {},
	resizeMouseMove: () => {},
	sortObject: () => {},
	sortArguments: () => {},
	rows: {},
	style: {},
	refs: {},
	colState: {},
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

function filterInputChange(e, columnKey) {}
function onContextMenu(e) {
	let items = reactive({});
	for (const [key, value] of Object.entries(colState)) {
		items[key] = {
			text: key,
			type: 'checkbox',
			value: colState[key].show,
			onChange: (e, e2) => {
				if (!colState[key].width) colState[key].width = 'fit-content(10ch)';
				colState[key].show = e.checked;
				refs.table.style.gridTemplateColumns = Object.entries(colState)
					.map(([key, value]) => (value.show ? value.width : ''))
					.join(' ');
			},
		};
	}
	store.contextMenuStore().onContextMenu(e, items, null, injected);
	return;
}

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});
//watch(()=>props.obj, (obj) => {})
</script>

<style scoped></style>
