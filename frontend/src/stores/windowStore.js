import {ref, computed, reactive, watch, unref, toRaw, markRaw, isRef, isProxy, isReadonly, nextTick} from 'vue';
import {defineStore} from 'pinia';
//import * as utils from './../utils.js'
//import c from './../constants.js'
import _ from 'lodash';

export const windowStore = defineStore('window', () => {
	const defaultWindow = ({name, component, id, title = null}) => ({
		e: null,
		id: id,
		component: component,
		title: title ? title : component,
		style: {zIndex: windows.length + 1},
		w: null,
		h: null,
		pos: {x: '0', y: '0'},
		holding: false,
		minimized: false,
		maximized: false,
	});
	const current = reactive({
		holdingCorner: false,
		activeWindow: null,
		holding: false,
		windowOrder: reactive([]),
		x1: 0,
		x2: 0,
		y1: 0,
		y2: 0,
	});

	const windows = reactive({});
	const holding = computed(() => Object.values(windows).reduce((acc, curr) => (curr.holding ? curr.holding : acc), false));
	const refs = reactive({});
	const props = reactive({});
	const windowListComponentKey = computed(() => {
		if (!windows || Object.keys(windows) == 0) return {};
		let list = {};
		for (const [key, value] of Object.entries(windows)) {
			if (!list[value.component]) list[value.component] = [];
			list[value.component].push(value);
		}
		console.log('windowKeyList', list);
		return list;
	});
	const windowExists = computed(() => {
		if (!windows || Object.keys(windows) == 0) return [];
		let list = Object.values(windows).map((el) => {
			return el.component;
		});
		return list;
	});

	function createNewWindow({name, title, component, style = {}, head = {}, body = {}, prop = null, _props = {}, componentSettings = {}, e} = {}) {
		console.log('createNewWindow', component, prop);
		const id = Date.now();
		current.windowOrder.push(id);
		refs[id] = null;
		windows[id] = reactive(defaultWindow({id, component, title}));
		windows[id].e = e;
		if (!prop) prop = {head: [{}], body: [{}]};
		if (!prop.head) prop['head'] = [{}];
		if (!prop.body) prop['body'] = [{}];
		props[id] = prop;
	}
	function createNewOrUnminimize({component, id, forceNew = false, prop = null, componentSettings = {}} = {}) {
		if (!windowExists.value.includes(component) || forceNew) return createNewWindow({component, prop, componentSettings});
		const minimize = windowListComponentKey.value[component][0].minimized ? false : true;
		for (let index = 0; index < Object.entries(windowListComponentKey.value[component]).length; index++) {
			let id = windowListComponentKey.value[component][index].id;
			focus(id);
			windows[id].minimized = minimize;
		}
	}

	function focus(id) {
		current.activeWindow = id;
		let index = current.windowOrder.indexOf(id);
		current.windowOrder.splice(index, 1);
		current.windowOrder.push(id);
		for (let i = 0; i < current.windowOrder.length; i++) {
			if (windows[current.windowOrder[i]]) windows[current.windowOrder[i]].style.zIndex = i + 1;
			if (refs[current.windowOrder[i]]) refs[current.windowOrder[i]].style.zIndex = i + 1;
		}
	}
	function coords(e, id) {
		let {clientX, clientY} = e;
		let nrtwo = () => {
			current.x2 = clientX;
			current.y2 = clientY;
		};
		current.x1 = current.x2 - clientX;
		current.y1 = current.y2 - clientY;
		nrtwo();
		let x = current.x1;
		let y = current.y1;
		return {x: x, y: y};
	}
	function ensureInside(e, divRef) {
		let minX = 5;
		let minY = 5;
		let titleBarHeight = '2.5ch';
		let clientX = e.clientX;
		let clientY = e.clientY;
		let screenX = e.screenX;
		let screenY = e.screenY;
		let rect = divRef.getBoundingClientRect();
		let winW = window.innerWidth;
		let winH = window.innerHeight;
		if (e.clientX - minX > 0 && e.clientX + minX < window.innerWidth && e.clientY - minY > 0 && e.clientY + minY < window.innerHeight) return true;

		return false;
	}
	let func = {
		mouseUp(id = current.activeWindow) {
			windows[id].holding = false;
			current.holding = false;
			current.holdingCorner = false;
		},
		mouseDownTitle(e, id = null, corner) {
			if (corner) current.holdingCorner = corner;
			windows[id].holding = true;
			current.holding = true;
			focus(id);
			coords(e, id);
		},
		mouseDownContainer(id = current.activeWindow) {
			focus(id);
		},
		mouseDownOutside(id = current.activeWindow) {
			windows[id].holding = false;
			current.holding = false;
			current.holdingCorner = false;
		},
		mouseDown(e, id = null, corner) {
			// nw n ne w e sw s se
			windows[id].holding = true;
			current.holding = true;
			current.holdingCorner = corner;
			focus(id);
			coords(e, id);
		},
		mouseMove(e, id = current.activeWindow, corner) {
			if (!ensureInside(e, refs[id])) return;
			if (current.holdingCorner && id == current.activeWindow) {
				let {x, y} = coords(e, id);
				let rect = refs[id].getBoundingClientRect();

				let obj = {};
				let ofsLeft = refs[id].offsetLeft;
				let ofsTop = refs[id].offsetTop;

				let xtra_1 = 0; //1.6//parseFloat(refs[id].style.height)-rect.height
				let xtra_2 = 3;
				let minW = 100;
				let minH = 100;
				if (current.holdingCorner == 'n' || current.holdingCorner == 'nw' || current.holdingCorner == 'ne') {
					if (rect.bottom < window.innerHeight || (y > 0 && e.clientY < window.innerHeight - minH)) refs[id].style.top = refs[id].offsetTop - y + 0 + 'px';
					let newH = rect.bottom - e.clientY + 0 - xtra_1;
					if (newH >= minH || (y > 0 && rect.bottom > window.innerHeight) || (y > 0 && e.clientY < window.innerHeight - minH)) refs[id].style.height = rect.height + y + 0 - xtra_1 + 'px';
				}
				if (current.holdingCorner == 's' || current.holdingCorner == 'sw' || current.holdingCorner == 'se') {
					if (rect.top + minH > e.clientY && y > 0 && rect.top - y > 0) refs[id].style.top = refs[id].offsetTop - y + 0 - 0.0 + 'px';
					let newH = e.clientY - rect.top + 0 - xtra_1;
					if (newH >= minH || (newH >= minH && y < 0 && e.clientY < window.innerHeight - minH)) refs[id].style.height = e.clientY - rect.top + 0 - xtra_1 + 'px';
				}

				if (current.holdingCorner == 'w' || current.holdingCorner == 'nw' || current.holdingCorner == 'sw') {
					if (rect.right < window.innerWidth || (x > 0 && e.clientX < window.innerWidth - minW)) refs[id].style.left = refs[id].offsetLeft - x + 0 + 'px';
					if (rect.width >= minW || (x > 0 && e.clientX < window.innerWidth - minW)) refs[id].style.width = rect.width + x + 0 - xtra_1 + 'px';
				}
				if (current.holdingCorner == 'e' || current.holdingCorner == 'ne' || current.holdingCorner == 'se') {
					if (rect.left + minW > e.clientX && x > 0 && rect.left - x > 0) refs[id].style.left = refs[id].offsetLeft - x + 0 + 'px';
					let newW = e.clientX - rect.left + 0 - xtra_1;
					if (newW >= minW || (newW >= minW && x < 0 && e.clientX < window.innerWidth - minW)) refs[id].style.width = e.clientX - rect.left + 0 - xtra_1 + 'px';
				}
			}
			if (!current.holding) return false;
			if (id != current.activeWindow) return;
			let {x, y} = coords(e, id);
			windows[id].pos.x = refs[id].offsetLeft - x + 'px';
			windows[id].pos.y = refs[id].offsetTop - y + 'px';
			refs[id].style.left = windows[id].pos.x;
			refs[id].style.top = windows[id].pos.y;
		},

		close(id = current.activeWindow) {
			let index = current.windowOrder.indexOf(id);
			current.windowOrder.splice(index, 1);
			delete refs[id];
			delete windows[id];
			for (let i = 0; i < current.windowOrder.length; i++) {
				if (windows[current.windowOrder[i]]) windows[current.windowOrder[i]].style.zIndex = i + 1;
			}
		},
		minimize(id = current.activeWindow, value) {
			if (value == true || value == false) windows[id].minimized = value;
			else windows[id].minimized = !windows[id].minimized;
		},
		maximize(id = current.activeWindow) {
			if (!windows[id]) return;
			windows[id].maximized = !windows[id].maximized;
			if (windows[id].maximized) {
				refs[id].style.maxWidth = 'none';
				refs[id].style.maxHeight = 'none';
				refs[id].style.left = '0px';
				refs[id].style.top = '0px';
				refs[id].style.width = '100vw';
				refs[id].style.height = '100vh';
			} else {
				refs[id].style.maxWidth = '100vw';
				refs[id].style.maxHeight = '100vh';
				refs[id].style.width = 'fit-content';
				refs[id].style.height = 'fit-content';
				let {x, y} = windows[id].pos;
				refs[id].style.left = `${x}`;
				refs[id].style.top = `${y}`;
			}
		},
		resize(e, id) {
			console.log();
		},
		onMounted(divRef, id) {
			refs[id] = divRef;
			focus(id);
			//new ResizeObserver(()=>{}).observe(refs[id])
			let pos = {
				x: windows[id].e.clientX,
				y: windows[id].e.clientY,
			};
			let width = refs[id].offsetWidth;
			if (windows[id].e?.clientX + width > window?.innerWidth) {
				pos.x = window.innerWidth - width;
			}
			let height = refs[id].offsetHeight;
			if (windows[id].e?.clientY + height > window?.innerHeight) {
				pos.y = window.innerHeight - height;
			}
			refs[id].style.left = `${pos.x}px`;
			refs[id].style.top = `${pos.y}px`;
		},
	};
	var handle = computed(() => {
		let resObj = {};
		for (const [key, value] of Object.entries(windows)) {
			if (value.minimized)
				resObj[key] = {
					text: value.title,
					value: value.minimized,
					type: 'checkbox',
					onChange: (obj) => {
						windows[key].minimized = obj.checked;
					},
				};
		}
		return resObj;
	});
	return {current, refs, windows, createNewWindow, func, windowExists, createNewOrUnminimize, windowListComponentKey, props, holding, handle};
});
