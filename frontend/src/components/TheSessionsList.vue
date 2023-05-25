<script setup>
import {reactive, onMounted, onUpdated, nextTick, computed} from 'vue';
import {storeStore, contextMenuStore, refStore, dataStore} from './../stores/store.js'; //imagesStore
import * as app from '../../wailsjs/go/main/App';
import * as store from './../stores/store.js'; //imagesStore
import {storeToRefs} from 'pinia';
const {config} = storeToRefs(storeStore());
const emit = defineEmits(['resize-width', 'selected']);
const props = defineProps({});
const refs = reactive({});

const symbols = {
	file: String.fromCodePoint('0x1F5CE'),
	circle: String.fromCodePoint('0x25CF'), // &#x23FA;		0x025CF 0x025C9 0x2B24 23FA, 2218, 20D8, 1F5CE, 229A, 2388	String.fromCodePoint('0x23FA')
	star: String.fromCodePoint('0x2605'), // &#9733;
	speechBubble: String.fromCodePoint('0x1F5E9'), // &#x1F5E9;
};
const color = {
	red: 'red',
	green: 'rgb(0, 255, 85)',
	gray: 'rgb(80,80,80)',
	yellow: 'yellow',
};
function selectRow(gameid) {
	emit('selected', gameid);
}

onMounted(() => {
	refStore().sideMenu.scrollTo(0, refStore().sideMenu.scrollHeight);
	new ResizeObserver(onResize).observe(refStore().sideMenu);
});
var currentLength = 0;
onUpdated(() => {
	if (dataStore().getSortedSessionsList.length != currentLength) {
		currentLength = dataStore().getSortedSessionsList.length;
		refStore().sideMenu.scrollTo(0, refStore().sideMenu.scrollHeight);
	}
});

function onResize() {
	try {
		let rect = refStore().sideMenu.getBoundingClientRect();
		emit('resize-width', rect.width);
	} catch {
		emit('resize-width', 0);
	}
}

function toggle(e) {
	if (storeStore().config?.state?.sidemenu?.[e] == undefined) storeStore().config.state.sidemenu[e] = false;
	storeStore().config.state.sidemenu[e] = !storeStore().config.state?.sidemenu[e];
}

function saveComment(resObj) {
	let {text, obj} = resObj;
	let {id} = obj;
	if (!store.dataStore().SessionJson[obj.id]) store.dataStore().SessionJson[id] = {comment: '', favorite: false};
	store.dataStore().SessionJson[id].favorite = store.dataStore().SessionJson[id].favorite;
	store.dataStore().SessionJson[id].comment = text;
	app.UpdateSessionJson(id, store.dataStore().SessionJson[id]);
}

const commentWindowObj = (id, e) => {
	return {
		component: 'TextBox',
		title: id,
		e,
		prop: {
			head: {
				text: id,
			},
			body: {
				obj: {id: id, ...store.dataStore().SessionJson[id]},
				click: saveComment,
			},
		},
	};
};
function contextMenu(e, id) {
	let items = {
		'Sort by': {
			'Sort by': {
				type: 'radioGroup',
				array: [
					{
						text: 'Datetime',
						id: 'datetime',
						name: 'sort_by',
						onChange: (e, name) => {
							console.log('click', name, e.target, e);
							storeStore().config.state.sidemenu.sort_by = name;
						},
						active: storeStore().config?.state?.sidemenu?.sort_by,
					},
					{text: 'Playtime', id: 'playtime', name: 'sort_by', onChange: (e, name) => (storeStore().config.state.sidemenu.sort_by = name), active: storeStore().config?.state?.sidemenu?.sort_by},
				],
			},
		},
		display_filestatus: {text: 'Display filestatus', type: 'checkbox', onChange: () => toggle('display_filestatus'), value: storeStore().config?.state?.sidemenu?.['display_filestatus']},
		display_input: {text: 'Display input', type: 'checkbox', onChange: () => toggle('display_input'), value: storeStore().config?.state?.sidemenu?.['display_input']},
		display_extra: {text: 'Display extra', type: 'checkbox', onChange: () => toggle('display_extra'), value: storeStore().config?.state?.sidemenu?.['display_extra']},
		Comment: () => store.windowStore().createNewWindow(commentWindowObj(id, e)),
	};
	contextMenuStore().onContextMenu(e, items);
}
</script>

