import {ref, computed, reactive, watch, onMounted, nextTick} from 'vue';
import {defineStore} from 'pinia';
import {storeToRefs} from 'pinia';
//import {windowStore} from './store.js'
import _ from 'lodash';
import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import * as utils from './../scripts/utils.js';

export const contextMenuStore = defineStore('contextMenu', () => {
	function clearItems(obj = items) {
		for (const [key, value] of Object.entries(obj)) {
			delete obj[key];
		}
		for (const [key, value] of Object.entries(config.value)) {
			delete config.value[key];
		}
	}
	async function onContextMenu(e, newItems, rect, conf) {
		//console.log('onContextMenu', conf)
		//let {_rect, } = conf
		await clearItems();
		await nextTick();
		//items.value = newItems
		//items = newItems
		//onContextMenu().items = newItems
		Object.assign(items, newItems);
		//console.log('onContextMenu', e.clientX, e.clientY, e)
		//prop['x'] = e.clientX + 'px'
		//prop['y'] = e.clientY + 'px'
		if (conf) _.merge(config.value, conf);
		if (conf?.rect) {
			config.value['x'] = rect.left + rect.width / 2 + 'px';
			config.value['y'] = rect.top + rect.height + 'px';
			pos['x'] = rect.left + rect.width + 'px';
			pos['y'] = rect.top + rect.height + 'px';
			open.value = !open.value;
		} else if (rect) {
			pos['x'] = rect.left + rect.width + 'px';
			pos['y'] = rect.top + rect.height + 'px';
			open.value = !open.value;
		} else {
			pos['x'] = e.clientX + 'px';
			pos['y'] = e.clientY + 'px';
			open.value = true;
		}
		//contextMenuStore().open =! contextMenuStore().open
		//open.value =! open.value

		//console.log('onContextMenu', open.value, pos, items)
	}
	const config = ref({});
	const items = reactive({initItem1: () => console.log('waa')});
	const divRef = ref(null);
	const prop = reactive({});
	const pos = reactive({});
	const open = ref(false);
	return {
		onContextMenu,
		config,
		open,
		divRef,
		items: items,
		//items: items.value,
		prop,
		pos,
		focused: reactive({}),
	};
});
