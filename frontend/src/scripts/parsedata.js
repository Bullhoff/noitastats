import * as utils from './../scripts/utils.js';
import {reactive, onMounted, computed, nextTick} from 'vue';
import {storeStore, imageStore, contextMenuStore, dataStore} from './../stores/store.js'; //imagesStore

export function JSONparse(data, key) {
	try {
		if (typeof data != 'string') return data;
		return JSON.parse(data);
	} catch (err) {
		return false;
	}
}
export function parsePlayerXml(data) {
	if (!data) return;
	dataStore().raw.PlayerXml = {...data};
	let res = iteratePlayerXml(dataStore().raw.PlayerXml);
	console.log('parsePlayerXml RES', res);
	dataStore().sorted.PlayerXml = res;
}
function iteratePlayerXml(data = {}, path = [], resObj = {}) {
	for (var [key, value] of Object.entries(data)) {
		let newPath = [...JSON.parse(JSON.stringify(path))];
		if (value.name == 'inventory_quick' || value.name == 'inventory_full') {
			newPath = [`${value.name}`];
			iteratePlayerXml(value.Entity, newPath, resObj);
		} else if (value.tags?.includes('perk_entity')) {
			newPath = [`perks`];
			if (!resObj.PerksAll) resObj.PerksAll = [];
			if (!resObj.Perks) resObj.Perks = [];
			if (!resObj.Perks_GameEffectComponent) resObj.Perks_GameEffectComponent = [];
			if (!resObj.Perks_MagicConvertMaterialComponent) resObj.Perks_MagicConvertMaterialComponent = [];
			value = {
				$: structureData(value),
				...value,
			};
			resObj.PerksAll.push(value);
			if (value.UIIconComponent?.[0]?.is_perk) resObj.Perks.push(value);
			else if (value.GameEffectComponent?.[0]) resObj.Perks_GameEffectComponent.push(value);
			else if (value.MagicConvertMaterialComponent?.[0]) resObj.Perks_MagicConvertMaterialComponent.push(value);
		} else if (path[0] == 'inventory_full') {
			newPath = [...path, `spells`];
			if (!resObj.Spells) resObj.Spells = [];
			value = {
				$: structureData(value),
				...value,
			};
			resObj.Spells.push(value);
		} else if (value.tags?.includes('wand') && !value.tags?.includes('broken_wand')) {
			newPath = [...path, `wands`];
			if (!resObj.Wands) resObj.Wands = [];
			value = {
				$: structureData(value),
				...value,
			};
			resObj.Wands.push(value);
			iteratePlayerXml(value.Entity, newPath, resObj.Wands.at(-1));
		} else if (path[0] == 'inventory_quick' && value.tags?.includes('item')) {
			// item_pickup
			newPath = [...path, `items`];
			if (!resObj.Items) resObj.Items = [];
			value = {
				$: structureData(value),
				...value,
			};
			resObj.Items.push(value);
		} else if (value.tags?.includes('card_action') && value.ItemComponent[0].permanently_attached == 1) {
			// UIIconComponent["-is_perk"]
			newPath = [`spells`];
			if (!resObj.Spells_alwayscast) resObj.Spells_alwayscast = [];
			value = {
				$: structureData(value),
				...value,
			};
			resObj.Spells_alwayscast.push(value);
		} else if (value.tags?.includes('card_action')) {
			newPath = [`spells`];
			if (!resObj.Spells) resObj.Spells = [];
			value = {
				$: structureData(value),
				...value,
			};
			resObj.Spells.push(value);
		} else if (key == 'Entity' || value.Entity) {
			iteratePlayerXml(value, newPath, resObj);
		} else {
		}
	}
	return resObj;
}
function structureData(obj) {
	let res = {};
	let filepath = null;
	if (!filepath) filepath = obj?.ItemComponent?.[0]?.['ui_sprite']; // data/ui_gfx/items/potion.png
	if (!filepath) filepath = obj?.SpriteComponent?.[0]?.['image_file']; // data/items_gfx/potion.png
	if (!filepath) filepath = obj?.SpriteComponent?.[0]?.['image_file'];
	if (!filepath) filepath = obj?.UIIconComponent?.[0]?.['icon_sprite_file'];
	if (filepath?.includes('.xml')) filepath = filepath.replace('.xml', '.png');
	res['img'] = filepath;
	if (obj?.UIIconComponent?.[0]) {
		if (obj?.UIIconComponent?.[0]?.description) res['_id'] = obj?.UIIconComponent?.[0]?.description; // ID
		if (obj?.UIIconComponent?.[0]?.description) {
			res['type'] = 'perk';
			res['id'] = obj?.UIIconComponent?.[0]?.description.split('_');
			res['id'].shift();
			res['id'] = res['id'].join('_').toUpperCase();
		}
	}
	if (obj?.GameEffectComponent?.[0]) {
		if (obj?.GameEffectComponent?.[0]?.effect) res['id'] = obj?.GameEffectComponent?.[0]?.effect; // ID
	}
	if (obj?.ItemComponent?.[0]) {
		if (obj?.ItemComponent?.[0]?.item_name) res['id'] = obj?.ItemComponent?.[0]?.item_name; // ID
		if (obj?.ItemComponent?.[0]?.['inventory_slot.x']) res['x'] = obj?.ItemComponent?.[0]?.['inventory_slot.x'];
		if (obj?.ItemComponent?.[0]?.['inventory_slot.y']) res['y'] = obj?.ItemComponent?.[0]?.['inventory_slot.y'];
		if (obj?.ItemComponent?.[0]?.permanently_attached) res['permanently_attached'] = obj?.ItemComponent?.[0]?.permanently_attached;
		if (obj?.ItemComponent?.[0]?.uses_remaining) res['uses_remaining'] = obj?.ItemComponent?.[0]?.uses_remaining;
	}
	if (obj?.AbilityComponent?.[0]) {
		if (obj?.AbilityComponent?.[0]?.ui_name) res['id'] = obj?.AbilityComponent?.[0]?.ui_name; // ID
		if (obj?.AbilityComponent?.[0]?.sprite_file) res['sprite_file'] = obj?.AbilityComponent?.[0]?.sprite_file;
		if (obj?.AbilityComponent?.[0]?.gun_level) res['gun_level'] = obj?.AbilityComponent?.[0]?.gun_level;
		if (obj?.AbilityComponent?.[0]?.gun_config?.[0]?.deck_capacity) res['deck_capacity'] = obj?.AbilityComponent?.[0]?.gun_config?.[0]?.deck_capacity;
		if (obj?.AbilityComponent?.[0]?.gun_config?.[0]?.actions_per_round) res['actions_per_round'] = obj?.AbilityComponent?.[0]?.gun_config?.[0]?.actions_per_round;
		if (obj?.AbilityComponent?.[0]?.gun_config?.[0]?.reload_time) res['reload_time'] = obj?.AbilityComponent?.[0]?.gun_config?.[0]?.reload_time;
		if (obj?.AbilityComponent?.[0]?.gun_config?.[0]?.shuffle_deck_when_empty) res['shuffle_deck_when_empty'] = obj?.AbilityComponent?.[0]?.gun_config?.[0]?.shuffle_deck_when_empty;
	}
	if (obj?.ItemActionComponent?.[0]) {
		if (obj?.ItemActionComponent?.[0]?.action_id) res['id'] = obj?.ItemActionComponent?.[0]?.action_id; // ID
	}
	/* if(obj?.PhysicsBodyComponent?.[0]){
				res['hax_wait_till_pixel_scenes_loaded'] = obj.PhysicsBodyComponent?.[0]?.hax_wait_till_pixel_scenes_loaded
			} */
	return res;
}