<template v-if="dataStore().getSortedSessionsList">
	<div class="container prevent-select" style="cursor: pointer" :ref="(el) => (refStore().sideMenu = el)" @scroll="" @resize="onResize">
		<li v-for="(value, key) in dataStore().getSortedSessionsList" style="list-style-type: none; white-space: nowrap" @contextmenu.prevent="contextMenu($event, value.$.gameid)">
			<span class="file-status prevent-select" style="" v-if="config.state?.sidemenu?.display_filestatus">
				<span :title="'/player.xml'" :ref="(el) => (refs[value.$.gameid + '__player.xml'] = el)" :class="[dataStore().raw.AvailabeFiles[value.$.gameid + '\\player.xml'] ? 'file-exists' : '']">{{ symbols.file }}</span>
				<span :title="'/world_state.xml'" :ref="(el) => (refs[value.$.gameid + '__world_state.xml'] = el)" :class="[dataStore().raw.AvailabeFiles[value.$.gameid + '\\world_state.xml'] ? 'file-exists' : '']">{{ symbols.file }}</span>
				<span :title="`/stats/sessions/${value.$.gameid}_kills.xml`" :ref="(el) => (refs[value.$.gameid + '__kills.xml'] = el)" :class="[value.kill_map ? 'file-exists' : '']">{{ symbols.file }}</span>
				<span :title="`/stats/sessions/${value.Gameid}_stats.xml`" :ref="(el) => (refs[value.$.gameid + '__stats.xml'] = el)" :class="[value.stats ? 'file-exists' : '']">{{ symbols.file }}</span>
			</span>

			<span class="status-container" v-if="storeStore().config?.state?.sidemenu?.display_input">
				<span
					:style="[store.dataStore().SessionJson[value.$.gameid]?.comment && store.dataStore().SessionJson[value.$.gameid]?.comment != '' ? {} : {color: color.gray}]"
					:title="store.dataStore().SessionJson[value.$.gameid]?.comment"
					@click="store.windowStore().createNewWindow(commentWindowObj(value.$.gameid, $event))"
				>
					{{ symbols.speechBubble }}
				</span>
				<span
					:style="[store.dataStore().SessionJson[value.$.gameid]?.favorite ? {color: color.yellow} : {color: color.gray}]"
					style="cursor: pointer"
					@click="
						() => {
							if (!dataStore().SessionJson[value.$.gameid]) dataStore().SessionJson[value.$.gameid] = {favorite: false};
							dataStore().SessionJson[value.$.gameid].favorite = !dataStore().SessionJson[value.$.gameid].favorite;
							app.UpdateSessionJson(value.$.gameid, dataStore().SessionJson[value.$.gameid]);
						}
					"
				>
					{{ symbols.star }}
				</span>
			</span>
			<span class="row prevent-select" :class="[dataStore().raw.Gameid == value.$.gameid ? 'row-selected' : '']" @click="selectRow(value.$.gameid)">{{ value.$.datetime }} </span>

			<span class="status-container prevent-select" v-if="storeStore().config?.state?.sidemenu?.display_extra" @click="test">
				<span>
					<span :title="'dead: ' + value?.stats?.[0]?.dead + ' killed by ' + value?.stats?.[0]?.killed_by" :style="[value?.stats?.[0]?.dead == 0 ? {color: color.green} : {color: color.red}]">{{ symbols.circle }}</span>
				</span>
				<span>
					<span :title="'playtime_str: ' + value?.stats?.[0]?.playtime_str" style="color: gray" :style="value?.['$']?.['-playtime']?.style">{{ symbols.circle }}</span>
				</span>
			</span>
		</li>
	</div>
</template>

<style scoped>
.status-container {
	font-size: 1ch;
	vertical-align: middle;
}
.file-status,
.inactive {
	font-size: 1ch;
	color: red;
	vertical-align: middle;
}
.file-exists,
.active {
	color: rgb(0, 255, 85);
}

.container {
	z-index: 3;
	background-color: black;
	height: 100%;
	position: fixed;
	overflow-y: scroll;
	resize: horizontal;
}
.row {
	cursor: pointer;
	border: 1px solid rgb(19, 18, 18);
	padding: 0px 5px 0px 5px;
	white-space: nowrap;
}
.row:hover {
	background-color: rgb(0, 106, 255);
}
.row-selected {
	background-color: rgb(0, 13, 255);
}
</style>
