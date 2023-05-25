<script setup>
import {reactive, onMounted, computed, nextTick, inject, provide} from 'vue';
import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import {storeToRefs} from 'pinia';
import * as store from './../stores/store.js'; //imageStore()
import {refStore, storeStore, imageStore, contextMenuStore, dataStore} from './../stores/store.js'; //imagesStore

import * as parse from './../scripts/parse.js';
//import * as utils from './../scripts/utils.js'

import TheLifetime from './TheLifetime.vue';
import TheSessionsList from './TheSessionsList.vue';
import TheDebug from './TheDebug.vue';
import TheSession from './TheSession/TheSession.vue';
import TheAllData from './TheAllData.vue';

var components = {debug: TheDebug, alldata: TheAllData, session: TheSession, lifetime_calculator: TheLifetime}; //

onMounted(() => {
	/* storeStore().addMsg({event: 'onMounted', text: 'yep', json: null});
	console.log('onMounted', window.wails);
	runtime.EventsOn('server-event', (e) => {
		if (!e.AvailabeFilesSingle) console.log('server-event', e);
		if (e == 'exiting' && store.storeStore().state.configLoaded) {  // store.storeStore().state.finisedLoading
			runtime.EventsEmit('client', {
				config: storeStore().config,
				session_list: storeStore().config?.session_list,
				highlight: storeStore().highlight,
			});
		} else {
			storeStore().addMsg({event: 'server-event', text: JSON.stringify(Object.keys(e)), json: e});
			dataStore().updateData(e);
		}
	});
	runtime.EventsEmit('client-event', 'init');
	parse.parseLua.luaInit();
	refStore().setContentPosition(); */
});

function getNoitaData(gameid = 'latest') {
	console.log('Noita.vue --> getNoitaData --> gameid: ', gameid);
	app.GetNoitaData(gameid).then((res) => {
		console.log('Noita.vue --> getNoitaData --> GetNoitaData --> res: ', res);
		return dataStore().updateData(res);
	});
}

function onWindowsDropDown(e) {
	let items = store.windowStore().handle;
	contextMenuStore().onContextMenu(e, items, refStore().mainTopMenu.getBoundingClientRect());
}
function onDropDown(e, divRef) {
	storeStore().state.dropdown = !storeStore().state.dropdown;
	//console.log('e', e.target)
	//var style = window.getComputedStyle(e.target, null).getPropertyValue('font-size');
	//var fontSize = parseFloat(style);
	//var buttonHeight = fontSize*2.5;
	//console.log('e.target.getBoundingClientRect()', buttonHeight, fontSize, style, refStore().mainTopMenu.getBoundingClientRect())
	if (!divRef) divRef = refStore().mainTopMenu;
	let items = {
		nav: {
			type: 'buttons',
			array: [
				{text: '\u{1F3AE}', tooltip: 'Start Noita (Steam)', onClick: () => app.StartExplorer({opt: 'startnoita'}), image: imageStore().images['data/ui_gfx/loading_splash/noita_logo.png']},
				{text: '\u2699\uFE0F', tooltip: 'Open config.json', onClick: () => app.ExecCommand({opt: 'open_config'})},
				{text: '\u{27F3}', tooltip: 'WindowReloadApp', onClick: () => runtime.WindowReloadApp()},
				{text: '\u{27F3}', tooltip: 'WindowReload', onClick: () => runtime.WindowReload()},
			],
		},
		Folders: {
			'NoitaStats data folder': () => app.StartExplorer({opt: 'start_datafolder'}),
			'Noita installed path': () => app.StartExplorer({opt: 'noitapath'}),
			save00: () => app.StartExplorer({opt: 'save00'}),
			save_rec: () => app.StartExplorer({opt: 'save_rec'}),
		},
		WindowAction: {
			WindowShow: () => app.WindowAction({opt: 'WindowShow'}),
			WindowHide: () => app.WindowAction({opt: 'WindowHide'}),
			WindowReloadApp: () => app.WindowAction({opt: 'WindowReloadApp'}),
			WindowCenter: () => app.WindowAction({opt: 'WindowCenter'}),
		},
	};
	contextMenuStore().onContextMenu(e, items, divRef.getBoundingClientRect());
}

function toggleOpen(OpenKey) {
	if (OpenKey == undefined) OpenKey.open = false;
	else OpenKey.open = !OpenKey.open;
	return OpenKey.open;
}
</script>

