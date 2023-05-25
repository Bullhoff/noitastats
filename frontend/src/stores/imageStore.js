import {ref, computed, reactive, watch, onMounted, nextTick} from 'vue';
import {defineStore} from 'pinia';
import _ from 'lodash';
import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import * as utils from './../scripts/utils.js';

import {dataStore} from './store.js';

export const imageStore = defineStore('images', () => {
	const test1 = ref('INIT VAL');
	const test1Computed = computed(() => test1.value);
	const images = ref({});
	const imagesComputed = computed(() => images.value);

	const imagesLoaded = computed(() => {
		let obj = {};
		for (const [key, value] of Object.entries(images.value)) {
			obj[key] = new Image(10, 10);
			obj[key].src = value;
		}
		return obj;
	});
	function setImages(obj) {
		console.log('setImages', obj);
		for (const [key, value] of Object.entries(obj)) {
			//localStorage.setItem(`${key}`, value)
		}
	}
	const getImages = computed(() => {
		if (!images.value) return null;
		return images.value;
	});
	const getImage = computed(() => (urlid) => {
		if (!urlid) return '';
		let img = images.value?.[urlid];
		if (!img) {
			getImageFromGo(urlid);
			return '';
		}
		return img;
	});
	const getImageFromGo = async (urlid) => {
		let img = localStorage.getItem(urlid);
		if (img) {
			images.value[urlid] = img;
			return;
		}
		let res = await app.GetImage(urlid);
		if (res) {
			images.value[urlid] = res;
			localStorage.setItem(urlid, res);
		}
		return;
	};

	return {
		setImages,
		getImages,
		getImage,
		images,
		imagesComputed,
		test1,
		test1Computed,
		imagesLoaded,
	};
});
