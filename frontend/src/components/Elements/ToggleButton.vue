<script setup>
import {reactive, onMounted, onUpdated, nextTick, computed} from 'vue';
//import _ from 'lodash';

const emit = defineEmits(['click']);
const props = defineProps(
	reactive({
		value: {
			type: Boolean,
			required: false,
			default: false,
		},
		text: {
			type: String,
			required: false,
			default: '',
		},
	})
);
const refs = reactive({});
const state = reactive({});

onMounted(() => {});
onUpdated(() => {});

const style = reactive({
	circle: {
		default: {
			fill: 'gray',
		},
		hover: {},
		active: {
			fill: 'green',
		},
	},
});
</script>

<template>
	<div class="container prevent-select">
		<svg class="container" viewBox="0 0 100 100" @click="emit('click')">
			<circle cx="50%" cy="50%" r="30%" :class="[value ? 'circle-active' : '']" />
			<text x="50%" y="50%" text-anchor="middle" dominant-baseline="central">{{ props.text }}</text>
		</svg>
	</div>
</template>

<style scoped>
.container {
	width: 100%;
	height: 100%;
}
circle {
	fill: gray;
}
svg:hover circle {
	stroke: rgb(0, 137, 250);
	stroke-width: 5;
}
text {
	font: bold 30px sans-serif;
}
.circle-active {
	fill: rgb(0, 255, 68);
}
.active {
	fill: rgb(0, 255, 55);
}
</style>
