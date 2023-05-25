<template>
	<table cellspacing="0" BORDER="0" :ref="(el) => (refs.table = el)" :style="[state.holding ? {cursor: 'col-resize'} : {}]">
		<tr style="visibility: hidden">
			<td :ref="(el) => (refs.leftColumn = el)"></td>
			<td></td>
		</tr>
		<tr v-for="(value, key) in obj">
			<td>
				<span class="key">
					{{ key }}
				</span>
				<div class="prevent-select resize-handle" @mousedown="resizeMouseDown($event)"></div>
			</td>
			<td :ref="(el) => (refs.rightColumn = el)" :title="'\ntypeof: ' + typeof value + '\ntype: ' + utils.type(value) + '\n'">
				<span v-if="utils.type(value) == 'bool'">
					<Checkbox :prop="{...prop, path: [...prop.path, key]}" :checked="value" :style="{container: {}}" @onChange="onChange" />
				</span>
				<span v-else-if="!editableValues" class="value">
					{{ value != '' ? value : '&nbsp;' }}
				</span>

				<span v-else-if="true" :contenteditable="!Array.isArray(editableValues) ? editableValues : listContainsValue(editableValues, key)">
					{{ value }}
				</span>
				<span v-else>
					<textarea
						class="value-editable"
						v-model="obj[key]"
						rows="1"
						spellcheck="”false”"
						data-gramm="false"
						:readonly="!Array.isArray(editableValues) ? !editableValues : !listContainsValue(editableValues, key)"
						:style="[style?.textarea, style?.textareaValue]"
						@change="emitHandler(emitObj($event, key, value, 'change', value));"
					></textarea>
				</span>
			</td>
		</tr>
	</table>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed} from 'vue';
import {storeToRefs} from 'pinia';
//import {stateStore, storeStore, imageStore, contextMenuStore} from "../../stores/store.js";
//import TreeView from "./TreeView.vue";
//import Elements from "./Elements.vue";

const emit = defineEmits(['emit-handler', 'onChange']);

const props = defineProps(
	reactive({
		obj: {default: {}},
		containerRef: {default: null},
		elements: {default: []},
		nested: {default: true},
		objtitle: {default: false},
		editableKeys: {default: false},
		editableValues: {default: false},
		treeState: {
			type: Object,
			required: false,
			default: {expanded: false},
		},
		prop: {
			type: Object,
			required: false,
			default: {
				state: {},
				path: [],
			},
		},
		obj: {
			type: Object,
			required: false,
		},
		path: {
			type: Array,
			required: false,
			default: [],
		},
		style: {
			type: Object,
			required: false,
			default: [],
		},
	})
);
const refs = reactive({});
const styleDefault = {
	table: {
		cellspacing: '30',
		cellpadding: '30',
	},
	container: {
		'max-width': '100%',
		position: 'absolute',
	},
	key: {
		'white-space': 'wrap',
		'padding-right': '1ch',
		border: '1px solid rgb(79, 79, 79)',
		width: 'fit-content',
		'min-width': 'fit-content',
	},
	textareaKey: {
		color: 'rgb(100,100,250)',
		resize: 'none',
		'font-weight': 'bold',
		width: 'fit-content',
		'white-space': 'wrap',
		resize: 'vertical',
		'min-width': 'fit-content',
	},
	value: {
		border: '1px solid rgb(79, 79, 79)',
		'max-width': '*',
		width: '*',
		'white-space': 'wrap',
	},
	textareaValue: {
		color: 'rgb(200,200,100)',
		resize: 'vertical',
	},
	textarea: {
		'font-size': '16px',
		overflow: 'hidden',
		'vertical-align': 'middle',
		'line-height': '15.999px',
		'word-break': 'break-all',
		width: '100%',
		'background-color': 'transparent',
		border: 'none',
		outline: 'none',
		'-webkit-box-shadow': 'none',
		'-moz-box-shadow': 'none',
		'box-shadow': 'none',
		resize: 'none',
	},
};

