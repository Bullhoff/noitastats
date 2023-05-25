<template>
	<template v-if="true">
		<div class="prevent-select resize-handle" :style="style.resizeHandle" @dblclick="dblclick($event, columnKey)" @mousedown="resizeMouseDown($event, columnKey)"></div>
	</template>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../../stores/store.js';
import _ from 'lodash';

const emit = defineEmits(['emit-handler', 'onChange', 'click']);

const props = defineProps(
	reactive({
		columnKey: {default: ''},
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
var {/* dblclick, resizeMouseDown, resizeMouseUp, resizeMouseMove, */ rows, style, refs, colState, table, uniqueColumnValues, sortArguments, sortObject} = inject(
	'tableInject',
	reactive({
		//dblclick:()=>{},
		//resizeMouseDown:()=>{},
		//resizeMouseUp:()=>{},
		//resizeMouseMove:()=>{},
		sortObject: () => {},
		sortArguments: () => {},
		rows: {},
		style: {},
		refs: {},
		colState: {},
		table: {},
		uniqueColumnValues: {},
	})
);
var injected = inject('injected', {
	open: true,
	path: [],
	treeObj: {},
	onBlur: () => {},
	onFocus: () => {},
});

function dblclick(e, key) {
	let rect = refs.columnKeys[key].getBoundingClientRect();
	colState[key].fit = !colState[key].fit;
	if (colState[key].fit) {
		colState[key].widthTemp = rect.width + 'px';
		refs.table.style.gridTemplateColumns = Object.entries(colState)
			.filter(([k, v]) => v.show)
			.map(([k, v]) => {
				if (k == key) {
					//let columnRect = refs.columnKeys[key].getBoundingClientRect()
					let longest = 0;
					for (const [rowKey, rowValue] of Object.entries(rows.value)) {
						if (rowValue[key] != undefined && rowValue[key].toString().length > longest) longest = rowValue[key].toString().length;
					}
					colState[key].width = `${longest + 1}ch`;
					return colState[key].width;
				}
				return colState[key].width ? `${v.width}` : '';
			})
			.join(' ');
	} else {
		refs.table.style.gridTemplateColumns = Object.entries(colState)
			.filter(([k, v]) => v.show)
			.map(([k, v]) => {
				if (k == key && colState[key].widthTemp) {
					colState[key].width = colState[key].widthTemp;
					return colState[key].width;
				}
				return colState[key].width ? v.width : '';
			})
			.join(' ');
		colState[key].widthTemp = false;
	}
}

function resizeMouseDown(e, colId) {
	document.body.style.cursor = 'col-resize';
	coords(e);
	state.holding = true;
	window.addEventListener('mousemove', resizeMouseMove);
	window.addEventListener('mouseup', resizeMouseUp);
	window.addEventListener('mouseleave', resizeMouseUp);
	window.column = colId;
}

function resizeMouseUp(e, colId) {
	document.body.style.cursor = 'auto';
	state.holding = false;
	window.removeEventListener('mousemove', resizeMouseMove);
	window.removeEventListener('mouseup', resizeMouseUp);
	window.removeEventListener('mouseleave', resizeMouseUp);
}
function resizeMouseMove(e) {
	if (!state.holding) return;
	if (!timelag()) return;
	let colId = e.currentTarget.column;
	let gridTemplateColumns = '';
	let toTheLeft = true;
	for (const [key, value] of Object.entries(colState).filter(([key, value]) => value.show)) {
		if (!value) {
		} else if (colId == key) {
			if (colState[key].widthTemp) colState[key].widthTemp = false;
			toTheLeft = false;
			let columnRect = refs.columnKeys[key].getBoundingClientRect();
			var newWidth = Math.round(e.clientX - columnRect.left + 3);
			if (newWidth < 16) newWidth = 16;
			colState[key].width = `${newWidth}px`;
			gridTemplateColumns += ` ${colState[key].width}`;
		} else if (toTheLeft && colState[key]?.fit) {
			let columnRect = refs.columnKeys[key].getBoundingClientRect();
			var newWidth = Math.round(columnRect.width);
			if (newWidth < 16) newWidth = 16;
			colState[key].width = `${newWidth}px`;
			gridTemplateColumns += ` ${colState[key].width}`;
		} else if (true) {
			gridTemplateColumns += ` ${colState[key].width}`;
		}
		if (colState[key]?.fit) colState[key].fit = false;
	}
	refs.table.style.gridTemplateColumns = gridTemplateColumns;
}
var timelagArr = [];
const msDelay = 50;
function timelag() {
	timelagArr[1] = Date.now();
	if (timelagArr[0] && timelagArr[1] < timelagArr[0] + msDelay) return false;
	timelagArr[0] = Date.now();
	return true;
}
function coords(e, {x1 = state.x1, x2 = state.x2, y1 = state.y1, y2 = state.y2} = {}) {
	let {clientX, clientY} = e;
	let nrtwo = () => {
		state.x2 = clientX;
		state.y2 = clientY;
	};
	state.x1 = state.x2 - clientX;
	state.y1 = state.y2 - clientY;
	nrtwo();
	return {x: state.x1, y: y1};
}

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});
//watch(()=>props.obj, (obj) => {})
</script>

<style scoped>
/* .resize-handle{} */
.cursor-resize {
	cursor: col-resize;
}

.hideFullColumn tr > .hidecol {
	display: none;
}
</style>
