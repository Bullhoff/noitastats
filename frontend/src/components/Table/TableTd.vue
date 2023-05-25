<template style="">
	<!-- v-if="props.obj" -->
	<template v-if="true">
		<td :title="columnKey + ': ' + rowObject[columnKey] + '\n'" v-if="colState[columnKey].show" :style="[style.td]">
			<span :style="style.item" @click="onClick({e: $event, columnClicked: columnKey, valueClicked: rowObject[columnKey], ...rowObject})">
				{{ rowObject[columnKey] }}
			</span>
			<ResizeHandle :columnKey="columnKey" />
			<!-- <div class="prevent-select resize-handle" :style="style.resizeHandle" @dblclick="dblclick($event, columnKey)" @mousedown="resizeMouseDown($event, columnKey)" ></div> -->
		</td>
	</template>
</template>

<script setup>
//import SocketioService from './../socketio.service.js';
//import { canvasStore, styleStore, configStore } from "@/stores/store.js";
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import {storeToRefs} from 'pinia';
//import {stateStore, storeStore, imageStore, contextMenuStore} from "../../stores/store.js";
import * as store from '../../stores/store.js';

import _ from 'lodash';
import TableTd from './TableTd.vue';
import ResizeHandle from './ResizeHandle.vue';

const emit = defineEmits(['emit-handler', 'onChange', 'click']);

const props = defineProps(
	reactive({
		rowObject: {default: {}},
		columnKey: {
			default: '',
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
	})
);

const p = reactive(props);
var state = reactive({});
const sortBy = reactive({});
var stateObj = inject('stateObj', {});
var onBlur = inject('onBlur', () => {});
var onFocus = inject('onFocus', () => {});

//var {table, uniqueColumnValues, colState, refs} = reactive(props)
var {onClick, rows, style, refs, colState, table, uniqueColumnValues, sortArguments, sortObject} = inject('tableInject', {
	onClick: () => {},
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
});
var injected = inject('injected', {
	open: true,
	path: [],
	treeObj: {},
	onBlur: () => {},
	onFocus: () => {},
});

function onContextMenu(e) {
	let items = reactive({});
	store.contextMenuStore().onContextMenu(e, items, null, injected);
}

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});
//watch(()=>props.obj, (obj) => {})
</script>

<style scoped>
.hideFullColumn tr > .hidecol {
	display: none;
}
</style>
