<script setup>
import {reactive, onMounted, computed} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../../stores/store';
import * as app from '../../../wailsjs/go/main/App';
import * as runtime from '../../../wailsjs/runtime/runtime.js';

const emit = defineEmits(['emit-handler', 'click']);
const props = defineProps({
	obj: {
		type: Object,
		default: {
			id: '',
			sessionJson: {
				daily: false,
				favorite: false,
				comment: '',
			},
		},
		required: false,
	},
	text: {
		type: String,
		default: '',
		required: false,
	},
	click: {
		type: Function,
		default: () => {},
		required: false,
	},
});
const p = computed(() => props);
const state = reactive({});
const refs = reactive({
	textarea: null,
});
function buttonClick(e) {}

onMounted(() => {});

function emitHandler(e) {
	emit('emit-handler', e);
}
function onKeyDown(e) {
	//console.log('onKeyDown', e.keyCode)
	let newKey = null;
	if (e.keyCode === 9) newKey = '\u0009'; // tab
	//if (e.keyCode === 13) newKey = "\r\n"		// enter
	//if (e.keyCode === 13) newKey = "\r\n\0"	// enter
	//if (e.keyCode === 8) newKey = "\b"		// backspace

	if (newKey) {
		e.preventDefault();
		var doc = refs.textarea.ownerDocument.defaultView;
		var sel = doc.getSelection();
		var range = sel.getRangeAt(0);
		var tabNode = document.createTextNode(newKey);
		range.insertNode(tabNode);
		range.setStartAfter(tabNode);
		range.setEndAfter(tabNode);
		sel.removeAllRanges();
		sel.addRange(range);
	}
}
function save(e) {
	let resObj = {
		e: e,
		text: refs.textarea.innerHTML.replaceAll('\0', '').replaceAll('<div>', '\n').replaceAll('</div>', '').replaceAll('<br>', ''),
		obj: props.obj,
		id: props.obj?.id,
	};
	props.click(resObj);
	emit('click', resObj);
}
</script>

<template>
	<div class="container" style="">
		<div class="textbox" spellcheck="false" :ref="(el) => (refs.textarea = el)" :style="{'min-height': 16 * 3 + 'px', height: '100%'}" contenteditable="true" @keydown="onKeyDown">
			{{ props.obj?.comment }}
		</div>
		<button @click="save" style="">Save</button>
	</div>
</template>

<style scoped>
.container {
	margin: 5px;
	display: flex;
	flex-direction: column;
	gap: 5px;
	height: 95%;
}

.textbox {
	text-align: left;
	white-space: pre;
	border: 1px solid red;
	display: block;
	border: 0;
	outline: 0;
	white-space: pre;
	line-height: 1;
}
</style>
