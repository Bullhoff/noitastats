import {ref, computed, reactive, watch, onMounted, nextTick} from 'vue';
import {defineStore} from 'pinia';
import {storeToRefs} from 'pinia';
import {storeStore} from './store.js';
import _ from 'lodash';
import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import * as utils from './../scripts/utils.js';
import * as parse from './../scripts/parse.js';

function collectTime(name, op) {
	if(!storeStore()?.config?.general?.debug) return
	if (!storeStore().latestCollect[name] || typeof storeStore().latestCollect[name] == 'string') storeStore().latestCollect[name] = {};
	storeStore().latestCollect[name][op] = performance.now();
	if (op == 'done') storeStore().latestCollect[name] = Math.round(storeStore().latestCollect[name]['done'] - storeStore().latestCollect[name]['start']) + 'ms';
}
function setNest(path = [], val, resObj = {}) {
	let next = path.shift();
	if (path.length > 0) {
		if (!resObj[next]) resObj[next] = {};
		return setNest(path, val, resObj[next]);
	}
	resObj[next] = val;
	return resObj;
}

export const dataStore = defineStore('data', {
	state: () => {
		return {
			raw: {Gameid: '', lastSync: '', PlayerXml: {}, WorldStateXml: {}, Stats: {}, Sessions: {}, AvailabeFiles: {}},
			sorted: {PlayerXml: {}, WorldStateXml: {}},
			dataRaw: {Player: {}, Sessions: {}, Stats: {}},
			dataSorted: {InventoryEntities: [], GameData: {}, WorldState: {}, FungalShifts: [], Wands: {}, Pocket: {}, Backpack: {}, Perks: {}, Perks_: {}, Perks_all: {}, Other: {}},
			SessionJson: {},
			templateData: {},
			DataEntityListHolder: {},
			LuaData: {perk_list: {}, gun_actions: {}},
			EntityTemplate: [],
		};
	},
	getters: {
		xmlAndLua:
			(state) =>
			(obj = dataStore().LuaData.gun_actions, resObj = {}, path) => {
				// obj=dataStore().DataEntityList
				/* gun_actions:{}, perk_list:{} */
				//if(!resObj.spellDropdown) resObj.spellDropdown = {}
				if (!resObj.perk_list) resObj.perk_list = {};
				resObj.perk_list = dataStore().LuaData.perk_list;

				for (var [key, value] of Object.entries(obj)) {
					if (!resObj.gun_actions) resObj.gun_actions = {};
					if (!resObj.gun_actions[key]) resObj.gun_actions[key] = {};
					resObj.gun_actions[key].lua = value;
					if (value.custom_xml_file) {
						if (!resObj.gun_actions[key].custom_xml_file) resObj.gun_actions[key].custom_xml_file = [];
						if (!Array.isArray(value.custom_xml_file)) obj[key].custom_xml_file = [value.custom_xml_file];
						for (let i = 0; i < obj[key].custom_xml_file.length; i++) {
							if (dataStore().DataEntityList[obj[key].custom_xml_file[i]]) resObj.gun_actions[key].custom_xml_file.push(dataStore().DataEntityList[obj[key].custom_xml_file[i]]);
						}
					}

					if (value.related_projectiles) {
						if (!resObj.gun_actions[key].related_projectiles) resObj.gun_actions[key].related_projectiles = [];
						if (!Array.isArray(value.related_projectiles)) obj[key].related_projectiles = [value.related_projectiles];
						for (let i = 0; i < obj[key].related_projectiles.length; i++) {
							if (dataStore().DataEntityList[obj[key].related_projectiles[i]]) resObj.gun_actions[key].related_projectiles.push(dataStore().DataEntityList[obj[key].related_projectiles[i]]);
						}
					}

					if (value.related_extra_entities) {
						if (!resObj.gun_actions[key].related_extra_entities) resObj.gun_actions[key].related_extra_entities = [];
						if (!Array.isArray(value.related_extra_entities)) obj[key].related_extra_entities = [value.related_extra_entities];
						for (let i = 0; i < obj[key].related_extra_entities.length; i++) {
							if (dataStore().DataEntityList[obj[key].related_extra_entities[i]]) resObj.gun_actions[key].related_extra_entities.push(dataStore().DataEntityList[obj[key].related_extra_entities[i]]);
						}
					}
				}
				return resObj;
			},
		perkList(state) {
			let resObj = {};
			for (var [key, value] of Object.entries(this.xmlAndLua().perk_list)) {
				let obj = {};
				utils.setObjectValue({obj: value, key: 'id', resObj: obj});
				utils.setObjectValue({obj: value, key: 'ui_icon', resObj: obj});
				utils.setObjectValue({obj: value, key: 'stackable', resObj: obj});
				utils.setObjectValue({obj: value, key: 'max_in_perk_pool', resObj: obj});
				utils.setObjectValue({obj: value, key: 'usable_by_enemies', resObj: obj});
				utils.setObjectValue({obj: value, key: 'game_effect', resObj: obj});
				utils.setObjectValue({obj: value, key: 'stackable_is_rare', resObj: obj});
				utils.setObjectValue({obj: value, key: 'one_off_effect', resObj: obj});
				utils.setObjectValue({obj: value, key: 'do_not_remove', resObj: obj});
				utils.setObjectValue({obj: value, key: 'remove_other_perks', resObj: obj, func: (x) => x.join(', ')});
				resObj[key] = obj;
			}
			return resObj;
		},
		spellList(state) {
			let resObj = {};
			if (!this.xmlAndLua()?.gun_actions) return resObj;
			for (var [key, value] of Object.entries(this.xmlAndLua().gun_actions)) {
				let obj = {};
				if (value.lua.id) obj.id = value.lua.id;
				if (value.lua.type) obj.type = value.lua.type;
				if (value.lua.pattern_degrees) obj.pattern_degrees = value.lua.pattern_degrees;
				if (value.lua.mana) obj.mana = value.lua.mana;
				if (value.lua.max_uses) obj.max_uses = value.lua.max_uses;
				if (value.lua.fire_rate_wait) obj.fire_rate_wait = value.lua.fire_rate_wait;
				if (value.lua.lifetime_add) obj.lifetime_add = value.lua.lifetime_add;
				if (value.lua.damage_critical_chance) obj.damage_critical_chance = value.lua.damage_critical_chance;
				if (value.lua.price) obj.price = value.lua.price;
				if (value.lua.screenshake) obj.screenshake = value.lua.screenshake;
				if (value.lua.sprite) obj.sprite = value.lua.sprite;
				if (value.lua.spawn_level) obj.spawn_level = value.lua.spawn_level;
				if (value.lua.spawn_probability) obj.spawn_probability = value.lua.spawn_probability;
				if (value.lua.spawn_level) obj.spawn_levels = value.lua.spawn_level.split(',--')[0].replaceAll('"', '').split(',');
				if (value.lua.spawn_probability) obj.spawn_probabilities = value.lua.spawn_probability.split(',--')[0].replaceAll('"', '').split(',');

				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'damage', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'lifetime', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'lifetime_randomness', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'knockback_force', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0]?.config_explosion?.[0], key: 'damage', resKey: 'explosion_damage', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'speed_min', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'speed_max', resObj: obj});
				utils.setObjectValue({obj: value?.related_projectiles?.[0]?.ProjectileComponent?.[0], key: 'projectile_type', resObj: obj});
				resObj[key] = obj;
			}
			return resObj;
		},

		getLocalStorage() {
			let res = {images: {}};
			for (const [key, value] of Object.entries(localStorage)) {
				if (key.split('/')[0] == 'data') res.images[key] = value;
				else res[key] = JSON.parse(value);
			}
			return res;
		},
		DataEntityEntity:
			(state) =>
			(obj = dataStore().DataEntityList, resObj = {}) => {
				//dataStore().DataEntityListNested.data.entities.projectiles.deck
				if (!dataStore().DataEntityListNested()?.data?.entities?.projectiles) return;
				for (const [key, value] of Object.entries(dataStore().DataEntityListNested().data.entities.projectiles)) {
					let split = key.split('/');
					if (split.at(-1) == 'deck') continue;
					resObj[split.at(-1)] = value;
				}
				if (!dataStore().DataEntityListNested()?.data?.entities?.projectiles?.deck) return;
				for (const [key, value] of Object.entries(dataStore().DataEntityListNested().data.entities.projectiles.deck)) {
					let split = key.split('/');
					//resObj[split.at(-1)] = value
					resObj[split.at(-1)] = _.merge(resObj[split.at(-1)], value);
				}
				return resObj;
			},
		DataEntityListNested:
			(state) =>
			(obj = dataStore().DataEntityList, resObj = {}) => {
				for (const [key, value] of Object.entries(obj)) {
					key.split('/');
					setNest(key.split('/'), value, resObj);
				}
				return resObj;
			},

		DataEntityList(state) {
			if (Object.keys(state.DataEntityListHolder).length == 0 && localStorage.getItem('DataEntityList')) {
				let res = localStorage.getItem('DataEntityList');
				if (res) state.DataEntityListHolder = JSON.parse(res);
			} else if (Object.keys(state.DataEntityListHolder).length == 0) {
				app.GetDataXmlFiles();
				return {};
			}
			return state.DataEntityListHolder;
		},

		FungalShifts(state) {
			return state.sorted.WorldStateXml.changed_materials.reduce((res, value, i, array) => {
				if (i % 2 === 0) res.push(array.slice(i, i + 2));
				return res;
			}, []);
		},
		getSpellsSummary() {
			if (!this.sorted?.PlayerXml?.Spells) return {};
			let res = {};
			for (let i = 0; i < this.sorted?.PlayerXml?.Spells?.length; i++) {
				let id = this.sorted?.PlayerXml?.Spells?.[i]?.$?.id;
				res[id] = {
					$: {img: this.sorted?.PlayerXml?.Spells?.[i]?.$?.img, id},
					count: res[id]?.count == undefined ? 1 : (res[id].count += 1),
					//PlayerXml: this.sorted.PlayerXml.Spells[i]
				};
			}
			res = Object.entries(res)
				.sort((a, b) => b[1].count - a[1].count)
				.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
			return res;
		},
		getPerksSummary() {
			if (!this.sorted?.WorldStateXml?.lua_globals?.PICKUP_COUNT) return {};
			let res = {};
			for (const [key, value] of Object.entries(this.sorted?.WorldStateXml?.lua_globals?.PICKUP_COUNT)) {
				let id = key.replaceAll(`PERK_PICKED_`, ``).replaceAll(`_PICKUP_COUNT`, ``);
				//console.log('id', id, this.LuaData.perk_list[id])
				res[id] = {
					lua: this.LuaData?.perk_list?.[id],
					$: {
						id,
						img: this.LuaData?.perk_list?.[id]?.ui_icon, //`data/ui_gfx/perk_icons/${id.toLowerCase()}.png`,
					},
					count: value,
				};
			}
			res = Object.entries(res)
				.sort((a, b) => b[1].count - a[1].count)
				.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
			return res;
		},
		fetchFromList:
			(state) =>
			(key = '', obj = state.$state.DataEntityList) => {
				if (key == '') return [null, 'Enter key as first parameter'];
				return obj[key];
			},
		structureData:
			(state) =>
			(obj = state.$state.DataEntityList) => {
				let res = {};
				let filepath = null;
				if (!filepath) filepath = obj?.ItemComponent?.[0]?.['ui_sprite']; // data/ui_gfx/items/potion.png
				if (!filepath) filepath = obj?.SpriteComponent?.[0]?.['image_file']; // data/items_gfx/potion.png
				if (!filepath) filepath = obj?.SpriteComponent?.[0]?.['image_file'];
				if (!filepath) filepath = obj?.UIIconComponent?.[0]?.['icon_sprite_file'];

				if (filepath?.includes('.xml')) filepath = filepath.replace('.xml', '.png');
				res['img'] = filepath;

				if (obj?.GameEffectComponent?.[0]) {
					if (obj?.GameEffectComponent?.[0]?.effect) res['id'] = obj?.ItemComponent?.[0]?.item_name; // ID
				}
				if (obj?.ItemComponent?.[0]) {
					if (obj?.ItemComponent?.[0]?.item_name) res['id'] = obj?.ItemComponent?.[0]?.item_name; // ID

					if (obj?.ItemComponent?.[0]?.['inventory_slot.x']) res['x'] = obj?.ItemComponent?.[0]?.['inventory_slot.x'];
					if (obj?.ItemComponent?.[0]?.['inventory_slot.y']) res['y'] = obj?.ItemComponent?.[0]?.['inventory_slot.y'];

					if (obj?.ItemComponent?.[0]?.permanently_attached) res['permanently_attached'] = obj?.ItemComponent?.[0]?.permanently_attached;
					if (obj?.ItemComponent?.[0]?.uses_remaining) res['uses_remaining'] = obj?.ItemComponent?.[0]?.uses_remaining;
				}
				if (obj?.ItemActionComponent?.[0]) {
					if (obj?.ItemActionComponent?.[0]?.action_id) res['id'] = obj?.ItemActionComponent?.[0]?.action_id; // ID
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
				if (obj?.PhysicsBodyComponent?.[0]) {
					res['hax_wait_till_pixel_scenes_loaded'] = obj.PhysicsBodyComponent?.[0]?.hax_wait_till_pixel_scenes_loaded;
				}
				return res;
			},
		PlayerSummary(state) {
			let obj = {
				'Max hp': parseInt(state.raw?.PlayerXml?.DamageModelComponent?.[0]?.max_hp) < 2 ? `${state.raw?.PlayerXml?.DamageModelComponent?.[0]?.max_hp?.toString()} x 25` : parseInt(state.raw?.PlayerXml?.DamageModelComponent?.[0]?.max_hp) * 25,
				Money: state?.raw?.PlayerXml?.WalletComponent?.[0]?.money,
				'Money spent': state?.raw?.PlayerXml?.WalletComponent?.[0]?.money_spent,
				Position: `x: ${Math.round(state?.raw?.PlayerXml?._Transform?.[0]?.['position.x'])},  y: ${Math.round(state?.raw?.PlayerXml?._Transform?.[0]?.['position.y'])}`,
				World: Math.round(state?.raw?.PlayerXml?._Transform?.[0]?.['position.x'] / 35840),
			};
			return obj;
		},
		WorldStateSummary(state) {
			let obj = {};
			let defaultConf = {includeEmpty: true, includeUndefined: true};
			utils.setObjectValue({...defaultConf, obj: state.raw?.WorldStateXml?.WorldStateComponent?.[0], key: 'mods_have_been_active_during_this_run', resObj: obj, resKey: 'Mods have been active', type: 'integer'});
			utils.setObjectValue({...defaultConf, obj: state.raw?.WorldStateXml?.WorldStateComponent?.[0], key: 'player_polymorph_count', resObj: obj, resKey: 'Polymorph count', type: 'integer'});
			utils.setObjectValue({...defaultConf, obj: state.raw?.WorldStateXml?.WorldStateComponent?.[0], key: 'player_polymorph_random_count', resObj: obj, resKey: 'Polymorph random count', type: 'integer'});
			utils.setObjectValue({...defaultConf, obj: state.raw?.WorldStateXml?.WorldStateComponent?.[0], key: 'day_count', resObj: obj, resKey: 'Day count', type: 'integer'});
			utils.setObjectValue({...defaultConf, obj: state.sorted?.WorldStateXml?.lua_globals, key: 'TEMPLE_ACTIVE', resObj: obj, resKey: 'Active temples', type: 'string', func: (x) => Object.entries(x).reduce((a, [key, value]) => a + parseInt(value), 0) + ` / ${Object.keys(x).length}`});
			utils.setObjectValue({...defaultConf, obj: state.raw?.WorldStateXml?.WorldStateComponent?.[0]?.orbs_found_thisrun?.[0], key: 'primitive', resObj: obj, resKey: 'Orbs', type: 'integer'});
			utils.setObjectValue({...defaultConf, obj: state.sorted?.WorldStateXml?.lua_globals, key: 'ORB_MAP_STRING', resObj: obj, resKey: 'Orb locations', type: 'string', func: (x) => x.split(' ')});
			utils.setObjectValue({...defaultConf, obj: state.sorted?.WorldStateXml?.lua_globals?.COUNT, key: 'TEMPLE_PERK_REROLL_COUNT', resObj: obj, resKey: 'Reroll count', type: 'integer'});
			utils.setObjectValue({...defaultConf, obj: state.sorted?.WorldStateXml?.lua_globals?.COUNT, key: 'TEMPLE_PERK_REROLL_COUNT', resObj: obj, resKey: 'Rerolls left until break', type: 'integer', func: (x) => 39 - parseInt(x), includeUndefined: false});
			utils.setObjectValue({...defaultConf, obj: state.sorted?.WorldStateXml?.lua_globals?.COUNT, key: 'TEMPLE_PERK_REROLL_COUNT', resObj: obj, resKey: 'Rerolls left until breaking the break', type: 'integer', func: (x) => 1017 - parseInt(x), includeUndefined: false});

			return obj;
		},
		SessionSummary:
			(state) =>
			(gameid = state.raw.Gameid, resObj = {}) => {
				resObj['player_projectile_count'] = state.raw.Sessions?.[gameid]?.player_projectile_count;
				resObj['kills'] = state.raw.Sessions?.[gameid]?.kills;
				resObj['player_kills'] = state.raw.Sessions?.[gameid]?.player_kills;
				resObj['deaths'] = state.raw.Sessions?.[gameid]?.deaths;
				resObj['playtime_str'] = state.raw.Sessions?.[gameid]?.stats?.[0]?.playtime_str;
				resObj['world_seed'] = state.raw.Sessions?.[gameid]?.stats?.[0]?.world_seed;
				resObj['hp'] = state.raw.Sessions?.[gameid]?.stats?.[0]?.hp;
				resObj['enemies_killed'] = state.raw.Sessions?.[gameid]?.stats?.[0]?.enemies_killed;
				resObj['damage_taken'] = state.raw.Sessions?.[gameid]?.stats?.[0]?.damage_taken;
				resObj['projectiles_shot'] = state.raw.Sessions[gameid]?.stats?.[0].projectiles_shot;
				resObj['death_pos.x'] = state.raw.Sessions[gameid]?.stats?.[0]?.['death_pos.x'];
				resObj['death_pos.y'] = state.raw.Sessions[gameid]?.stats?.[0]?.['death_pos.y'];
				resObj['killed_by'] = state.raw.Sessions[gameid]?.stats?.[0].killed_by;
				resObj['death_map'] = state.raw.Sessions?.[gameid]?.death_map?.[0]?.E;
				resObj['kill_map'] = state.raw.Sessions?.[gameid]?.kill_map?.[0]?.E;
				return resObj;
			},
		GetEntityData: (state) => (multiplier) => state.counter * multiplier,
		EntityListLifeTimeAffecting: computed(() => {
			return Object.entries(this.EntityList).reduce((acc, [key, value]) => {
				if (!acc[key]) {
					acc[key] = value;
				}
				acc[key] = value;
				return acc;
			}, {});
		}),
		InventoryFlat() {
			if (!this.sorted?.PlayerXml?.Wands) return [];
			let arr = [];
			arr = [...this.sorted.PlayerXml?.Wands?.map((x) => x), ...this.sorted.PlayerXml?.Wands?.map((x) => (x.Spells ? x.Spells.map((y) => y) : [])).flat(), ...this.sorted.PlayerXml?.Items?.map((x) => x), ...this.sorted.PlayerXml?.Spells?.map((x) => x), ...this.sorted.PlayerXml?.Perks?.map((x) => x)];
			return arr;
		},
		summarySpells() {
			let resObj = {};
			for (let i = 0; i < this.InventoryFlat.length; i++) {
				let item = this.InventoryFlat?.[i];
				let action_id = item?.ItemActionComponent?.[0]?.action_id;
				if (action_id) {
					if (!resObj[action_id]) resObj[action_id] = item;
					if (!resObj[action_id]?.$) resObj[action_id].$ = {};
					if (!resObj[action_id]?.$?.count) resObj[action_id].$.count = 0;
					if (!resObj[action_id]?.$?.type) resObj[action_id].$.type = 'spell';
					resObj[action_id].$.count += 1;
				}
			}
			return resObj;
		},
		summaryPerks() {
			if (!this.sorted?.WorldStateXml?.lua_globals?.PICKUP_COUNT) return {};
			let res = {};
			for (const [key, value] of Object.entries(this.sorted?.WorldStateXml?.lua_globals?.PICKUP_COUNT)) {
				let id = key.replaceAll(`PERK_PICKED_`, ``).replaceAll(`_PICKUP_COUNT`, ``);
				//console.log('id', id, this.LuaData.perk_list[id])
				res[id] = {
					lua: this.LuaData?.perk_list?.[id],
					$: {
						id,
						img: this.LuaData?.perk_list?.[id]?.ui_icon, //`data/ui_gfx/perk_icons/${id.toLowerCase()}.png`,
					},
					count: value,
				};
			}
			res = Object.entries(res)
				.sort((a, b) => b[1].count - a[1].count)
				.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
			return res;
		},
		sessionList(state) {
			if (!state.raw?.Sessions) return {};
			let res = {};
			let defaultConf = {};
			for (const [key, value] of Object.entries(state.raw?.Sessions)) {
				res[key] = {};
				res[key].id = key;
				utils.setObjectValue({...defaultConf, obj: value, key: 'deaths', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value, key: 'kills', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value, key: 'player_kills', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value, key: 'player_projectile_count', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value, key: 'BUILD_NAME', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'biomes_visited_with_wands', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'damage_taken', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'dead', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'death_count', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'death_pos.x', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'death_pos.y', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'enemies_killed', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'killed_by', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'gold', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'gold_all', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'gold_infinite', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'healed', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'heart_containers', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'hp', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'items', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'kicks', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'places_visited', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'playtime', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'playtime_str', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'projectiles_shot', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'streaks', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'teleports', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'wands_edited', resObj: res[key]});
				utils.setObjectValue({...defaultConf, obj: value.stats?.[0], key: 'world_seed', resObj: res[key]});
			}
			//res =  Object.entries(res).sort((a,b)=> b[1].count - a[1].count ).reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
			return res;
		},

		killedBy() {
			let res = {deathCount: [0, 0], countPerEnemy: {}, countPerType: {}};
			if (!dataStore()?.raw?.Sessions) return res;
			for (const [key, value] of Object.entries(dataStore().raw.Sessions)) {
				if (!value.stats) continue;
				res.deathCount[0] += parseInt(value.stats[0].dead);
				res.deathCount[1] += 1;

				if (!value.stats?.[0]?.killed_by) continue;
				let [enemy, attack] = value.stats[0].killed_by.split('|');
				res.countPerEnemy[enemy] = res.countPerEnemy[enemy] == undefined ? 1 : (res.countPerEnemy[enemy] = res.countPerEnemy[enemy] + 1);
				res.countPerType[attack] = res.countPerType[attack] == undefined ? 1 : (res.countPerType[attack] = res.countPerType[attack] + 1);
			}
			res.deathCount = res.deathCount.join(' / ');
			res.countPerEnemy = Object.entries(res.countPerEnemy)
				.sort((a, b) => b[1] - a[1])
				.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
			res.countPerType = Object.entries(res.countPerType)
				.sort((a, b) => b[1] - a[1])
				.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
			return res;
		},

		getSortedSessionsList() {
			let sortBy = '-' + storeStore().config?.state?.sidemenu?.sort_by;
			if (!sortBy) console.log('getSortedSessionsList !sortBy', sortBy);
			const arr = Object.values(dataStore().raw.Sessions).sort((a, b) => {
				if (!a['$']) return [];
				if (!storeStore().config?.state?.sidemenu?.sort_by) sortBy = '-datetime';
				return a['$'][sortBy]['place'][0] - b['$'][sortBy]['place'][0];
			});
			return arr;
		},
	},

	actions: {
		async unpackAssets() {
			let res = await app.UnpackAssets();
			console.log('RES', res);
			parse.parseLua.luaInit();
		},
		init() {},
		setDataEntityList(obj) {
			this.DataEntityListHolder = obj;
			localStorage.setItem('DataEntityList', JSON.stringify(this.DataEntityListHolder));
		},
		getEntityData(entity, obj = {}) {
			//obj['img'] = await getImage(entity)
			obj['x'] = entity?.ItemComponent?.['-inventory_slot.x'];
			obj['y'] = entity?.ItemComponent?.['-inventory_slot.y'];
			obj['uses_remaining'] = entity?.ItemComponent?.['-uses_remaining'];
			obj['permanently_attached'] = entity?.ItemComponent?.['-permanently_attached'];
			console.log('getEntityData', obj);
			return obj;
		},

		async updateData(res = {}) {
			//console.log("updateData", res)
			this.raw.lastSync = await utils.currentDateTime();
			if (res.Gameid == '') this.raw.Gameid = '';
			if (res.PlayerXml) this.raw.PlayerXml = {};
			if (res.WorldStateXml) this.raw.WorldStateXml = {};
			if (res.EntityTemplate) {
			}
			collectTime('SessionJson', 'start');
			if (res.SessionJson) dataStore().SessionJson = {...dataStore().SessionJson, ...res.SessionJson};
			collectTime('SessionJson', 'done');

			collectTime('EntityList', 'start');
			if (res.EntityList) dataStore().EntityList = res.EntityList;
			collectTime('EntityList', 'done');

			collectTime('EntityList', 'start');
			if (res.DataEntityList) dataStore().setDataEntityList(res.DataEntityList);
			collectTime('EntityList', 'done');

			collectTime('AvailabeFiles', 'start');
			if (res.AvailabeFiles) this.raw.AvailabeFiles = res.AvailabeFiles; //allData.noitaDataRaw.AvailabeFiles = res.AvailabeFiles
			collectTime('AvailabeFiles', 'done');
			collectTime('AvailabeFilesSingle', 'start');
			if (res.AvailabeFilesSingle) this.raw.AvailabeFiles[Object.entries(res.AvailabeFilesSingle)[0][0]] = Object.entries(res.AvailabeFilesSingle)[0][1]; //allData.noitaDataRaw.AvailabeFiles = res.AvailabeFiles
			collectTime('AvailabeFilesSingle', 'done');

			///// LastDailyRunPlayed
			collectTime('LastDailyRunPlayed', 'start');
			if (res.LastDailyRunPlayed) this.raw.LastDailyRunPlayed = res.LastDailyRunPlayed;
			collectTime('LastDailyRunPlayed', 'done');

			///// Gameid
			if (res.Gameid) this.raw.Gameid = res.Gameid;
			if (res.Gameid) storeStore().state.selected = res.Gameid;

			////// PlayerXml
			collectTime('parsePlayerXml', 'start');
			if (res.PlayerXml) await parse.parsedata.parsePlayerXml(res.PlayerXml, this.raw);
			collectTime('parsePlayerXml', 'done');

			////// WorldStateXml
			collectTime('parseWorldStateXml', 'start');
			if (res.WorldStateXml) await parse.parsedata.parseWorldStateXml(res.WorldStateXml, this.raw);
			collectTime('parseWorldStateXml', 'done');

			////// StatsXml
			collectTime('parseStatsXml', 'start');
			if (res.StatsXml) await parse.parsedata.parseStatsXml(res.StatsXml, this.raw);
			collectTime('parseStatsXml', 'done');

			////// SessionsXml
			collectTime('parseSessionsXml', 'start');
			if (res.SessionsXml) await parse.parsedata.parseSessionsXml(res.SessionsXml, this.raw);
			collectTime('parseSessionsXml', 'done');

			collectTime('config', 'start');
			if (res.Config) this.updateConfig(res.Config); //Object.assign(storeStore().config, res.Config)
			collectTime('config', 'done');

			collectTime('Templates', 'start');
			if (res.Templates) {
				storeStore().templates = res.Templates;
			}
			collectTime('Templates', 'done');
			storeStore().addMsg({event: 'performance', text: ``, json: storeStore().latestCollect});
			storeStore().latestCollect = {};
		},
		async updateConfig(config) {
			console.log('updateConfig', config);
			storeStore().config = config;
			if (config?.highlight) storeStore().highlight = config?.highlight;
			if (storeStore().config?.session_list?.display_filestatus) storeStore().config.session_list.display_filestatus = JSON.parse(config.session_list.display_filestatus);
			if (storeStore().config?.session_list?.display_extra) storeStore().config.session_list.display_extra = JSON.parse(config.session_list.display_extra);
			//Object.assign(storeStore().config, config)
			if (config.window?.window_title) runtime.WindowSetTitle(config.window.window_title);
			if (config.window?.pos_x && config.window?.pos_y) {
				runtime.WindowSetPosition(config.window.pos_x, config.window.pos_y);
			}
			if (config.window?.size_x && config.window?.size_y) {
				runtime.WindowSetSize(config.window.size_x, config.window.size_y);
			}
			storeStore().state.configLoaded = true
			storeStore().addMsg({event: '*', text: `highlight`, json: storeStore().highlight});
		},
		parseTemplates(templates) {
			for (const [key, value] of Object.entries(templates)) {
				//parse.parsetemplates.init(key, value, this.$state)
			}
		},
	},
});
