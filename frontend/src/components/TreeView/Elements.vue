<template>
	<span v-for="(item, index) in elements" style="display: inline-block; padding: 0 1px 0 1px">
		<template v-if="item.placement == placement && item.rows == rows" @click="emit('emit-handler', item)" @change="emit('emit-handler', item)">
			<span v-if="item.type == 'button'" class="prevent-select" style="cursor: pointer">
				<Button :text="item.text" @click="item.onClick({e: $event, obj: currentValue, key: currentKey})" />
			</span>
			<span v-else-if="item.type == 'unicodeKey'" :title="currentKey + '\n' + '0x' + currentKey + '\n' + '0x1F800' + '\n' + String.fromCharCode('0x' + currentKey)">
				{{ item.text != undefined ? item.text : String.fromCharCode('0x' + currentKey) }}
			</span>
			<div v-else-if="item.type == 'checkbox'">
				<Checkbox />
			</div>
		</template>
	</span>
</template>

<script setup>
import {h, onMounted, onBeforeMount, watch, onUpdated, nextTick, ref, reactive, unref, toRaw, isRef, onRenderTriggered, computed} from 'vue';

const emit = defineEmits(['emit-handler']);
const props = defineProps(
	reactive({
		placement: {default: ''},
		rows: {default: ''},
		currentKey: {default: ''},
		currentValue: {default: ''},
		elements: {default: []},
		obj: {
			type: Object,
			required: false,
		},
		prop: {
			type: Object,
			required: false,
			default: {
				state: {},
			},
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
const refs = reactive({});
const styleDefault = {};
const style = computed(() => {
	return {...styleDefault, ...props.style};
});

onBeforeMount(() => {});
onMounted(() => {});
onUpdated(() => {});
</script>

<style scoped></style>
