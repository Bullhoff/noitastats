<script setup>
import {ref, reactive, onMounted, onBeforeMount, onUpdated, nextTick, computed, watch} from 'vue';
import Entity from './Entity.vue';

const emit = defineEmits(['emit-handler', '']);
const props = defineProps(
	reactive({
		prop: {
			type: Object,
			required: false,
			default: {},
		},
		parent: {
			type: Object,
			required: false,
			default: {},
		},
		obj: {
			type: Object,
			required: false,
			default: {},
		},
		images: {
			type: Object,
			required: false,
			default: {},
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
	nextclass: null,
	size: 0,
};

const propComputed = computed(() => {
	let res = JSON.parse(JSON.stringify(Object.assign(propDefaults, props.prop)));
	res.size = parseInt(res.size) ? parseInt(res.size) : 0;
	if (res.class == 'wands') {
		res.nextclass = 'wand';
	}
	if (res.class == 'wand') {
	}
	if (res.class == 'spells') {
		res.nextclass = 'spell';
	}
	if (res.class == 'spells_alwayscast') {
		res.nextclass = 'spell_alwayscast';
	}

	for (let i = 0; i < props.obj?.length; i++) {
		let x = parseInt(props?.obj[i]?.['$']?.x);
		if (x >= props.prop?.size) res.size = x + 1;
	}
	return res;
});

const arrComputed = computed(() => {
	if (Object.keys(props.obj).length == 0) return [];
	let res = props.obj;
	if (propComputed.value.size)
		res = props.obj.reduce((a, c) => {
			let x = parseInt(c['$'].x); //a[c['$'].x]
			if (a[x]) {
				for (let i = x; i <= propComputed.value.size; i++) {
					if (!a[i]) {
						a[i] = c;
						break;
					}
				}
			} else {
				a[x] = c;
			}
			return a;
		}, Array(propComputed.value.size));
	return res;
});
onBeforeMount(async () => {});
onMounted(async () => {});
</script>

<template>
	<template :class="['entities-container', propComputed.class]" :id="propComputed.class">
		<span v-for="n in propComputed.size" :key="JSON.stringify(obj) + JSON.stringify(arrComputed) + n" v-if="propComputed.size != 0" :title="JSON.stringify(obj, null, 2)">
			<Entity
				:obj="arrComputed[n - 1]"
				:prop="{
					...propComputed,
					n: propComputed.class == 'wands' ? n : propComputed.n,
					class: propComputed.nextclass ? propComputed.nextclass : propComputed.class,
				}"
				:size="size"
			/>
		</span>

		<span v-for="item in arrComputed" :key="JSON.stringify(item)" v-else>
			<Entity :obj="item" :prop="{...propComputed, class: propComputed.nextclass ? propComputed.nextclass : propComputed.class}" :size="size" />
		</span>
	</template>
</template>

<style scoped>
.wands span,
.spells,
.spells_alwayscast,
#backpack,
#pocket,
#perks,
.entities-container {
	display: flex;
	flex-direction: row;
}
.wands {
	display: flex;
	flex-direction: column;
}
.info {
	display: flex;
	flex-direction: row;
	background-color: aqua;
}
</style>