export function parseWorldStateXml(data) {
	if (!data?.WorldStateComponent?.[0]?.lua_globals) return;
	dataStore().raw.WorldStateXml = data;
	dataStore().sorted.WorldStateXml = {
		//flags: [],
		changed_materials: [],
		lua_globals_: [],
		lua_globals: {
			ORB_MAP_STRING: '',
			PICKUP_COUNT: {},
			TOTAL_COUNT: {},
			COUNT: {},
			LEVEL: {},
			TEMPLE_ACTIVE: {},
			TEMPLE_COLLAPSED: {},
			TEMPLE_LEAKED: {},
			other: {},
		},
	};
	if (dataStore().raw?.WorldStateXml?.WorldStateComponent?.[0]?.changed_materials?.[0]?.string)
		dataStore().sorted.WorldStateXml.changed_materials = dataStore().raw?.WorldStateXml?.WorldStateComponent?.[0]?.changed_materials?.[0]?.string.reduce((res, value, i, array) => {
			if (i % 2 === 0) res.push(array.slice(i, i + 2));
			return res;
		}, []);
	for (let i = 0; i < dataStore().raw.WorldStateXml.WorldStateComponent[0].lua_globals[0].E.length; i++) {
		let {key, value} = dataStore().raw.WorldStateXml.WorldStateComponent[0].lua_globals[0].E?.[i];
		if (key == 'ORB_MAP_STRING') dataStore().sorted.WorldStateXml.lua_globals.ORB_MAP_STRING = value;
		else if (key.includes('_PICKUP_COUNT')) dataStore().sorted.WorldStateXml.lua_globals.PICKUP_COUNT[key] = value;
		else if (key.includes('_TOTAL_COUNT')) dataStore().sorted.WorldStateXml.lua_globals.TOTAL_COUNT[key] = value;
		else if (key.includes('_COUNT')) dataStore().sorted.WorldStateXml.lua_globals.COUNT[key] = value;
		else if (key.includes('_LEVEL')) dataStore().sorted.WorldStateXml.lua_globals.LEVEL[key] = value;
		else if (key.includes('TEMPLE_ACTIVE_')) {
			let [temple, active, x, y] = key.split('_');
			dataStore().sorted.WorldStateXml.lua_globals.TEMPLE_ACTIVE[`${x}, ${y}`] = value;
		} else if (key.includes('TEMPLE_COLLAPSED_')) {
			let [temple, collapsed, x, y] = key.split('_');
			dataStore().sorted.WorldStateXml.lua_globals.TEMPLE_COLLAPSED[`${x}, ${y}`] = value;
		} else if (key.includes('TEMPLE_LEAKED_')) {
			let [temple, leaked, x, y] = key.split('_');
			dataStore().sorted.WorldStateXml.lua_globals.TEMPLE_LEAKED[`${x}, ${y}`] = value;
		} else dataStore().sorted.WorldStateXml.lua_globals.other[key] = value;
		dataStore().sorted.WorldStateXml.lua_globals_.push(`${dataStore().raw.WorldStateXml.WorldStateComponent[0].lua_globals[0].E?.[i].key} = ${dataStore().raw.WorldStateXml.WorldStateComponent[0].lua_globals[0].E?.[i].value}`);
	}
}

