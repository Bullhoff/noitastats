<template>
	<div class="container prevent-select" tabindex="0" style="min-width: 10ch" :ref="(el) => (containerRef = el)" @blur="onFocusOut" @focus="onFocusIn" :style="[style.container ? style.container : {}]">
		<template v-if="config.row && config.row.length > 0">
			<Table :obj="items" :td="config.row" :th="config.th" :containerRef="containerRef" :prop="{}" :columns="config.columns" :table="config.table" />
		</template>
		<div v-else v-for="(value, key, index) in p.items" :key="p.path + '__' + key" style="padding: 0ch 0ch 0ch 0ch; width: 100%" class="row">
			<template v-if="typeof value == 'object' && key[0] == '$'"></template>

			<template v-if="config.row && config.row.length > 0">
				<Row :prop="{...prop, key, value, config}" />
			</template>

			<template v-else-if="typeof value == 'object' && value?.type" :title="props.items['$' + key]?.tooltip">
				<div v-if="value.type == 'keyvalue'" style="position: relative"><!-- class="enable-select" --></div>
				<div v-if="value.type == 'checkbox'" style="white-space: nowrap">
					<Clickable>
						<input :ref="key" :id="key" type="checkbox" :checked="p.items[key].value" v-model="p.items[key].value" @mousedown.stop.prevent="" @change="value.onChange({...prop, e: event, checked: $event.target.checked, key, value, v: value.value})" />
						<label :for="key" style="width: 100%">{{ value.text != undefined ? value.text : key }}&#xa0;</label>
					</Clickable>
				</div>

				<div v-else-if="value.type == 'enlarged'" class="clickable">
					<p :style="{fontSize: '6ch', textAlign: 'center', ...value.style}" :title="props.items['$' + key]?.tooltip">
						{{ value.value }}
					</p>
				</div>
				<Clickable v-else-if="value.type == 'function'" :prop="prop" :p.path="p.path + '__' + key">
					<Function :prop="{key, value}" />
				</Clickable>

				<div v-else-if="value.type == 'slider'"></div>
				<template v-else-if="value.type == 'active'">
					<div class="clickable" style="">
						<input type="checkbox" :checked="value.active" @focusin="onFocusIn($event, p.path + '__' + key + '_checkbox')" @focusout="onFocusOut($event, p.path + '__' + key + '_checkbox')" @click="value.onChange($event, $event.target.checked)" />
						<p
							@click="
								value.active = !value.active;
								value.onChange($event, $event.target.checked);
							"
						>
							{{ key }}
						</p>
					</div>
				</template>
				<div v-else-if="value.type == 'radioGroup'">
					<Clickable v-for="(value, key, index) in value.array" @focusin="onFocusIn($event, p.path + '__' + key + '_radio')" @focusout="onFocusOut($event, p.path + '__' + key + '_radio')">
						<Radio :prop="{key, value}" />
					</Clickable>
				</div>
				<div v-else-if="value.type == 'radioGroup'">
					<Radio :prop="{key, value}" />
				</div>
				<Buttons
					:prop="{key, value}"
					v-else-if="value.type == 'buttons' && value.array"
					@click="
						(e) => {
							if (e) {
								contextMenuStore().open = false;
							}
						}
					"
				/>
			</template>

			<template v-else-if="p.prop?.type && p.prop.type == 'keyvalue'">
				<div :title="key + '\n' + p.prop.tooltip">
					<input type="checkbox" :checked="value.active" @focusin="onFocusIn($event, p.path + '__' + key + '_checkbox')" @focusout="onFocusOut($event, p.path + '__' + key + '_checkbox')" @click="value.onChange({e: $event, value: $event.target.checked, element: 'active', key})" />
					<input
						type="text"
						:placeholder="p.prop?.placeholder?.key"
						:value="value.key"
						style="position: relative"
						@focusin="onFocusIn($event, p.path + '__' + key + '_key')"
						@focusout="onFocusOut($event, p.path + '__' + key + '_key')"
						@change="
							props.items[key].key = $event.target.value;
							value.onChange({e: $event, value: $event.target.value, element: 'key', key});
						"
					/>
					<input
						type="text"
						:placeholder="p.prop?.placeholder?.value"
						:value="value.value"
						style="position: relative"
						@focusin="onFocusIn($event, p.path + '__' + key + '_value')"
						@focusout="onFocusOut($event, p.path + '__' + key + '_value')"
						@change="
							props.items[key].value = $event.target.value;
							value.onChange({e: $event, value: $event.target.value, element: 'value', key});
						"
					/>
					<button
						@focusin="onFocusIn($event, p.path + '__' + key + '_add')"
						@focusout="onFocusOut($event, p.path + '__' + key + '_add')"
						@click="
							props.items.push({active: false, string: 'text', color: 'yellow', onChange: value.onChange});
							value.onChange({e: $event, value: 'add', element: 'add', key});
						"
					>
						+
					</button>
					<button @focusin="onFocusIn($event, p.path + '__' + key + '_remove')" @focusout="onFocusOut($event, p.path + '__' + key + '_remove')" @click="value.onChange({e: $event, value: 'remove', element: 'remove', key})">-</button>
				</div>
			</template>

			<template v-else-if="p.items['$' + key]?.type">
				<p class="clickable" :ref="(el) => (treeObj[key]['$rowRef'] = el)" @click.stop.prevent="openSubMenu($event, key, value)" :title="props.items['$' + key]?.tooltip" style="display: flex; flex-direction: row; justify-self: right">
					<span style="width: 100%">{{ key }}</span>
					<span style="justify-self: flex-end">&#xa0;🡲</span>
				</p>

				<div style="position: absolute" :style="treeObj[key]?.['$subList']" :ref="(el) => (treeObj[key]['$subListRef'] = el)">
					<div v-if="treeObj[key]?.['$subListOpen']">
						<ContextMenu :items="p.items[key]" @opening-sub="openingSub($event)" :prop="{...p.prop, type: p.items['$' + key].type, ...p.items['$' + key]}" :path="p.path + '__' + key" :init="{}" :containerRect="containerRef.getBoundingClientRect()" />
					</div>
				</div>
			</template>
			<template v-else-if="typeof value == 'string'" style="">
				<Clickable>
					<p title="string" style="width: 100%" :title="props.items['$' + key]?.tooltip">{{ key }}: {{ value }}</p>
				</Clickable>
			</template>
			<div v-else-if="typeof value == 'function'">
				<Clickable @click.prevent="onClick($event, value, key)" :title="props.items['$' + key]?.tooltip">{{ key }} </Clickable>
			</div>

			<template v-else-if="typeof value == 'object'">
				<Clickable @click.stop.prevent="openSubMenu($event, key, value)" :title="props.items['$' + key]?.tooltip" style="display: flex; flex-direction: row">
					<span style="width: 100%">{{ key }}</span>
					<span style="justify-self: flex-end">&#xa0;🡲</span>
				</Clickable>
				<div v-if="treeObj[key]" style="position: absolute" :style="treeObj[key]?.['$subList']" :ref="(el) => (treeObj[key]['$subListRef'] = el)">
					<div v-if="treeObj[key]?.['$subListOpen']" style="padding-left: 3px">
						<ContextMenu :items="p.items[key]" @opening-sub="openingSub($event)" :path="p.path + '__' + key" :init="{}" :containerRect="containerRef.getBoundingClientRect()" />
					</div>
				</div>
			</template>

			<template v-else-if="typeof value == 'boolean'" style="">
				<Clickable>
					<label style="width: 100%; display: flex; align-items: center">
						<Checkbox :checked="value" @emit-handler="emitHandler" @change="" @onChange="emitHandler" />
						<span>{{ key }}</span>
					</label>
				</Clickable>
			</template>

			<template v-else style="">
				<Clickable>
					<p :title="typeof value" style="width: 100%">(UNSORTED) {{ typeof value }} {{ key }}: {{ value }}</p>
				</Clickable>
			</template>
		</div>
	</div>
