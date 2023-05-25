<template style="">
	<label :style="style.container">
		<input type="checkbox" class="checkbox" :style="style.checkbox" :checked="checked" @change="emitHandler({e: $event, _checked: checked})" />
		<svg class="svg-checkbox" viewBox="0 0 20 20">
			<path d="M3,1 L17,1 L17,1 C18.1045695,1 19,1.8954305 19,3 L19,17 L19,17 C19,18.1045695 18.1045695,19 17,19 L3,19 L3,19 C1.8954305,19 1,18.1045695 1,17 L1,3 L1,3 C1,1.8954305 1.8954305,1 3,1 Z"></path>
			<polyline points="4 11 8 15 16 6"></polyline>
		</svg>
	</label>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed, inject, provide} from 'vue';
import _ from 'lodash';

const emit = defineEmits(['emit-handler', 'onChange']);

const props = defineProps(
	reactive({
		checked: {default: true},
		buttons: {default: []},
		obj: {
			type: Object,
			required: false,
		},
		style: {
			type: Object,
			required: false,
			default: {},
		},
		prop: {
			type: Object,
			required: false,
			default: {},
		},
	})
);
const refs = reactive({});
const styleDefault = {
	container: {
		'background-color': 'blue',
		position: 'relative',
		display: 'flex',
		'align-items': 'center',
		cursor: 'pointer',
		width: '16px',
		height: '16px',
	},
	checkbox: {
		visibility: 'hidden',
	},
};
const style = computed(() => _.merge(styleDefault, props.style));
//const p = inject('p', () => {})
//const onChange = inject('event', (obj) => {emit('onChange', obj)})

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});

function emitHandler(obj) {
	let {e} = obj;
	obj = {
		...props.prop,
		...obj,
		element: 'checkbox',
		checked: e?.target?.checked,
	};
	emit('onChange', obj);
	emit('emit-handler', obj);
	return obj;
}
</script>

<style scoped>
svg {
	position: absolute;
	left: 0px;
	vertical-align: middle;
	width: 100%;
	height: 100%;
	fill: rgb(108, 108, 108);
}
svg polyline {
	visibility: hidden;
	stroke: red;
	stroke-width: 3;
}

svg:hover {
	fill: rgb(125, 127, 0);
}
svg:hover polyline {
	visibility: visible;
	stroke: rgba(247, 5, 5, 0.2);
}

input[type='checkbox']:checked + svg polyline,
input[type='checkbox']:checked + svg {
	visibility: visible;
	fill: rgb(105, 105, 105);
	stroke: rgb(234, 234, 234);
}
</style>