export function parseSessions_kills(data, dataRaw, dataSorted) {
	if (!data) return;
	for (const [key, value] of Object.entries(data)) {
		let [k, v] = [key.split('_')[0], key.split('_')[1]];
		if (!dataRaw.Sessions[k]) dataRaw.Sessions[k] = {};
		let json = JSONparse(value, 'Sessions_kills');
		if (!json) return;
		dataRaw.Sessions[k]['kills'] = json;
		dataRaw.Sessions[k]['Gameid'] = k;
		dataRaw.Sessions[k]['DateTime'] = utils.gameidToDateTimeStr(key);
	}
}
export function parseSessions_stats(data, dataRaw, dataSorted) {
	if (!data) return;
	for (const [key, value] of Object.entries(data)) {
		let [k, v] = [key.split('_')[0], key.split('_')[1]];
		if (!dataRaw.Sessions[k]) dataRaw.Sessions[k] = {};
		let json = JSONparse(value, 'Sessions_stats');
		if (!json) return;
		dataRaw.Sessions[k]['stats'] = json;
		dataRaw.Sessions[k]['Gameid'] = k;
		dataRaw.Sessions[k]['DateTime'] = utils.gameidToDateTimeStr(key);
		if (!dataRaw.Sessions[k]['$']) dataRaw.Sessions[k]['$'] = {};
		dataRaw.Sessions[k]['$']['datetime'] = utils.gameidToDateTimeStr(key);
	}
	dataStore().init();
}

