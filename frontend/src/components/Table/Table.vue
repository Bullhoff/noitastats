<template>
	<div :style="style.container">
		<template v-if="props.obj">
			<table cellspacing="0" BORDER="0" :ref="(el) => (refs.table = el)" :style="style.table">
				<TableHead :style="style" :path="path" :colState="colState" :obj="obj" />
				<TableBody :style="style" :path="path" :colState="colState" :obj="obj" />
			</table>
		</template>
	</div>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import {storeToRefs} from 'pinia';
//import {stateStore, storeStore, imageStore, contextMenuStore} from "../../stores/store.js";
import * as store from '../../stores/store.js';

import _ from 'lodash';
import TableHead from './TableHead.vue';
import TableBody from './TableBody.vue';

const emit = defineEmits(['emit-handler', 'onChange', 'click', 'table-state', 'obj']);

const props = defineProps(
	reactive({
		obj: {default: {}},
		containerRef: {default: null},
		prop: {
			type: Object,
			required: false,
			default: {
				state: {},
				path: [],
			},
		},
		path: {
			type: Array,
			required: false,
			default: ['Table'],
		},
		style: {
			type: Object,
			required: false,
			default: {},
		},
		obj: {
			type: [Boolean, Object],
			required: false,
			default: false,
		},
		table: {
			type: Object,
			required: false,
			default: {sortBy: {key: false, ascending: false, columns: null}},
		},
	})
);
const path = reactive(props.path);
const tableDefaults = {
	//columns: {},
	showFilterButton: true,
	showFilter: false,
	showColumns: false,
	sortBy: {key: false, ascending: false},
	showValues: false,
	excludeColumnValues: false,
	includeColumnValues: false,
	filterInput: {},
};
const table = reactive(_.merge(tableDefaults, props.table));
const p = reactive(props);
const refs = reactive({
	columnHandle: [],
	column: [],
	columns: {},
	columnKeys: {},
});
const uniqueColumnValues = reactive({});
const colState = reactive({});

const styleDefault = {
	container: {
		//'position':'unset',
		display: 'flex',
		'flex-direction': 'column',
		width: '100%',
		'white-space': 'nowrap',
		'box-sizing': 'border-box',
		//'overflow':'hidden',
	},
	table: {
		display: 'grid',
		width: '100%',
		'line-height': 1,
		'box-sizing': 'border-box',
		'border-collapse': 'collapse',
		'border-spacing': 0,
	},
	resizeHandle: {
		'z-index': 1,
		position: 'absolute',
		cursor: 'col-resize',
		width: '6px',
		right: 0,
		top: 0,
		bottom: 0,
		'box-sizing': 'border-box',
	},
	tr: {
		'box-sizing': 'border-box',
		'background-color': 'rgb(10,10,15)',
		display: 'contents',
	},
	th: {
		position: 'relative',
		width: '100%',
		height: '16px',
		'background-color': 'rgb(10,10,155)',
		'text-overflow': 'ellipsis',
		padding: '0 3px 0 3px',
		'white-space': 'nowrap',

		display: 'flex',
		'flex-direction': 'row',
		'flex-wrap': 'nowrap',

		'border-collapse': 'collapse',
		'box-sizing': 'border-box',
		'border-spacing': 0,
		border: '1px solid white',
		'line-height': 1,
	},
	td: {
		position: 'relative',
		'border-collapse': 'collapse',
		'box-sizing': 'border-box',
		'border-spacing': 0,
		border: 'solid #5B6DCD 1px',
		overflow: 'hidden',
	},
	item: {
		width: '100%',
		display: 'block',
		position: 'relative',
		'text-overflow': 'ellipsis',
		'white-space': 'nowrap',
		overflow: 'hidden',
		'border-spacing': 0,
		'box-sizing': 'border-box',
		'text-align': 'left',
		padding: '0 3px 0 3px',
	},
};
const style = computed(() => _.merge(styleDefault, props.style));

function onClick(obj) {
	emit('click', obj);
}

