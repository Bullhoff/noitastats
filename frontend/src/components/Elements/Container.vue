<template>
	<div :style="[style.container, state.open ? style.containerOpen : {width: 'fit-content'}]">
		<div
			class="prevent-select"
			@click="state.open = !state.open"
			:style="[
				style.head,
				{
					cursor: 'pointer',
					display: 'flex',
					'flex-direction': 'row',
					'align-items': 'center',
					'align-content': 'center',
					width: 'fit-content',
				},
				state.open ? {} : {},
				state.hover ? {'background-color': 'yellow'} : {},
			]"
			@mouseenter="state.hover = true"
			@mouseleave="state.hover = false"
		>
			<slot name="head" style="" @click="onClick" />
		</div>
		<div v-if="state.open" :style="[style.content, [props.editable ? {} : {}]]">
			<slot />
		</div>
	</div>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../../stores/store.js';

const emit = defineEmits(['emit-handler', 'onChange', 'click']);

const props = defineProps(
	reactive({
		selected: {default: ''},
		editable: {default: false},
		containerRef: {default: null},
		th: {default: null},
		td: {default: {}},
		obj: {default: {}},
		elements: {default: []},
		nested: {default: true},
		objtitle: {default: false},
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
		table: {
			type: Object,
			required: false,
			default: {thVisible: false, th: null, td: null},
		},
	})
);
const p = reactive(props);
const refs = reactive({
	columnHandle: [],
	column: [],
});
const styleDefault = {
	container: {
		position: 'static',
		'box-sizing': 'border-box',
		display: 'flex',
		'flex-direction': 'column',
		'flex-wrap': 'none',
		'background-color': 'rgb(10,10,55)',
		border: '2px solid rgb(110,110,115)',
		'border-radius': '3px',
		'line-height': '1',
		width: 'fit-content',
	},
	containerOpen: {
		width: '100%',
		'box-sizing': 'border-box',
	},
	head: {
		position: 'relative',
		'background-color': 'rgb(50,50,55)',
		border: '2px solid rgb(110,110,115)',
		'border-radius': '3px',
	},
	content: {
		'box-sizing': 'border-box',
		overflow: 'auto',
		'white-space': 'nowrap',
		width: '100%',
	},
};

var state = reactive({
	open: false,
});

const style = computed(() => {
	let obj = {...styleDefault, ...props.style};
	return obj;
});

function emitHandler(obj) {
	emit('emit-handler', obj);
}

function onClick(e, field) {
	if (field == 'selected' && props.editable) {
	} else emit('click', e, refs.arrow);
}

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});
</script>

<style scoped>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}
</style>
