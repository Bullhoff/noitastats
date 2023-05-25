<script setup>
import {ref, reactive, onMounted, onBeforeMount, onUpdated, nextTick, computed} from 'vue';
//import _ from 'lodash';
import {imageStore, dataStore} from './../stores/store.js';
import * as store from './../stores/store.js';

const emit = defineEmits(['emit-handler', '']);
const props = defineProps(
	reactive({
		obj: {
			type: Object,
			required: false,
			default: {},
		},
	})
);
const refs = reactive({
	container: null,
});

const state = reactive({
	tooltip: false,
	info: false,
	toggle: false,
	everything: false,
});

function uniquekey(obj, key, i = 1) {
	i += 1;
	if (obj[key]) return uniquekey(obj, key + '_' + i, i);
	return key;
}
function getLastKeyValue(startObj, obj = {}) {
	if (!startObj) return obj;
	for (var [key, value] of Object.entries(startObj)) {
		if (key[0] == '$' || key == 'Entity' || key == 'Spells_alwayscast' || key == 'Spells') {
		} else if (typeof value == 'object') Object.assign(obj, getLastKeyValue(value));
		else {
			key = uniquekey(obj, key);
			if (key.includes('damage_critical')) value = value * 25;
			else if (key.includes('damage')) value = value * 25;
			else if (key.includes('rate')) value = (value / 60).toFixed(2);
			if (key == 'current_reload_time') key = 'reload_time';
			if (key == 'mana') key = 'mana_drain';
			obj[key] = value;
		}
	}
	return obj;
}
function update() {
	var fullData = {};
	fullData.PlayerXml = props.obj;
	let id = 'ITEM';
	if (props.obj?.AbilityComponent?.[0]?.ui_name) {
		id = props.obj?.AbilityComponent?.[0]?.ui_name;
	}
	if (props.obj?.ItemActionComponent?.[0]?.action_id) {
		id = props.obj.ItemActionComponent?.[0].action_id;
		fullData.gun_actions = dataStore().LuaData?.gun_actions?.[props.obj.ItemActionComponent?.[0].action_id];
	}
	if (props.obj?.UIIconComponent?.[0]?.name) {
		let name = props.obj?.UIIconComponent?.[0]?.name.split('_');
		name.shift();
		id = name.join(' ');
		name = name.join('_').toUpperCase();
		if (dataStore().LuaData?.perk_list?.[name]) fullData.gun_actions = dataStore().LuaData?.perk_list?.[name];
	}
	var obj = getLastKeyValue(fullData);
	refs.everything.innerHTML = '';
	refs.container.innerHTML = '';
	var arr = Object.entries(obj).sort((a, b) => a[0].localeCompare(b[0]));
	for (let i = 0; i < arr.length; i++) {
		if (!arr[i][1]) continue;
		let imgUrl = 'data/ui_gfx/inventory/icon_' + arr[i][0] + '.png';
		let img = imageStore().getImage(imgUrl);
		var tooltip = `${arr[i][0]} = ${arr[i][1]}`;
		tooltip = JSON.stringify(arr[i], null, 2);

		if (!imageStore().images[imgUrl]) {
			const text = document.createElement('span');
			text.innerText = arr[i][0];
			text.style.textAlign = 'left';
			text.title = tooltip;
			refs.everything.appendChild(text);
			const val = document.createElement('span');
			val.innerText = arr[i][1];
			val.style.textAlign = 'left';
			val.title = tooltip;
			refs.everything.appendChild(val);
		} else {
			const img = document.createElement('img');
			img.src = imageStore().getImage('data/ui_gfx/inventory/icon_' + arr[i][0] + '.png');
			img.alt = '';
			img.width = '8';
			img.height = '8';
			img.title = tooltip;
			refs.container.appendChild(img);
			const val = document.createElement('span');
			val.innerText = arr[i][1];
			val.style.textAlign = 'left';
			val.style.width = '3ch';

			val.title = tooltip;
			refs.container.appendChild(val);
		}
	}
	const btn = document.createElement('button');
	btn.style.height = '10px';
	btn.style.gridColumn = 'span 2';
	btn.addEventListener('click', ($event) => store.windowStore().createNewWindow({component: 'TreeView', title: id, e: $event, prop: {head: {text: id}, body: {obj: JSON.parse(JSON.stringify(fullData))}}}));
	refs.container.appendChild(btn);
}
onUpdated(() => {
	update();
});

onBeforeMount(async () => {});
onMounted(async () => {
	update();
});
</script>

<template>
	<div :class="[]" class="container">
		<span class="entity-info" :ref="(el) => (refs.container = el)"> </span>
		<span class="entity-info" :ref="(el) => (refs.everything = el)" v-show="state.everything"> </span>
	</div>
</template>

<style scoped>
.container {
	width: inherit;
}
.entity-info {
	padding: 0 1px 0 1px;
	font-size: 8px;
	display: grid;
	grid-template-columns: auto auto;
	gap: 1px;
	overflow: hidden;
}
</style>
