<template>
	<span style="width: 100%; display: flex">
		<input type="radio" :checked="prop.value.active == prop.value.id" :name="prop.value.name" :id="prop.value.id" @click="prop.value.onChange($event, $event.target.id, $event.target)" />
		<label :for="prop.value.id" style="width: 100%">
			{{ prop.value.text }}
		</label>
	</span>
</template>

<script setup>
import _ from 'lodash';
import Clickable from './Clickable.vue';
//import * as utils from './../utils.js'

import {ref, reactive, toRefs, watch, watchEffect, onMounted, onBeforeMount, onUpdated, nextTick, onDeactivated, onRenderTriggered, computed} from 'vue';

const emit = defineEmits([]);
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

onBeforeMount(async () => {});
onMounted(async () => {});
onUpdated(async () => {});
</script>

<style scoped></style>
