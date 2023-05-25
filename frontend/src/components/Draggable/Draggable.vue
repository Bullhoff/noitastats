<template>
	<div v-show="!props.windowState?.minimized" style="">
		<div id="draggable-container" @scroll="func.onScroll" style="" @mousedown.passive="mouseDownContainer(window_id)" :style="[style.container, windowState.style]" :ref="(el) => (divRef = el)" :class="[currentState.activeWindow == window_id ? 'window-active' : '', 'window']">
			<div id="draggable-header" :style="[style.header]" @mousedown.passive="mouseDownTitle($event, window_id)" style="min-width: 1ch; width: 100%" :ref="(el) => (refs['header'] = el)">
				<header @dblclick="func.maximize()" class="prevent-select" style="position: relative; white-space: nowrap; width: 100%; height: fit-content; display: flex; flex-direction: row">
					<slot name="header" style="width: fit-content"></slot>
					<span style="position: absolute; width: fit-content; right: 0px; top: 0px">
						<Button
							:text="String.fromCodePoint('0x1F5D5')"
							@click="
								windowState.minimized = !windowState.minimized;
								func.minimize();
							"
							style="padding: 0px"
						/>
						<Button :text="props.windowState.maximized ? String.fromCodePoint('0x1F5D7') : String.fromCodePoint('0x1F5D6')" @click="func.maximize()" style="padding: 0px" />
						<Button :text="String.fromCodePoint('0x1F5D9')" @click="func.close()" style="padding: 0px" />
					</span>
				</header>
			</div>

			<div id="draggable-body" :style="style.body" style="position: relative">
				<slot name="body" style="width: fit-content"></slot>
				<slot />
			</div>

			<div class="prevent-select" style="position: absolute; width: 0.2ch; height: 100%; left: -0.4ch; top: 0; cursor: ew-resize" @mousedown="mouseDownResize($event, 'w')"></div>
			<div class="prevent-select" style="position: absolute; width: 0.2ch; height: 100%; right: -0.4ch; top: 0; cursor: ew-resize" @mousedown="mouseDownResize($event, 'e')"></div>

			<div class="prevent-select" style="position: absolute; width: 100%; height: 0.2ch; top: -0.4ch; cursor: ns-resize" @mousedown="mouseDownResize($event, 'n')"></div>
			<div class="prevent-select" style="position: absolute; width: 100%; height: 0.2ch; bottom: -0.4ch; cursor: ns-resize" @mousedown="mouseDownResize($event, 's')"></div>

			<div class="prevent-select" style="position: absolute; width: 1ch; height: 1ch; left: -0.4ch; top: -0.4ch; cursor: nwse-resize" @mousedown="mouseDownResize($event, 'nw')"></div>
			<div class="prevent-select" style="position: absolute; width: 1ch; height: 1ch; right: -0.4ch; bottom: -0.4ch; cursor: nwse-resize" @mousedown="mouseDownResize($event, 'se')"></div>

			<div class="prevent-select" style="position: absolute; width: 1ch; height: 1ch; top: -0.4ch; right: -0.4ch; cursor: nesw-resize" @mousedown="mouseDownResize($event, 'ne')"></div>
			<div class="prevent-select" style="position: absolute; width: 1ch; height: 1ch; left: -0.4ch; bottom: -0.4ch; cursor: nesw-resize" @mousedown="mouseDownResize($event, 'sw')"></div>
		</div>
	</div>
</template>

<script setup>
import {defineComponent, onMounted, onBeforeMount, onUpdated, nextTick, ref, reactive, toRef, isRef, computed} from 'vue';
//import * as utils from '@/utils.js'
import _ from 'lodash';

const emit = defineEmits(['reset-zoom', 'test-emit', 'close', 'onHolding', 'onRelease', 'minimize', 'maximize']);

const props = defineProps({
	window_id: {
		required: false,
		default: null,
	},
	/* divRef: {
		required: false,
		default: null,
	}, */
	propStyle: {
		type: Object,
		required: false,
	},
	windowState: {
		type: Object,
		required: false,
	},
	currentState: {
		type: Object,
		required: false,
	},
	functions: {
		type: Object,
		required: false,
	},
});
const refs = reactive({});
const divRef = ref(null);
//const divRef = ref(props.divRef)
const defaultWindowState = {divRef: null, style: {}, holding: false, posX: 0, posY: 0, name: '', open: false, minimized: false, maximized: false};
const windowState = reactive(_.merge(defaultWindowState, props.windowState));

const defaultCurrentState = {activeWindow: null, x1: 0, x2: 0, y1: 0, y2: 0};

const currentState = reactive(_.merge(defaultCurrentState, computed(() => props.currentState).value));

const defaultStyle = {
	container: {
		display: 'flex',
		flexDirection: 'column',
		maxWidth: '100vw',
		maxHeight: '100vh',
		borderRadius: '5px',
		position: 'fixed',
		border: '1px solid #d3d3d3',
		color: '#fff',
		'z-index': '99',
	},
	header: {
		position: 'sticky',
		height: '2.5ch',
		cursor: 'move',
		Zindex: 10,
		backgroundColor: '#1f1b7e',
		color: '#fff',
		borderRadius: '5px',
	},
	body: {
		position: 'relative',
		width: '100%',
		height: '100%',
		alignItems: 'left',
		alignContent: 'left',
		overflow: 'auto',
		maxWidth: '95vw',
		maxHeight: '95vh',
	},
};

const style = reactive(_.merge(defaultStyle, props.propStyle));

function mouseDownResize(e, side) {
	func.mouseDown(e, props.window_id, side);

	window.addEventListener('mousemove', mouseMove);
	window.addEventListener('mouseup', mouseUp);
	window.addEventListener('mouseleave', mouseUp);
}

function mouseDownContainer() {
	func.mouseDownContainer(props.window_id);
}
function mouseDownTitle(e) {
	func.mouseDownTitle(e, props.window_id);
	window.addEventListener('mousemove', mouseMove);
	window.addEventListener('mouseup', mouseUp);
	window.addEventListener('mouseleave', mouseUp);
}
function mouseUp() {
	func.mouseUp(props.window_id);
	window.removeEventListener('mousemove', mouseMove);
	window.removeEventListener('mouseup', mouseUp);
	window.removeEventListener('mouseleave', mouseUp);
}
function mouseMove(e) {
	func.mouseMove(e);
}

const defaultFunctions = {
	onScroll(e) {
		if (divRef.value.scrollTop > refs.header.offsetHeight * 2) {
			style.header.position = 'fixed';
			style.header.width = 'inherit';
		} else if (divRef.value.scrollTop < refs.header.offsetHeight) {
			style.header.position = 'sticky';
			style.header.width = '100%';
		}
	},
};
const func = reactive(_.merge(defaultFunctions, props.functions));

onBeforeMount(() => {});
onMounted(async () => {
	await nextTick();
	var nIntervId = setInterval(() => {
		console.log('nIntervId', nIntervId, divRef.value, props.window_id);
		if (func?.onMounted) {
			func.onMounted(divRef.value, props.window_id);
			clearInterval(nIntervId);
		}
	}, 10);
});
</script>

<style scoped>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}

.window {
	background-color: #3b3838;
}
.window-active {
	background-color: #595454;
}

.prevent-select {
	-webkit-user-select: none; /* Safari */
	-ms-user-select: none; /* IE 10 and IE 11 */
	user-select: none; /* Standard syntax */
}
</style>
