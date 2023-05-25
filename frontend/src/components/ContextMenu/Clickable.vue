<template>
	<div class="prevent-select clickable">
		<slot />
	</div>
</template>

<script setup>
import _ from 'lodash';
//import * as utils from './../utils.js'

import {ref, reactive, toRefs, watch, watchEffect, onMounted, onBeforeMount, onUpdated, nextTick, onDeactivated, onRenderTriggered, computed, inject} from 'vue';

const emit = defineEmits(['opening-sub']);
const props = defineProps(
	reactive({
		styles: {
			type: Object,
			default: {},
		},
	})
);

var defaultStyle = {
	clickable: {
		'background-color': 'rgba(75, 75, 75, 0.95)',
		display: 'flex',
		'flex-direction': 'row',
		padding: '0ch 0.5ch 0ch 0.5ch',
		color: 'white',
	},
	clickableHover: {
		'background-color': 'blue',
		'border-radius': '3px',
	},
	items: {},
	open: false,
};

var style = reactive({...defaultStyle, ...inject('style', {})});

const refs = reactive({});
const containerRef = ref(null);
const treeObj = reactive({});
const p = reactive(Object.assign({}, props));

onBeforeMount(async () => {});
onMounted(async () => {});
onUpdated(async () => {});
</script>

<style scoped>
*:focus {
	outline: none;
}
.container {
	display: block;
	position: fixed;
	z-index: 9999;
	cursor: pointer;
	margin: 0px;
	padding: 0px;
	white-space: nowrap;
	height: fit-content;
	width: fit-content;
}
.row {
	margin: 0px;
	padding: 0px;
	white-space: nowrap;
	text-align: left;
	border-radius: 15px;
}
.clickable {
	background-color: rgba(75, 75, 75, 0.95);
	display: flex;
	flex-direction: row;
	padding: 0ch 0.5ch 0ch 0.5ch;
	color: white;
	border: 1px solid rgb(0, 0, 0);
	border-radius: 3px;
}
.clickable:hover {
	background-color: blue;
	border-radius: 3px;
}
</style>