</template>

<script setup>
import _ from 'lodash';
import {ref, reactive, toRefs, watch, watchEffect, onMounted, onBeforeMount, onUpdated, nextTick, onDeactivated, onRenderTriggered, computed, inject, provide} from 'vue';
import {contextMenuStore} from './../../stores/store.js';
import Clickable from './Clickable.vue';
import Function from './Function.vue';
import Buttons from './Buttons.vue';
import Radio from './Radio.vue';
import Row from './Row.vue';

const emit = defineEmits(['opening-sub', 'emit-handler', 'onChange']);
const props = defineProps(
	reactive({
		prop: {
			type: Object,
			default: {
				type: null,
				focused: {},
				open: false,
			},
		},
		config: {
			type: Object,
			default: {rect: {}, row: []},
		},
		open: {
			type: Boolean,
			default: false,
		},
		init: {
			type: Object,
			default: {$containerRef: null},
		},
		containerRect: {
			type: Object,
			default: {},
		},
		path: {
			type: String,
			default: 'root',
		},
		items: {
			type: Object,
			default: {},
		},
		sideObject: {
			type: Object,
			default: {$containerRef: null},
		},
		itemsSettings: {
			type: Object,
			default: {},
		},
		styles: {
			type: Object,
			default: {},
		},
		divRef: {
			default: null,
		},
	})
);

var defaultStyle = {
	container: {
		'max-width': '70%',
		'max-height': '90%',
		'overflow-x': 'hidden',
		'overflow-y': 'auto',
		//'position': 'absolute',
		position: 'fixed',
		'z-index': '999',
	},
	clickable: {
		'background-color': 'rgba(75, 75, 75, 0.95)',
		display: 'flex',
		'flex-direction': 'row',
		padding: '0ch 0.5ch 0ch 0.5ch',
		color: 'white',
	},
	clickableHover: {
		// activeHover : false
		'background-color': 'blue',
		'border-radius': '3px',
	},
	items: {},
	open: false,
};
var style = reactive({...defaultStyle, ...style});

var prop = reactive(props.prop);
const element = reactive([]);
const refs = reactive({});
const containerRef = ref(null);
const treeObj = reactive({});
const p = reactive(Object.assign({}, props));