<template>
	<main :ref="(el) => (refStore().main = el)">
		<span></span>
		<div class="side-menu" v-if="storeStore()?.config?.state?.sidemenu?.open && Object.keys(dataStore().raw.Sessions).length > 0">
			<TheSessionsList :allData="dataStore().$state" @resize-width="refStore().sidemenuResize" :selected="storeStore().state.selected" @selected="getNoitaData" />
		</div>

		<div class="container" :ref="(el) => (refStore().container = el)">
			<span :ref="(el) => (refStore().topMenu = el)">
				<div class="top-menu" style="z-index: 2" :ref="(el) => (refStore().mainTopMenu = el)">
					<Button
						:active="storeStore()?.config?.state?.sidemenu?.open"
						:text="[String.fromCodePoint('0x1F872'), String.fromCodePoint('0x1F870')]"
						:texts="[String.fromCodePoint('0x1F872'), String.fromCodePoint('0x1F870')]"
						:prop="{text: [String.fromCodePoint('0x1F872'), String.fromCodePoint('0x1F870')]}"
						:title="storeStore()?.config?.state?.sidemenu?.open.toString()"
						:onClick="
							(e) => {
								toggleOpen(storeStore()?.config?.state?.sidemenu);
								refStore().sidemenuResize(0);
							}
						"
					/>

					<Button :text="'All Data'" :active="storeStore()?.config?.state?.alldata?.open" @click="toggleOpen(storeStore()?.config?.state?.alldata)" />
					<Button :text="'Session'" :active="storeStore()?.config?.state?.session?.open" @click="toggleOpen(storeStore()?.config?.state?.session)" />
					<Button :text="'LifetimeCalc'" :active="storeStore()?.config?.state?.lifetime_calculator?.open" @click="storeStore().config.state.lifetime_calculator.open = !storeStore()?.config?.state?.lifetime_calculator?.open" />
					<Button :text="'Debug'" v-if="storeStore()?.config?.general?.debug" :active="storeStore()?.config?.state?.debug?.open" @click="storeStore().config.state.debug.open = !storeStore()?.config?.state?.debug?.open" />
					<div style="width: auto; display: flex; flex-direction: row; margin-left: auto; margin-right: 5px">
						<!-- justify-content:flex-end; right:0; height:2ch;-->
						<Button v-if="Object.keys(store.windowStore().handle).length > 0" :text="'&#128468;'" @click="onWindowsDropDown($event)" />
						<Button title="Close all" v-if="Object.keys(store.windowStore().windows).length > 0" :text="'&#128465;'" @click="Object.keys(store.windowStore().windows).map((id) => delete store.windowStore().windows[id])" />
						<Button title="Restore all" v-if="Object.keys(store.windowStore().windows).length > 0" :onDisabled="!(Object.keys(store.windowStore().handle).length > 0)" :text="'+'" @click="Object.keys(store.windowStore().windows).map((id) => store.windowStore().func.minimize(id, false))" />
						<Button
							title="Minimize all"
							v-if="Object.keys(store.windowStore().windows).length > 0"
							:onDisabled="Object.keys(store.windowStore().handle).length == Object.keys(store.windowStore().windows).length"
							:text="'-'"
							@click="Object.keys(store.windowStore().windows).map((id) => store.windowStore().func.minimize(id, true))"
						/>
						<ToggleButton
							:prop="{value: storeStore().config?.window?.always_on_top}"
							:value="storeStore().config?.window?.always_on_top"
							@click="
								storeStore().config.window.always_on_top = !storeStore().config?.window?.always_on_top;
								runtime.WindowSetAlwaysOnTop(storeStore().config?.window?.always_on_top);
							"
							:title="'AlwaysOnTop: ' + storeStore().config?.window?.always_on_top"
							:text="'Top'"
						/>
						<Button :text="'*'" @click="onDropDown($event)" />
					</div>
				</div>

				<div class="top-menu" v-if="storeStore()?.config?.state?.session?.open" style="top: 2.5ch; z-index: 1">
					<Button :text="'GameInfo'" :active="storeStore()?.config?.state?.session?.subviews?.game_info?.open" @click="toggleOpen(storeStore()?.config?.state?.session?.subviews?.game_info)" />
					<Button :text="'Wands'" :active="storeStore()?.config?.state?.session?.subviews?.wands?.open" @click="toggleOpen(storeStore()?.config?.state?.session?.subviews?.wands)" />
					<Button :text="'Items'" :active="storeStore()?.config?.state?.session?.subviews?.items?.open" @click="toggleOpen(storeStore()?.config?.state?.session?.subviews?.items)" />
					<Button :text="'Spells'" :active="storeStore()?.config?.state?.session?.subviews?.spells?.open" @click="toggleOpen(storeStore()?.config?.state?.session?.subviews?.spells)" />
					<Button :text="'Perks'" :active="storeStore()?.config?.state?.session?.subviews?.perks?.open" @click="toggleOpen(storeStore()?.config?.state?.session?.subviews?.perks)" />
					<Button :text="'Fungal Shifts'" :active="storeStore()?.config?.state?.session?.subviews?.fungal_shifts?.open" @click="toggleOpen(storeStore()?.config?.state?.session?.subviews?.fungal_shifts)" />
				</div>
			</span>

			<div class="content" :ref="(el) => (refStore().content = el)" style="">
				<template v-for="(value, key) in components" :key="key">
					<template v-if="store.storeStore().finishedLoading[key] || true">
						<component :is="value" v-if="storeStore()?.config?.state?.[key]?.open" />
					</template>
				</template>
			</div>
		</div>
	</main>
</template>

<style scoped>
.content {
	width: 100%;
}

.top-menu {
	position: sticky;
	top: 0px;
	float: top;
	white-space: nowrap;
	overflow: hidden;
	width: 100%;
	height: 2.5ch;
	background-color: #000000;
	display: flex;
	flex-direction: row;
}
</style>
