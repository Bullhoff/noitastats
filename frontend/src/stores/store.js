import {ref, computed, reactive, watch, onMounted, nextTick} from 'vue';
import {defineStore} from 'pinia';
//import { storeToRefs } from 'pinia'
import _ from 'lodash';
import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import * as utils from './../scripts/utils.js';

import {windowStore} from './windowStore.js';
import {lifetimeStore} from './lifetimeStore.js';
import {contextMenuStore} from './contextMenuStore.js';
import {dataStore} from './dataStore.js';
import {imageStore} from './imageStore.js';
import {refStore} from './refStore.js';
export {windowStore, lifetimeStore, contextMenuStore, dataStore, imageStore, refStore};

export const storeStore = defineStore('store', () => {
	const state = reactive({
		loading: {Stats: false, Sessions: false, Player: false, WorldState: false},
		configLoaded: false, 
		finisedLoading: false,
		selected: '',
		sidemenuWidth: 0,
		alwaysOnTop: false,
		dropdown: false,
	});
	const msg = reactive({
		status: [],
	});
	function addMsg() {
		let event = arguments[0]?.event;
		if (event) {
			let json = {};
			if (msg.json) json = {...JSON.parse(JSON.stringify(msg.json))};
			if (!msg[event]) Object.assign(msg, {[event]: []});
			msg[event].push({datetime: utils.currentDateTime(), json, ...arguments[0]});
		}
	}
	const config = ref({
		state: {},
	});
	const templates = reactive({});
	const latestCollect = reactive({});
	
	const cursor = computed(()=>{
		if(state.configLoaded) return 'auto'
		return 'wait'
	})
	return {
		cursor,
		//cursor: ref('auto'),
		/* configComputed: computed(()=>{
			return config.value
		}), */
		configState: computed(() => config.value?.state),
		config,
		msg,
		addMsg,
		state,
		templates,
		latestCollect,
		entitiesExpanded: reactive({wands: [false, false, false, false, false, false, false, false, false], pocket: false, backpack: false, perks: false}),
		finishedLoading: ref({}),
	};
});