const rows = computed(() => {
	let list = {...p.obj};
	for (const [rowKey, rowObj] of Object.entries(list)) {
		for (const [columnKey, columnValue] of Object.entries(rowObj)) {
			if (table.columns && !table.columns[columnKey]) {
				delete list[rowKey];
			}
			if (typeof columnValue == 'object') {
				delete list[rowKey][columnKey];
				continue;
			}
			if (table?.excludeColumnValues?.[columnKey] != undefined) {
				if (typeof table.excludeColumnValues[columnKey] == 'string') table.excludeColumnValues[columnKey] = [table.excludeColumnValues[columnKey]];
				let removed = false;
				for (let i = 0; i < table.excludeColumnValues[columnKey].length; i++) {
					if (columnValue == table.excludeColumnValues[columnKey][i]) {
						if (removed) continue;
						delete list[rowKey];
						removed = true;
					}
				}
				if (removed) continue;
			}
			if (table?.filterInput?.[columnKey]) {
				if (!columnValue.toString().toLowerCase().includes(table.filterInput[columnKey].toString().toLowerCase())) {
					delete list[rowKey];
					continue;
				}
			}

			if (!uniqueColumnValues[columnKey]) uniqueColumnValues[columnKey] = {};
			if (uniqueColumnValues[columnKey][columnValue] == undefined) uniqueColumnValues[columnKey][columnValue] = true;
			if (uniqueColumnValues[columnKey]?.[columnValue]) continue;
			//delete list[rowKey][columnKey]
			delete list[rowKey];
		}
	}
	return sortObject(list, table.sortBy.key, table.sortBy.ascending);
});

function sortObject(obj, sortBy, ascending = table.sortBy.ascending) {
	return Object.entries(obj)
		.sort((a, b) => {
			return sortArguments(a[1][sortBy], b[1][sortBy], ascending);
		})
		.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
}

function sortArguments(val1, val2, ascending = table.sortBy.ascending) {
	let a = ascending ? val1 : val2;
	let b = ascending ? val2 : val1;
	if (a == '') return -1;
	if (b == '') return 1;
	if (a == null || a == undefined) return -1;
	if (b == null || b == undefined) return 1;
	if (!isFinite(a) && !isFinite(b)) {
		return isNaN(a) && isNaN(b) ? a.localeCompare(b) : a < b ? -1 : a === b ? 0 : 1;
	}
	return Number(a) - Number(b);
}

var injected = inject('injected', {
	open: true,
	path: [],
	treeObj: {},
	onBlur: () => {},
	onFocus: () => {},
});

onBeforeMount(() => {});

onMounted(() => {
	if (!props.obj) return;
	if (Object.keys(props.obj).length == 0) return;
	for (const [rowKey, rowValue] of Object.entries(props.obj)) {
		for (const [key, value] of Object.entries(rowValue)) {
			newColumn(colState, {
				key,
				columnKey: key,
				type: typeof value == 'object' ? 'object' : 'string',
				show: !table.showColumns ? true : table.showColumns.includes(key) ? true : false,
			});
		}
	}
	if (Object.keys(colState).length == 0) return;
	let rect = refs.table.getBoundingClientRect();
	let colAmount = Object.values(colState).filter((value) => value.show).length;

	let colSize = Math.round((rect.width - 1) / (colAmount + 1));
	for (const [key, value] of Object.entries(colState)) {
		if (!value || !value.show) continue;
		colState[key].width = `${colSize}px`;
	}
	colState[Object.keys(colState)[0]].width = `${colSize * 2}px`;
	refs.table.style.gridTemplateColumns = Object.values(colState)
		.filter((value) => value.show)
		.map((v) => v.width)
		.join(' ');
	for (const [rowKey, rowObj] of Object.entries(rows.value)) {
		for (const [colKey, colObj] of Object.entries(colState)) {
			if (!rowObj[colKey]) {
				p.obj[rowKey][colKey] = '';
			}
		}
	}
});
onUpdated(() => {});

watch(
	() => p.obj,
	(obj) => {
		emit('obj', p.obj);
	}
);
watch(table, (newValue, oldValue) => {
	emit('table-state', table);
});
provide('tableInject', {
	style,
	refs,
	colState,
	table,
	uniqueColumnValues,
	rows,
	sortArguments,
	sortObject,
	onClick,
});

function newColumn(colState, obj) {
	let {columnKey} = obj;
	let defaultObj = {
		sub: {},
		columnKey: columnKey,
		length: null,
		text: columnKey,
		//show: true,
		show: !table.showColumns ? true : table.showColumns.includes(colState) ? true : false,
		width: `16px`,
		widthTemp: null,
		fit: false,
		hover: false,
	};
	if (!colState[columnKey]) colState[columnKey] = _.merge(defaultObj, obj);
	else colState[columnKey] = _.merge(colState[columnKey], obj);
}
</script>

<style scoped>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}
</style>