export function parseStatsXml(data, raw, dataSorted) {
	if (!data) return;
	for (const [key, value] of Object.entries(data)) {
		raw.Stats[key] = value;
	}
}

export function parseSessionsXml(data, raw) {
	if (!data) return;
	for (const [key, value] of Object.entries(data)) {
		raw.Sessions[key] = {
			$: {gameid: key, datetime: utils.gameidToDateTimeStr(key)},
			...value,
		};
	}
	heater([1, 'stats', '0', 'playtime'], '-playtime');
	heater([1, '$', 'datetime'], '-datetime');
}

function type(val) {
	if (Array.isArray(val)) return 'array';
	if (typeof val == 'object') return 'object';
	if (!isNaN(Date.parse(val))) return 'date';
	if (val == 'true' || val == 'false') return 'bool';
	if (isNaN(val)) return 'string';
	val = Number(val);
	if (isFloat(val)) return 'float'; //parseFloat(val)
	if (isInt(val)) return 'int'; //parseInt(val)
}
function isInt(n) {
	return Number(n) === n && n % 1 === 0;
}
function isFloat(n) {
	return Number(n) === n && n % 1 !== 0;
}
function parseType(arr) {
	var newArr = Array.isArray(arr) ? arr : [arr];
	for (let i = 0; i < newArr.length; i++) {
		let t = type(newArr[i]);
		if (t == 'date') {
			newArr[i] = Date.parse(newArr[i]);
		} else if (t == 'int') {
			newArr[i] = parseInt(newArr[i]);
		} else if (t == 'float') {
			newArr[i] = parseFloat(newArr[i]);
		}
	}
	arr = newArr;
	return arr;
}
function heater(path = [], key, colorLow, colorHigh) {
	let arrTemp = Object.entries(dataStore().raw.Sessions);
	let arr = arrTemp.sort((a, b) => {
		// utils.getObjectValue(a, [...path])-utils.getObjectValue(b, [...path])
		let valA = utils.getObjectValue(a, [...path]);
		let valB = utils.getObjectValue(b, [...path]);
		let res = parseType([valA, valB]);
		valA = res[0];
		valB = res[1];
		return valA - valB;
	});
	let minValue = utils.getObjectValue(arr[0], [...path]);
	let maxValue = utils.getObjectValue(arr.at(-1), [...path]);
	let res = parseType([minValue, maxValue]);
	minValue = res[0];
	maxValue = res[1];
	const ratio = 255 / (maxValue - minValue);
	const remove = minValue * ratio;
	if (!key) key = `${path.at(-1)}`;
	for (let i = 0; i < arr.length; i++) {
		if (!arr[i][1]['$']) arr[i][1]['$'] = {};
		if (!arr[i][1]['$'][key]) arr[i][1]['$'][key] = {};
		if (!arr[i][1]['$'][key]['style']) arr[i][1]['$'][key]['style'] = {};
		arr[i][1]['$'][key]['$'] = `${i} / ${arr.length - 1}`;
		let value = utils.getObjectValue(arr[i], [...path]);
		let res = parseType([value]);
		value = res[0];
		arr[i][1]['$'][key]['style']['color'] = `rgb(${Math.round(value * ratio - remove).toString()}, ${Math.round(255 - value * ratio + remove).toString()}, ${10})`;
		arr[i][1]['$'][key]['place'] = [i, arr.length - 1];
	}
	nextTick().then(() => {
		let obj = arr.reduce((acc, c) => {
			return {...acc, [c[0]]: c[1]};
		}, {});
		dataStore().raw.Sessions = obj;
	});
}
