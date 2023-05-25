import {ref, computed, reactive, watch, onMounted, nextTick} from 'vue';
import {defineStore} from 'pinia';
//  import { storeToRefs } from 'pinia'
//import * as utils from './../utils.js'
import _ from 'lodash';

import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import * as utils from './../scripts/utils.js';

import {dataStore} from './store.js';

export const lifetimeStore = defineStore('lifetimeStore', {
	state: () => {
		return {
			resultArray: ref([]),
			spellSelected: ref({id: 'BLACK_HOLE', lifetime: 120}),
			allItems: ref({
				PROJECTILE_HOMING_SHOOTER: {
					type: 'perk',
					name: 'Boomerang Spells',
					image: 'data/ui_gfx/perk_icons/projectile_homing_shooter.png',
					lifetime: +20,
					amount: 0,
				},
				BOUNCE: {
					type: 'perk',
					name: 'Bouncing Spells',
					image: 'data/ui_gfx/perk_icons/bounce.png',
					lifetime: +60,
					amount: 0,
				},
				LIFETIME: {
					name: 'Increase Duration',
					image: 'data/ui_gfx/gun_actions/lifetime.png',
					lifetime: +75,
					amount: 0,
				},
				PINGPONG_PATH: {
					name: 'Ping-Pong Path',
					image: 'data/ui_gfx/gun_actions/pingpong_path.png',
					lifetime: +25,
					amount: 0,
				},
				ORBIT_SHOT: {
					name: 'Orbiting Arc',
					image: 'data/ui_gfx/gun_actions/orbit_shot.png',
					lifetime: +25,
					amount: 0,
				},
				SPIRALING_SHOT: {
					name: 'Spiral Arc',
					image: 'data/ui_gfx/gun_actions/spiraling_shot.png',
					lifetime: +50,
					amount: 0,
				},
				PHASING_ARC: {
					name: 'Phasing Arc',
					image: 'data/ui_gfx/gun_actions/phasing_arc.png',
					lifetime: +80,
					amount: 0,
				},
				LIFETIME_DOWN: {
					name: 'Reduce Lifetime',
					image: 'data/ui_gfx/gun_actions/lifetime_down.png',
					lifetime: -42,
					amount: 0,
				},
				CHAIN_SHOT: {
					name: 'Chain Spell',
					image: 'data/ui_gfx/gun_actions/chain_shot.png',
					lifetime: -30,
					amount: 0,
				},
			}),
		};
	},
	getters: {
		InventoryLifeTimeAffecting() {
			return dataStore().InventoryFlat.reduce((acc, value) => {
				if (this.allItems[value.$.id] || (value?.ItemActionComponent?.[0]?.action_id && dataStore().LuaData?.gun_actions?.[value?.ItemActionComponent?.[0]?.action_id] && dataStore().LuaData?.gun_actions?.[value?.ItemActionComponent?.[0]?.action_id]?.lifetime_add)) acc.push(value);
				return acc;
			}, []);
		},
		result() {
			let obj = {
				count: 0,
				count_spells: 0,
				count_perks: 0,
				sum: 0,
			};
			for (const [key, value] of Object.entries(this.allItems)) {
				obj.sum += value.amount * value.lifetime;
				obj.count += value.amount;
				if (value.type == 'perk') obj.count_perks += value.amount;
				else obj.count_spells += value.amount;
			}

			return obj;
		},

		lifetimeSpellDropdown: (state) => {
			let resObj = {};
			let obj = dataStore().LuaData.gun_actions;
			for (var [key, value] of Object.entries(obj)) {
				if (!resObj[key]) resObj[key] = {};
				resObj[key].lua = value;
				if (value.custom_xml_file) {
					if (!resObj[key].custom_xml_file) resObj[key].custom_xml_file = [];
					if (!Array.isArray(value.custom_xml_file)) obj[key].custom_xml_file = [value.custom_xml_file];
					for (let i = 0; i < obj[key].custom_xml_file.length; i++) {
						if (dataStore().DataEntityList[obj[key].custom_xml_file[i]]) resObj[key].custom_xml_file.push(dataStore().DataEntityList[obj[key].custom_xml_file[i]]);
					}
				}

				if (value.related_projectiles) {
					if (!resObj[key].related_projectiles) resObj[key].related_projectiles = [];
					if (!Array.isArray(value.related_projectiles)) obj[key].related_projectiles = [value.related_projectiles];
					for (let i = 0; i < obj[key].related_projectiles.length; i++) {
						if (dataStore().DataEntityList[obj[key].related_projectiles[i]]) resObj[key].related_projectiles.push(dataStore().DataEntityList[obj[key].related_projectiles[i]]);
					}
				}

				if (value.related_extra_entities) {
					if (!resObj[key].related_extra_entities) resObj[key].related_extra_entities = [];
					if (!Array.isArray(value.related_extra_entities)) obj[key].related_extra_entities = [value.related_extra_entities];
					for (let i = 0; i < obj[key].related_extra_entities.length; i++) {
						if (dataStore().DataEntityList[obj[key].related_extra_entities[i]]) resObj[key].related_extra_entities.push(dataStore().DataEntityList[obj[key].related_extra_entities[i]]);
					}
				}
			}
			return resObj;
		},
	},
	actions: {},
});
