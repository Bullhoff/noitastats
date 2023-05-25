<script setup>
import {reactive, onMounted, onUnmounted, onUpdated, nextTick, computed} from 'vue';
import * as app from '../wailsjs/go/main/App';
import * as runtime from '../wailsjs/runtime/runtime.js';
import {storeToRefs} from 'pinia';
import * as store from './stores/store.js';
import * as parse from './scripts/parse.js';
import Noita from './components/Noita.vue';

onMounted(() => {
	store.storeStore().addMsg({event: 'onMounted', text: 'yep', json: null});
	console.log('onMounted', window.wails);
	runtime.EventsOn('server-event', (e) => {
		if (!e.AvailabeFilesSingle) console.log('server-event', e);
		if (e == 'exiting' && store.storeStore().state.configLoaded) {
			// store.storeStore().state.finisedLoading
			runtime.EventsEmit('client', {
				config: store.storeStore().config,
				session_list: store.storeStore().config?.session_list,
				highlight: store.storeStore().highlight,
			});
		} else {
			store.storeStore().addMsg({event: 'server-event', text: JSON.stringify(Object.keys(e)), json: e});
			store.dataStore().updateData(e);
		}
	});
	runtime.EventsEmit('client-event', 'init');
	parse.parseLua.luaInit();
	store.refStore().setContentPosition();
});
onUnmounted(() => {});
</script>

<template>
	<span :style="{cursor: store.storeStore().cursor}" >
    
		<div v-if="!store.storeStore().state.configLoaded" style="width: 100%; height: 100%"> 
      Loading
    </div>

		<template v-else>
			<div v-if="store.contextMenuStore().open" style="position: fixed; z-index: 999">
				<ContextMenu :items="store.contextMenuStore().items" :init="store.contextMenuStore().pos" :config="store.contextMenuStore().config" />
			</div>

			<div style="position: absolute; width: 100%; height: 100%; z-index: 99; overflow: hidden" :style="[store.windowStore().holding ? {pointerEvents: 'auto'} : {pointerEvents: 'none'}]">
				<template v-for="(item, index) in store.windowStore().windows" :key="item.id">
					<Draggable style="pointer-events: auto" :functions="store.windowStore().func" :window_id="item.id" :windowState="store.windowStore().windows[item.id]">
						<template v-slot:header>
							<p v-if="store.windowStore()?.props?.[item?.id]?.head?.text">{{ store.windowStore()?.props?.[item?.id]?.head?.text }}</p>
							<p v-else>{{ item?.id }}</p>
						</template>
						<template v-slot:body>
							<component
								:is="item.component"
								@click="store.windowStore()?.props?.[item?.id]?.body?.click"
								:props="store.windowStore()?.props?.[item?.id]"
								:prop="store.windowStore()?.props?.[item?.id]"
								:style="store.windowStore()?.props?.[item?.id]?.body?.style"
								:obj="store.windowStore()?.props?.[item?.id]?.body?.obj"
								:text="store.windowStore()?.props?.[item?.id]?.body.text"
								:treeState="{}"
								:containerRef="store.windowStore().refs[item.id]"
							></component>
						</template>
					</Draggable>
				</template>
			</div>
			<Noita />
		</template>
	</span>
</template>

<style>
.wait {
	cursor: wait !important;
}
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}
button {
	background-color: #4b4b4b;
}

.prevent-select {
	-webkit-user-select: none; /* Safari */
	-ms-user-select: none; /* IE 10 and IE 11 */
	user-select: none; /* Standard syntax */
}
.enable-select {
	-webkit-user-select: auto; /* Safari */
	-ms-user-select: auto; /* IE 10 and IE 11 */
	user-select: auto; /* Standard syntax */
}

::-webkit-scrollbar {
	width: 10px;
}
::-webkit-scrollbar-track {
	background: #3d3d3d;
}
::-webkit-scrollbar-thumb {
	background: #888;
}
::-webkit-scrollbar-thumb:hover {
	background: #555;
}
</style>
