<template>
	<div style="position: unset; text-align: left; box-sizing: border-box; white-space: nowrap" :style="style.container">
		<Table @onChange="onChange" :obj="list.keyValues" :containerRef="containerRef" />
		<div v-if="nested" v-for="(value, key, index) in list.objects" :key="[...path, key].join('__')">
			<template v-if="false"></template>
			<template v-else>
				<Elements placement="left" rows="objects" :elements="elements" :obj="obj" :currentKey="key" :currentValue="value" @emit-handler="(item) => item.onClick({e: $event, value, key, obj: obj[key]})" />

				<Button @click="buttonExpand(key, $event)" :text="[String.fromCodePoint('0x276F'), String.fromCodePoint('0x276E')]" :active="treeState[key]?.expanded" :onDisabled="list?.objects?.[key] && Object.keys(list.objects?.[key])?.length == 0 ? true : false" />
				<button
					v-if="false"
					class="prevent-select"
					style="width: 3ch; text-align: center"
					@click="buttonExpand(key, $event)"
					:ref="(el) => (refs[key] = el)"
					:style="[list?.objects?.[key] && Object.keys(list.objects?.[key])?.length == 0 ? {disabled: 'true', opacity: '0.3'} : treeState[key]?.expanded ? {backgroundColor: 'darkcyan'} : {backgroundColor: 'gray'}]"
				>
					{{ list?.objects?.[key] && Object.keys(list.objects?.[key])?.length == 0 ? String.fromCodePoint('0x276F') : treeState[key]?.expanded ? String.fromCodePoint('0x276E') : String.fromCodePoint('0x276F') }}
				</button>
				<span style="vertical-align: middle">
					{{ objtitle && Array.isArray(objtitle) && arrayIndexOf(objtitle, value) ? arrayIndexOf(objtitle, value) : objtitle && value?.[objtitle] ? value[objtitle] : key }}
				</span>
				<Elements placement="right" rows="objects" :elements="elements" :obj="obj" :currentKey="key" :currentValue="value" @emit-handler="(item) => item.onClick({e: $event, value, key, obj: obj[key]})" />

				<div v-if="treeState[key]?.expanded" :style="{paddingLeft: `${8 * 3}px`}">
					<TreeView
						@onChange="onChange"
						:obj="typeof value == 'object' ? value : [value]"
						:path="[...path, key]"
						@emit-handler="emitHandler"
						:treeState="treeState[key]"
						:prop="prop"
						:objtitle="objtitle"
						:elements="elements"
						:containerRef="containerRef"
						:editableKeys="editableKeys"
						:editableValues="editableValues"
					/>
				</div>
			</template>
		</div>
	</div>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed} from 'vue';
import {storeToRefs} from 'pinia';
import {storeStore, contextMenuStore} from '../../stores/store.js';

import TreeView from './TreeView.vue';
import Elements from './Elements.vue';
import Table from './Table.vue';

const emit = defineEmits(['emit-handler', 'onChange']);

const props = defineProps(
	reactive({
		containerRef: {default: null},
		elements: {default: []},
		nested: {default: true},
		objtitle: {default: false},
		editableKeys: {default: false},
		editableValues: {default: false},
		treeState: {
			type: Object,
			required: false,
			default: {expanded: false},
		},
		prop: {
			type: Object,
			required: false,
			default: {
				state: {},
			},
		},
		obj: {
			type: Object,
			required: false,
		},
		treeSettings: {
			type: Object,
			required: false,
			default: {},
		},
		parent: {
			type: String,
			required: false,
			//default:
		},
		path: {
			type: Array,
			required: false,
			default: [],
		},
		style: {
			type: Object,
			required: false,
			default: {},
		},
	})
);

const styleDefault = {
	table: {
		cellspacing: '30',
		cellpadding: '30',
	},
	container: {
		'max-width': '100%',
	},
	key: {
		'white-space': 'wrap',
		'padding-right': '1ch',
		border: '1px solid rgb(79, 79, 79)',
		width: 'fit-content',
		'min-width': 'fit-content',
	},
	textareaKey: {
		color: 'rgb(100,100,250)',
		resize: 'none',
		'font-weight': 'bold',
		width: 'fit-content',
		'white-space': 'wrap',
		resize: 'vertical',
		'min-width': 'fit-content',
	},
	value: {
		border: '1px solid rgb(79, 79, 79)',
		width: '*',
		'white-space': 'wrap',
	},
	textareaValue: {
		color: 'rgb(200,200,100)',
		resize: 'vertical',
	},
	textarea: {
		'font-size': '16px',
		overflow: 'hidden',
		'vertical-align': 'middle',
		'line-height': '15.999px',
		'word-break': 'break-all',
		width: '100%',
		'background-color': 'transparent',
		border: 'none',
		outline: 'none',
		'-webkit-box-shadow': 'none',
		'-moz-box-shadow': 'none',
		'box-shadow': 'none',
		resize: 'none',
	},
};