provide('p', p);

function emitHandler(obj) {
	emit('onChange', obj);
	emit('emit-handler', obj);
	return obj;
}
function onBlur() {}

var injected = inject('injected', {
	open: true,
	path: [],
	treeObj: {},
	onBlur: () => {},
	onFocus: () => {},
});
provide('config', props.config);

async function onFocusOut(e, path = p.path) {
	injected.onBlur({e: e, path: [...injected.path, ...path.split('__')]});
	if (props.config.onBlur) props.config.onBlur({e: e, path: [...props.config.path, ...path.split('__')]});
	await setTimeout(() => {
		contextMenuStore().focused[path] = false;
		prop.focused[path] = false;
		let open = false;
		for (const [key, value] of Object.entries(contextMenuStore().focused)) {
			if (value == true) open = true;
		}
		props.init.open = open;
		contextMenuStore().open = open;
		prop.open = open;
	}, 10);
	return;
}
function onFocusIn(e, path = p.path) {
	injected.onFocus({e: e, path: [...injected.path, ...path.split('__')]});
	if (props.config.onFocus) props.config.onFocus({e: e, path: [...props.config.path, ...path.split('__')]});
	contextMenuStore().focused[path] = true;
	prop.focused[path] = true;
}
async function openSubMenu(e, key, value) {
	if (treeObj[key]['$subListOpen']) {
		let parent = p.path.split('__').pop();
		await onFocusIn(e, parent);
		treeObj[key]['$subListOpen'] = false;
	} else {
		onFocusIn(e, p.path + '__' + key);
		treeObj[key]['$subListOpen'] = true;
	}
	if (!treeObj[key]['$subListOpen']) return;
	await nextTick();
	let rectContainer = containerRef.value.getBoundingClientRect();
	let rectRow = e.target.getBoundingClientRect();
	let left = containerRef.value.offsetWidth;
	let top = e.target.offsetTop;
	treeObj[key]['$subListRef'].style.left = `${left - 3}px`;
	treeObj[key]['$subListRef'].style.top = `${top}px`;
	let scrollTop = containerRef.value.scrollTop;
	treeObj[key]['$subListRef'].style.top = `${treeObj[key]?.['$subListRef']?.offsetTop - scrollTop}px`;
}

async function onClick(e, value, key) {
	value();
	await nextTick();
	contextMenuStore().open = false;
}

function openingSub(e) {
	for (const [key, value] of Object.entries(treeObj)) {
		if (treeObj && treeObj?.[key]?.['$subListOpen'] && treeObj?.[key]?.['$subListOpen']) {
			if (e != props.path + '__' + key) treeObj[key]['$subListOpen'] = false;
		}
		let newE = e.split('__');
		newE.pop();
		newE = newE.join('__');
		emit('opening-sub', newE);
	}
}

onBeforeMount(async () => {
	treeObj['$containerRef'] = null;
	let index = 0;
	if (!p.items) return;
	for (const [key, value] of Object.entries(p.items)) {
		element[index] = {};
		index += 1;
		p.styles[key] = {};
		p.styles[key]['$row'] = {};
		if (typeof value == 'object') p.styles[key]['$subListOpen'] = false;
		treeObj[key] = {};
		treeObj[key]['$row'] = {};
		treeObj[key]['$rowRef'] = null;
		if (typeof value == 'object') {
			treeObj[key]['$subListOpen'] = false;
			treeObj[key]['$subListRef'] = null;
		}
	}
});

onMounted(async () => {
	if (props.init.x && props.init.y) {
		//await nextTick()
		console.log('containerRef.value', containerRef.value, containerRef);
		containerRef.value.style.left = props.init.x;
		containerRef.value.style.top = props.init.y;
	}
	containerRef.value.focus();
	emit('opening-sub', props.path);

	await nextTick();
	const containerRect = containerRef.value.getBoundingClientRect();

	var paddingX = 10;
	var paddingY = 15;
	if (containerRect.right > window.innerWidth - paddingX) {
		if (!props.containerRect.width) {
			containerRef.value.style.left = `${window.innerWidth - containerRect.width - paddingX}px`;
		} else {
			containerRef.value.style.left = `${props.containerRect.left - containerRect.width}px`;
		}
	}
	if (containerRect.bottom > window.innerHeight - paddingY) {
		containerRef.value.style.top = `${window.innerHeight - containerRect.height - paddingY}px`;
	}
	treeObj['$startOffsetTop'] = containerRef.value.offsetTop;
});
onUpdated(async () => {});
</script>

<style scoped>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}

*:focus {
	outline: none;
}
.container {
	display: block;
	cursor: pointer;
	margin: 0px;
	padding: 0px;
	white-space: nowrap;
	height: fit-content;
	width: fit-content;
}
.row {
	margin: 0px;
	padding: 0px;
	white-space: nowrap;
	text-align: left;
	border-radius: 15px;
}
div {
	margin: 0px;
	padding: 0px;
}
template {
	margin: 0px;
	padding: 0px;
}
</style>
