<script setup>
import {ref, reactive, onMounted, onBeforeMount, onUpdated, nextTick, computed} from 'vue';
//import * as utils from '@/utils.js'
//import _ from 'lodash';
import InfoBox from './InfoBox.vue';
import Image from './Image.vue';
//import Tooltip from "./Tooltip.vue";
import Entities from './Entities.vue';
import {storeStore} from './../stores/store.js';

const emit = defineEmits(['emit-handler', '']);
const props = defineProps(
	reactive({
		prop: {
			type: Object,
			required: false,
			default: {},
		},
		obj: {
			type: Object,
			required: false,
			default: {},
		},
		image: {
			type: String,
			required: false,
			default: '',
		},
		size: {
			type: Number,
			required: false,
			default: 0,
		},
	})
);
const propDefaults = {
	class: null,
};
const propComputed = computed(() => {
	return Object.assign(propDefaults, props.prop);
});

const state = reactive({
	tooltip: false,
	info: false,
});

onBeforeMount(async () => {});
onMounted(async () => {});

function toggle(e) {
	let bool;
	if (propComputed.value.class == 'wand' || propComputed.value.class == 'spell' || propComputed.value.class == 'spell_alwayscast') bool = storeStore().entitiesExpanded.wands[propComputed.value.n];
	if (propComputed.value.class == 'pocket') bool = storeStore().entitiesExpanded.pocket;
	if (propComputed.value.class == 'backpack') bool = storeStore().entitiesExpanded.backpack;
	if (propComputed.value.class == 'perks') bool = storeStore().entitiesExpanded.perks;
	if (bool == undefined) bool = false;
	if (e == null) return bool;
	if (propComputed.value.class == 'wand' || propComputed.value.class == 'spell' || propComputed.value.class == 'spell_alwayscast') storeStore().entitiesExpanded.wands[propComputed.value.n] = !storeStore().entitiesExpanded.wands[propComputed.value.n];
	if (propComputed.value.class == 'pocket') bool = storeStore().entitiesExpanded.pocket = !storeStore().entitiesExpanded.pocket;
	if (propComputed.value.class == 'backpack') bool = storeStore().entitiesExpanded.backpack = !storeStore().entitiesExpanded.backpack;
	if (propComputed.value.class == 'perks') bool = storeStore().entitiesExpanded.perks = !storeStore().entitiesExpanded.perks;
}
</script>

<template>
	<div :class="['entity-container-wrapper', propComputed.class + '-wrapper']">
		<span :class="propComputed.class">
			<div
				:class="['entity-container', 'prevent-select', propComputed.class, 'entity-container-default']"
				v-if="obj"
				style="border: 1px solid rgb(0, 0, 0)"
				:title="JSON.stringify(obj?.['$'], null, 2)"
				@click="
					toggle($event);
					state.info = !state.info;
				"
			>
				<span v-if="propComputed.class == 'spell_alwayscast'" style="">
					<Image
						:style="{border: '5px solid rgb(0, 0, 0)'}"
						:obj="obj"
						:image="obj['$']?.img ? obj['$'].img : ''"
						@emit-handler="
							test($event, '***********emit');
							emit('emit-handler', $event);
						"
					/>
				</span>
				<span v-else style="">
					<Image
						:obj="obj"
						:image="obj['$']?.img ? obj['$'].img : ''"
						@emit-handler="
							test($event, '***********emit');
							emit('emit-handler', $event);
						"
					/>
					<span class="badge-count prevent-select" v-if="obj?.count">{{ obj?.count }}</span>
				</span>
				<p v-if="obj['$']?.uses_remaining != '-1' && obj['$']?.uses_remaining != null" class="uses-remaining" style="">
					{{ obj['$']?.uses_remaining }}
				</p>
				<p v-if="propComputed.class == 'wand' && obj['$']?.gun_level != '-1' && obj['$']?.gun_level != null" class="gun-level" style="">
					{{ obj['$']?.gun_level }}
				</p>
			</div>
			<span v-if="toggle() && obj && Object.keys(obj).length > 0">
				<InfoBox :obj="obj" />
			</span>
		</span>
		<span v-if="propComputed.class == 'wand'" style="display: flex; flex-direction: row">
			<Entities :obj="obj['Spells_alwayscast']" :prop="{...propComputed, class: 'spells_alwayscast', size: obj['spells_alwayscast']?.length}" :size="size" />
			<Entities :obj="obj['Spells']" :prop="{...propComputed, class: 'spells', size: obj['$']?.deck_capacity}" :size="size" />
		</span>
	</div>
