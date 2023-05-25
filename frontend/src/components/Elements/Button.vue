<script setup>
import {reactive, onMounted, computed} from 'vue';

const emit = defineEmits([]);
const props = defineProps({
	onDisabled: {default: false},
	style: {
		type: Object,
		default: {},
		required: false,
	},
	text: {
		type: [String, Array],
		default: 'Text',
		required: false,
	},
	texts: {
		type: Array,
		//default: "",
		required: false,
	},
	prop: {
		type: Object,
		default: {
			text: null,
			checkbox: null,
		},
		required: false,
	},
	obj: {
		type: Object,
		default: {
			text: null,
			checkbox: null,
		},
		required: false,
	},

	elements: {
		type: Object,
		default: {
			checkbox: null,
		},
		required: false,
	},
	title: {
		type: String,
		default: '',
		required: false,
	},
	active: {
		type: Boolean,
		//default: "Text",
		required: false,
	},
	onClick: {
		type: Function,
		default: () => {},
		required: false,
	},
});
const p = computed(() => props);
const state = reactive({
	width: 0,
});
const refs = reactive({});
const activeNumber = computed(() => Number(props.active));

const currentText = computed(() => {
	if (props.onDisabled) return props.onDisabled != true ? props.onDisabled : Array.isArray(props.text) ? props.text[0] : props.text;
	if (props.text && Array.isArray(props.text)) return props.text[activeNumber.value] == undefined ? props.text[0] : props.text[activeNumber.value];
	else if (props.text) return props.text;
	else return '*';
	let text = '';
});

const defaultStyle = {
	container: {paddingLeft: '2px', paddingRight: '2px', lineHeight: '1'},
	svg: {},
	text: {},
	checkbox: {
		/* width:"15", height:"50%" */
		width: 0,
		height: 0,
	},
};
const style = reactive(Object.assign(defaultStyle, props.style));

function buttonClick(e) {
	props.onClick(e);
}

onMounted(() => {
	adaptSvg();
});

function adaptSvg() {
	if (!refs.text) return;
	style.svg.width = 6 + refs.text.getBBox().width;
	style.svg.height = refs.text.getBBox().height;
}
</script>

<template>
	<div class="container prevent-select" :title="title" :ref="(el) => (refs.container = el)" :style="style.container">
		<template v-if="!text">
			<svg class="container" viewBox="0 0 100 100" @click="emit('click')">
				<circle cx="50%" cy="50%" r="30%" :class="[value ? 'circle-active' : '']" />
				<text x="50%" y="50%" text-anchor="central" dominant-baseline="central">{{ props.text }}</text>
				<!-- middle -->
			</svg>
		</template>
		<template v-else-if="text == '*'">
			<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg" :class="[props.onDisabled ? {inactive: props.onDisabled} : {'svg-active': props.active}]" :ref="(el) => (refs.svg = el)" :style="style.svg" style="width: 2.2ch; justify-self: center; align-self: center" @click="buttonClick">
				<rect width="100" height="100" rx="30" ry="30" stroke-width="5" class="svg-rect-container" />
				<!-- fill="blue"  -->
				<line x1="20" y1="20" x2="80" y2="20" stroke="white" stroke-width="10" />
				<line x1="20" y1="50" x2="80" y2="50" stroke="white" stroke-width="10" />
				<line x1="20" y1="80" x2="80" y2="80" stroke="white" stroke-width="10" />
			</svg>
		</template>
		<template v-else>
			<svg :class="[props.onDisabled ? {inactive: props.onDisabled} : {'svg-active': props.active}]" :ref="(el) => (refs.svg = el)" :style="style.svg" style="min-width: 2ch" @click="buttonClick">
				<rect x="0" width="100%" height="100%" rx="5" stroke-width="1" class="svg-rect-container" />
				<text x="50%" y="50%" font-size="100%" text-anchor="middle" dominant-baseline="central" :ref="(el) => (refs.text = el)">
					{{ currentText ? currentText : props.prop?.text[0] }}
				</text>
			</svg>
		</template>
	</div>
</template>

<style scoped>
.container {
	cursor: pointer;
	display: inline-flex;
	flex-direction: row;

	position: relative;
	vertical-align: middle;
	width: fit-content;
	height: 100%;
}

svg {
	flex-grow: 1;
}
.svg-rect-container {
	fill: black;
	/* stroke-width: 1; */
	stroke: rgb(0, 21, 255);
}
svg text {
	fill: white;
}
svg line {
	stroke: white;
}

/* HOVER */
svg:hover .svg-rect-container {
	/* fill:rgb(57, 57, 57); */
}
svg:hover text {
	fill: rgb(249, 240, 0);
}
svg:hover line {
	stroke: rgb(249, 240, 0);
}

/* ACTIVE */
.svg-active .svg-rect-container {
	fill: rgb(56, 56, 56);
}
.svg-active text {
	fill: rgb(0, 255, 34);
}
.svg-active line {
	stroke: rgb(0, 255, 34);
}

.inactive .svg-rect-container {
	/* fill:rgb(42, 42, 42); */
}
.inactive text {
	stroke: rgb(56, 56, 56);
	fill: rgb(56, 56, 56);
}

input:checked + label > svg rect {
	fill: rgb(249, 0, 0);
}

.button {
	cursor: pointer;
	display: inline-block;
	background-color: black;
	padding: 0px 3px 0px 3px;
	margin: 0px 3px 0px 3px;
	color: white;
	border: 1px solid rgb(0, 21, 255);
	border-radius: 5px;
}
.button:hover {
	color: rgb(249, 240, 0);
}
.button-active {
	background-color: rgb(56, 56, 56);
	color: rgb(0, 255, 34);
}
</style>