var pathLocal = reactive([...props.path]);
const treeState = reactive(props.treeState);
const refs = reactive({});

function rectLeft(r) {
	let rect = refs.values?.getBoundingClientRect();
	return rect?.left;
}
const style = computed(() => {
	let obj = {...styleDefault, ...props.style};
	if (obj.value['width'] == '*' && rectLeft(refs.values)) obj.value['width'] = window.innerWidth - (rectLeft(refs.values) + 20) + 'px';
	return obj;
});

function onChange(obj) {
	emit('onChange', obj);
}
function emitHandler(e, key, value, path) {
	emit('emit-handler', e, key, value, path);
}

const arrayIndexOf = computed(() => (arr, val) => {
	if (!val) return false;
	if (typeof val != 'object') return false;
	for (const [key, value] of Object.entries(val)) {
		let index = arr.indexOf(key);
		if (index != -1) return value;
	}
	return false;
});

const list = computed(() => {
	let obj = {objects: {}, keyValues: {}, columns: {}};
	if (!props.obj) return obj;
	for (var [key, value] of Object.entries(props.obj)) {
		if (true && Array.isArray(value) && value.length == 1) {
			obj.objects[key] = value[0];
			continue;
		}
		if (typeof value == 'object') {
			obj.objects[key] = value;
		} else {
			obj.keyValues[key] = value;
		}
	}
	return obj;
});

function checkIfHightlight(key, value) {
	for (let i = 0; i < storeStore().config.highlight.length; i++) {
		if (storeStore().config.highlight[i].active && (String(key).toLowerCase().includes(String(storeStore().config.highlight[i].string).toLowerCase()) || String(value).toLowerCase().includes(String(storeStore().config.highlight[i].string).toLowerCase())))
			return storeStore().config.highlight[i].color;
	}
	return 'initial';
}

function contextMenu(e, key) {
	let items = reactive({});
	items['$Hightlight'] = {
		type: 'keyvalue',
		placeholder: {key: 'Search string', value: 'Highlight color'},
		onClick: (e) => {
			console.log('UNUSED KEY', e);
		},
	};
	items['Hightlight'] = [];
	for (let i = 0; i < storeStore().highlight.length; i++) {
		items['Hightlight'].push({
			active: storeStore().highlight[i].active,
			key: storeStore().highlight[i].string,
			value: storeStore().highlight[i].color,
			onChange: (obj) => {
				let {key, value, element} = obj;
				if (element == 'active') storeStore().highlight[key].active = value;
				if (element == 'key') storeStore().highlight[key].string = value;
				if (element == 'value') storeStore().highlight[key].color = value;
				if (element == 'add') storeStore().highlight.push({active: false, string: '', color: ''});
				if (element == 'remove') storeStore().highlight.slice(i, i + 1);
			},
		});
	}
	contextMenuStore().onContextMenu(e, items);
}

function buttonExpand(key, e) {
	treeStateDefault(key);
	treeState[key].expanded = !treeState[key].expanded;
}

onBeforeMount(() => {
	if (!props.obj) return;
	for (const [key, value] of Object.entries(props.obj)) {
		try {
			props.obj[key] = JSON.parse(value);
		} catch (e) {}
		if (typeof value == 'object') {
			treeStateDefault(key);
		}
	}
});
onMounted(() => {
	if (!props.obj) return;
	if (pathLocal && pathLocal.length >= 2) {
		let last = pathLocal.pop();
		pathLocal.push(Number.isInteger(last) ? {type: 'array', key: pathLocal.at(-1)} : last);
	}
});
onUpdated(() => {
	if (!props.obj) return;
	for (const [key, value] of Object.entries(props.obj)) {
		if (typeof value == 'object') treeStateDefault(key);
	}
});

function treeStateDefault(key) {
	if (!treeState[key]) treeState[key] = {expanded: false};
}
</script>

<style scoped>
table {
	border-collapse: collapse;
}
tr {
	width: 100%;
}
table,
th,
td {
	text-align: left;
}

.highlight {
	background-color: aqua;
}

.value-changed {
	background-color: bisque;
}

input {
	background-color: gray;
	max-width: 100%;
}

li,
ul,
button {
	text-align: left;
	white-space: nowrap;
}
button {
	background-color: gray;
}
.button-expanded {
	background-color: darkcyan;
}
</style>
