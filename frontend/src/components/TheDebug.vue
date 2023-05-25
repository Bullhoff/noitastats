<template>
	<div>
		<div style="display: flex; flex-direction: row">
			<button @click="bool.buttons = !bool.buttons">Buttons</button>
			<button @click="bool.showMessages = !bool.showMessages">Messages</button>
		</div>
		<div v-if="bool.buttons">
			<div>
				<button @click="clickHandler(2)">2</button>
				<button @click="clickHandler(3)">3</button>
				<button @click="clickHandler(4)">4</button>
				<button @click="clickHandler(5)">5</button>
			</div>
		</div>

		<div v-if="bool.showMessages" style="display: flex; flex-direction: column">
			<template v-for="(categoryArray, msgCategory) in storeStore().msg">
				<button @click="!bool[msgCategory] ? (bool[msgCategory] = true) : (bool[msgCategory] = !bool[msgCategory])">{{ msgCategory }}</button>
				<template v-if="bool[msgCategory]">
					<div v-for="(value, i) in categoryArray" style="width: 100%" class="row-container">
						<p>
							{{ value.datetime }}
						</p>
						<p>
							{{ value.event }}
						</p>
						<p>
							{{ value.text }}
						</p>
						<span v-if="value.json" style="text-align: left; left: 0px; margin-right: auto">
							<TreeView :obj="value.json" />
						</span>
					</div>
				</template>
			</template>
		</div>

		<div v-if="bool.showValues">
			<dl>
				<template v-for="(catObj, cat) in storeStore().msg.status">
					<dt style="background-color: brown" @click="bool[cat] = !bool[cat]">{{ cat }}</dt>
					<template v-for="(value, key) in catObj" v-if="bool[cat]">
						<dd style="text-decoration: underline">
							<p style="display: inline-block; text-decoration: underline">{{ key }}</p>
							<template v-if="typeof value == 'number' || typeof value == 'string'">
								<p style="display: inline-block">: {{ typeof value == 'number' ? utils.round(value, 2) : value }}</p>
							</template>
							<template v-else-if="typeof value == 'object'">
								<p v-for="(value2, key2) in value" style="display: inline-block">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{ key2 }}: {{ typeof value == 'number' ? utils.round(value2, 2) : value2 }}</p>
							</template>
						</dd>
					</template>
				</template>
			</dl>
		</div>
		<button @click="localStorageClear">Clear localStorage</button>
	</div>
</template>

<script setup>
import * as store from './../stores/store.js';
import {defineComponent, onMounted, onBeforeMount, onUpdated, nextTick, ref, reactive, computed} from 'vue';

import * as utils from './../scripts/utils.js';
//const emit = defineEmits(['resetZoom'])

const refs = reactive([]);
let messageQueue = computed(() => {});

const bool = reactive({
	showMessages: false,
	showValues: false,
});

function localStorageClear() {
	localStorage.clear();
}
onMounted(() => {
	for (const [key, value] of Object.entries(store.storeStore().msg)) {
		//if(value.json) {}
		bool[key] = false;
		for (let i = 0; i < value.length; i++) {
			bool[key + '__' + i] = false;
		}
	}
});

function clickHandler(e) {
	console.log('clickHandler', e);
	if (e == 1) SocketioService.socket.emit('test');
	if (e == 2) SocketioService.socket.emit('stroke:remove', {canvas_id: 'yee--b0d2b197-d57a-447b-969e-cab256d298e2', layer_id: 'Hans_default--e65d2bdc-e3a2-4958-a3cf-85061c244392', strokeIndex: null});
}
</script>

<style scoped>
button {
	background-color: gray;
}
textarea {
	background-color: gray;
	caret-color: red;
	cursor: cell;
}

.row-container {
	display: flex;
	flex-direction: row;
	/* gap: 1ch; */
	border: 1px solid red;
}
.row-container p {
	padding: 0ch 1ch 0ch 1ch;
	border: 1px solid red;
}
</style>
