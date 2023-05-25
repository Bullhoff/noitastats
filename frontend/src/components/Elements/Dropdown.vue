<template>
	<div :ref="(el) => (refs.container = el)" :style="[style.container, state.open ? style.containerOpen : {width: 'fit-content'}]">
		<div
			class="prevent-select"
			@click="state.open = !state.open"
			:style="[style.head, state.open ? {} : {}, state.hover ? style.headHover : {}]"
			@mouseenter="state.hover = true"
			@mouseleave="state.hover = false"
			@blur="onBlur({e: $event, path: props.path})"
			@focus="onFocus({e: $event, path: props.path})"
		>
			<span :style="[style.arrow]" @click="onClick" :ref="(el) => (refs.arrow = el)" v-if="config?.arrowLeft">
				{{ String.fromCodePoint('0x1F87B') }}
			</span>
			<slot name="head" style="" @click="onClick" />
			<span :style="[style.arrow]" @click="onClick" :ref="(el) => (refs.arrow = el)" v-if="config?.arrowRight">
				{{ String.fromCodePoint('0x1F87B') }}
			</span>
		</div>
		<div @blur="onBlur({e: $event, path: props.path})" @focus="onFocus({e: $event, path: props.path})" :ref="(el) => (refs.content = el)" tabindex="0" v-if="state.open" :style="[style.contentContainer]">
			<div :style="[style.content, [props.editable ? {} : {}]]">
				<slot :path="[...props.path, 'aa']" />
			</div>
		</div>
	</div>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import * as store from '../../stores/store.js';
import _ from 'lodash';

const emit = defineEmits(['click']);

const props = defineProps(
	reactive({
		selected: {default: ''},
		editable: {default: false},
		obj: {
			type: Object,
			required: false,
		},
		path: {
			type: Array,
			required: false,
			default: ['root'],
		},
		style: {
			type: Object,
			required: false,
			default: [],
		},
		config: {
			type: Object,
			required: false,
			default: {},
		},
	})
);

const styleDefault = {
	arrow: {
		'justify-self': 'end',
		cursor: 'pointer',
		'margin-right': 0,
		right: 0,
	},
	container: {
		position: 'static',
		'box-sizing': 'border-box',
		display: 'flex',
		'flex-direction': 'column',
		'flex-wrap': 'none',
		'background-color': 'rgb(10,10,55)',
		'border-radius': '3px',
		'line-height': '1',
		width: 'fit-content',
	},
	containerOpen: {
		'box-sizing': 'border-box',
	},
	head: {
		'box-sizing': 'border-box',
		position: 'relative',
		'background-color': 'rgb(50,50,55)',
		border: '2px solid rgb(110,110,115)',
		'border-radius': '3px',

		cursor: 'pointer',
		display: 'flex',
		'flex-direction': 'row',
		'align-items': 'center',
		'align-content': 'center',
		width: 'fit-content',
	},
	headHover: {
		'background-color': 'rgb(150,150,155)',
	},
	contentContainer: {
		'box-sizing': 'border-box',
	},
	content: {
		position: 'absolute',
		'background-color': 'rgb(0, 26, 51)',
		'min-width': '20%',
		'min-height': '50%',
		'z-index': 1,
		'box-sizing': 'border-box',
		'white-space': 'nowrap',
		overflow: 'auto',
		'overflow-x': 'hidden',
		resize: 'both',
	},
};
const configDefaults = {
	arrowLeft: false,
	arrowRight: false,
};
const config = reactive(_.merge(configDefaults, props.config));
const p = reactive(props);
const refs = reactive({
	container: null,
	content: null,
	columnHandle: [],
	column: [],
});

const style = computed(() => {
	let obj = _.merge(styleDefault, props.style);
	let rect = refs.container ? refs.container.getBoundingClientRect() : null;
	if (rect) obj.content['max-height'] = `${window.innerHeight - rect.top - 50}px`;
	console.log('**style**', obj.content['max-height'], window.innerHeight, refs.container?.offsetTop, rect?.top);
	return obj;
});
var state = reactive({
	open: false,
});

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});

watch(
	() => refs.content,
	(val) => {
		if (refs.content) refs.content.focus();
	}
);

const treeObj = reactive({});
const stateObj = reactive({});

async function onBlur(obj) {
	let {e, element, fromSub, subId, path} = obj;
	await setTimeout(() => {
		treeObj[path.join('__')] = false;
		let open = false;
		for (const [key, value] of Object.entries(treeObj)) {
			if (value == true) open = true;
		}
		state.open = open;
	}, 10);
}
function onFocus(obj) {
	// , path = p.path
	let {e, element, fromSub, subId, path} = obj;
	treeObj[path.join('__')] = true;
}
provide('injected', {
	path: props.path,
	treeObj: treeObj,
	onBlur: onBlur,
	onFocus: onFocus,
});

provide('stateObj', stateObj);
provide('onBlur', onBlur);
provide('onFocus', onFocus);

function onClick(e, field) {
	if (field == 'selected' && props.editable) {
	} else emit('click', e, refs.arrow);
}
</script>

<style scoped>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}
</style>
