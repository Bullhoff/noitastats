<template>
	<div style="display: flex; flex-direction: row">
		<Clickable
			v-for="(value, key, index) in prop.value.array"
			@click="
				value.onClick($event);
				onClick($event, value);
			"
			:title="value.tooltip"
		>
			{{ value.text }}
		</Clickable>
	</div>
</template>

<script setup>
import _ from 'lodash';
//import * as utils from './../utils.js'
import Clickable from './Clickable.vue';
import {ref, reactive, toRefs, watch, watchEffect, onMounted, onBeforeMount, onUpdated, nextTick, onDeactivated, onRenderTriggered, computed} from 'vue';

const emit = defineEmits(['click']);
const props = defineProps(
	reactive({
		prop: {
			type: Object,
			default: {key: null, value: null, style: {}, type: null},
		},
	})
);

const refs = reactive({});
const containerRef = ref(null);
const treeObj = reactive({});
const p = reactive(Object.assign({}, props));

function onClick(e, obj) {
	if (obj.close != false && obj.closeOnClick != false) emit('click', true);
}

onBeforeMount(async () => {});
onMounted(async () => {});
onUpdated(async () => {});
</script>

<style scoped></style>