</template>

<style scoped>
/* <span class="badge-count prevent-select" v-if="obj?.count" style="position:absolute; font-size:1ch; padding: 0 4px 0 4px; bottom:1px; left:1px; border: 1px solid red; border-radius: 10px; background-color: white; color:black; font-weight: bold;" >{{ obj?.count }}</span> */
.badge-count {
	position: absolute;
	/* font-size:1ch;  */
	padding: 0 2px 0 2px;
	bottom: 1px;
	left: 1px;
	border: 1px solid red;
	border-radius: 10px;
	background-color: rgb(0, 0, 0);
	color: rgb(255, 255, 255);
	font-weight: 900;
	/* text-shadow: 1px 1px red; */
	/* text-shadow: 0 0 3px #3300ff; */
	/* text-shadow: 0 0 3px #ff0000, 0 0 5px #0000ff; */

	font-size: 12px;
	line-height: 1;
}

.wand,
.spell,
.spell_alwayscast,
.entity-container {
	/* background-color: rgb(62, 47, 47); */
	/* margin: 0px 0px 0px 0px;
	padding: 0px 0px 0px 0px; */
	border-radius: 5px;

	align-self: start;
	/* display:flex; */
	/* align-self: start; */
	/* align-self: stretch; */
	/* align-items: start; */
	/* align-items: start;
	background-color: bisque; */
}
.entity-container {
	/* background-color: rgb(62, 47, 47);
	border-radius: 5px; */

	/* display:inline-block;  */
	position: relative;
	width: 32px;
	height: 32px;

	/* 
	align-items: start;
	justify-items: start;
	 */
	/* align-items: start; */
}
.entity-container-wrapper {
	display: flex;
	flex-direction: row;
	/* flex-direction:column; */
	/* border: 1px solid rgb(0, 0, 0); */
	/* border-radius: 1px; */
	margin: 0px;
	padding: 0px;
	width: fit-content;
	height: fit-content;

	/* background-color: bisque; */
	align-items: start;
	align-self: start;
	/* display:flex;
	align-self: start; */
	/* align-items: start; */

	/* align-items: start;
	justify-items: start; */
	/* align-items: flex-start;
	justify-items: flex-start;
	justify-content: flex-start; */
}
.entity-container div {
	display: flex;
	align-items: center;
	/* align-content: center; */
	/* justify-items: start; */
}

.wand {
	/* display:inline-flex;	
	flex-direction: row; */
	/* background-color: rgba(19, 25, 95, 0.7); */
	/* border-radius: 1px; */

	/* background-color: bisque;
	display:flex;
	align-items: start; */
	/* justify-items: start; */
}
.spell_alwayscast {
	/* margin: 1px ; */
	/* border: 2px solid rgb(0, 0, 0); */
	background-color: rgba(0, 0, 0, 0.9);
	/* background-color: rgba(0, 0, 0, 0.5); */
}

.uses-remaining {
	position: absolute;
	bottom: 0px;
	right: 0px;
	font-size: 8px;
	height: fit-content;
	background-color: rgba(0, 0, 0, 0.5);

	border: 1px solid rgb(0, 0, 0);
	border-radius: 5px;
}
.gun-level {
	position: absolute;
	bottom: 0px;
	left: 1px;
	font-size: 8px;
	height: fit-content;
	background-color: rgba(0, 0, 0, 0.5);

	border: 1px solid rgb(0, 0, 0);
	border-radius: 5px;
}

.entity-info {
	/* font-size:8px;
	display:inline-flex;
	flex-direction: column; */
	/* display:inline-block; */
	/* position:relative; */
	/* display:inline-flex;
	flex-direction: row; */
}
.entity-info span {
	display: inline-flex;
	flex-direction: row;
}
</style>