var state = reactive({
	x1: null,
	y1: null,
	x2: null,
	y2: null,
	holding: false,
});

function coords(e) {
	let {clientX, clientY} = e;
	let nrtwo = () => {
		state.x2 = clientX;
		state.y2 = clientY;
	};
	state.x1 = state.x2 - clientX;
	state.y1 = state.y2 - clientY;
	nrtwo();
	let x = state.x1;
	let y = state.y1;
	return {x: x, y: y};
}
function onChange(obj) {
	emit('onChange', obj);
}

function resizeMouseDown(e) {
	coords(e);
	state.holding = true;
	window.addEventListener('mousemove', resizeMouseMove);
	window.addEventListener('mouseup', resizeMouseUp);
	window.addEventListener('mouseleave', resizeMouseUp);
}

function resizeMouseUp(e) {
	state.holding = false;
	window.removeEventListener('mousemove', resizeMouseMove);
	window.removeEventListener('mouseup', resizeMouseUp);
	window.removeEventListener('mouseleave', resizeMouseUp);
}
function resizeMouseMove(e) {
	if (!state.holding) return;
	let {x, y} = coords(e);
	let rect = refs.leftColumn.getBoundingClientRect();
	var newWidth = e.clientX - rect.left;
	refs.table.style.gridTemplateColumns = `${newWidth}px auto`;
}

function listContainsValue(arr, value) {
	for (let i = 0; i < arr.length; i++) {
		if (arr[i] == value || arr[i] == '*') return true;
	}
	return false;
}

function rectLeft(r) {
	let rect = refs.values?.getBoundingClientRect();
	return rect?.left;
}
const style = computed(() => {
	let obj = {...styleDefault, ...props.style};
	if (obj.value['width'] == '*' && rectLeft(refs.values)) obj.value['width'] = window.innerWidth - (rectLeft(refs.values) + 20) + 'px';
	return obj;
});
function emitObj(obj) {
	return {...obj, path: [...props.path, obj.key]};
}
function emitHandler(obj) {
	emit('emit-handler', obj);
}
const treeState = reactive(props.treeState);

onBeforeMount(() => {
	if (!props.obj) return;
	for (const [key, value] of Object.entries(props.obj)) {
		try {
			props.obj[key] = JSON.parse(value);
		} catch (e) {}
		if (typeof value == 'object') {
			treeStateDefault(key);
		}
	}
});

onMounted(() => {});
onUpdated(() => {});

function treeStateDefault(key) {
	if (!treeState[key]) treeState[key] = {expanded: false};
}
</script>

<style scoped>
* {
	margin: 0;
	padding: 0;
}
table {
	display: grid;
	grid-template-columns: fit-content(1000px) minmax(150px, 1fr);
	overflow: auto;
	width: 100%;
	padding-right: 5px;
	border-collapse: collapse;
	border-spacing: 0;
	line-height: 1;
}

table thead,
table tbody,
table tr {
	display: contents;
}

table th {
	position: relative;
	font-size: 18px;
}

td,
td div,
td span,
table th,
table td {
	position: relative;
	text-align: left;
	white-space: nowrap;
	overflow: hidden;
	width: 100%;
}

td div,
td span {
	padding: 0 3px 0 3px;
	display: block;
	position: relative;
	border-spacing: 0;
	width: 100%;
	text-overflow: ellipsis;
}
td span {
	border-spacing: 0;
	box-sizing: border-box;
	border: solid #5b6dcd 1px;
}

table thead,
table tbody,
table tr {
	display: contents;
}
.resize-handle {
	z-index: 1099999;
	position: absolute;
	cursor: col-resize;
	width: 6px;
	right: 0;
	top: 0;
	bottom: 0;
	box-sizing: border-box;
}

/* .container{
	position:relative;
	display:grid;
	grid-template-columns: minmax(150px, 3fr) minmax(150px, 3fr);
}
.col{
	position:relative;
	border: 1px solid red;
	width:100px;
} */
</style>
