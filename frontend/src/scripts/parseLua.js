import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';

import * as utils from './../scripts/utils.js';
import {reactive, onMounted, computed, nextTick} from 'vue';
import {storeStore, dataStore} from './../stores/store.js'; //imagesStore

export async function luaInit() {
	//return
	//await localStorage.removeItem('lua_perk_list')
	//await localStorage.removeItem('lua_gun_actions')

	let lua_perk_list = await localStorage.getItem('lua_perk_list');
	if (lua_perk_list && lua_perk_list.length > 2) dataStore().LuaData.perk_list = await JSON.parse(lua_perk_list);
	else await parseLua(await app.GetNoitaDataFile(String.raw`data\scripts\perks\perk_list.lua`), dataStore().LuaData.perk_list, 'lua_perk_list');

	let lua_gun_actions = await localStorage.getItem('lua_gun_actions');
	if (lua_gun_actions && lua_gun_actions.length > 2) dataStore().LuaData.gun_actions = await JSON.parse(lua_gun_actions);
	else await parseLua(await app.GetNoitaDataFile(String.raw`data\scripts\gun\gun_actions.lua`), dataStore().LuaData.gun_actions, 'lua_gun_actions');

	//localStorage.setItem('lua_perk_list', JSON.stringify(dataStore().LuaData.perk_list))
	//localStorage.setItem('lua_gun_actions', JSON.stringify(dataStore().LuaData.gun_actions))
}
export function parseLua(str, resultObj = {}, name = '') {
	console.log('parseLua', name, resultObj);
	str = str.replaceAll(/perk_list =/gi, '');
	str = str.replaceAll(/actions =/gi, '');
	str = str.replaceAll(/--?.*--/gi, '');

	var re = /\.*(?<key>[A-z]+)\s*=\s*(shot_effects.|c.[A-z_]+)?\s*(?<value>[\.\w\W\u0020\u0024\u0022+-._/,{}]*)\s*(--)?/i; // ([--]|.$[,])?
	var objects = str.split('\r\n\t{');

	for (var i = 0; i < objects.length; i++) {
		var rows = objects[i].split('\r\n');
		//console.log('parseLuaPerks rows', rows)
		var resObj = {};
		for (let y = 0; y < rows.length; y++) {
			rows[y] = rows[y].trim();
			if (!rows[y] || rows[y] == 'end' || rows[y] == '') continue;
			let res = rows[y].match(re);
			if (res?.groups?.key && res?.groups?.value) {
				res.groups.value = res.groups.value.trim();
				if (res?.groups?.value.at(-1) == ',') res.groups.value = res.groups.value.substring(0, res.groups.value.length - 1);
				//if(/[+-]\s([0-9]*[.])?[0-9]+/) res.groups.value = res.groups.value.replace(' ', '')
				if (/[+-]\s([0-9]*[.])?[0-9]+/) {
					res.groups.value = res.groups.value.replace(' ', '');
				}

				if (res.groups.value.includes('{') && res.groups.value.includes('}')) {
					res.groups.value = res.groups.value.replaceAll('{', '[');
					res.groups.value = res.groups.value.replaceAll('}', ']');
				}
				try {
					res.groups.value = JSON.parse(res.groups.value);
				} catch (e) {}

				if (res?.groups?.value[0] == '+' || res?.groups?.value[0] == '-') res.groups.value = Number(res.groups.value);

				if (!resObj[res?.groups?.key]) resObj[res?.groups?.key] = res?.groups?.value;
			}
		}
		if (resObj?.id && !resultObj[resObj?.id] && !resObj?.id.includes('TESTBULLET')) resultObj[resObj.id] = resObj;
	}
	localStorage.setItem(name, JSON.stringify(resultObj));
	console.log('lua_list', name, resultObj);
}
