<template>
	<template v-if="true">
		<tr :style="style.tr" v-for="(rowObject, rowIndex) in rows">
			<template v-for="(columnDisplay, columnKey) in colState">
				<TableTd :style="style" :path="path" :columnKey="columnKey" :rowObject="rowObject" />
			</template>
		</tr>
	</template>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../../stores/store.js';

import _ from 'lodash';
import TableTd from './TableTd.vue';

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
