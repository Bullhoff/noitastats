<script setup>
import {ref, reactive, onMounted, onBeforeMount, onUpdated, nextTick, computed, watch} from 'vue';
import * as store from './../stores/store.js';
import {storeToRefs} from 'pinia';
import {imageStore} from './../stores/store.js';
const {} = storeToRefs(imageStore());

import * as app from '../../wailsjs/go/main/App';

const emit = defineEmits(['emit-handler', '']);
const props = defineProps(
	reactive({
		obj: {
			type: Object,
			required: false,
			default: {},
		},
		config: {
			type: Object,
			required: false,
			default: {
				backgroundImage: 'data/ui_gfx/inventory/full_inventory_box.png',
				backgroundImageHover: 'data/ui_gfx/inventory/full_inventory_box_highlight.png',
			},
		},
		image: {
			type: String,
			required: false,
			default: '',
		},
		backgroundImage: {
			type: String,
			required: false,
			default: '',
		},
		style: {
			type: Object,
			required: false,
			default: {},
		},
	})
);
const refs = reactive({img: null});
const state = reactive({hover: false});
const img = computed(() => props.image);

onBeforeMount(async () => {});
onMounted(async () => {});
</script>

<template>
	<span></span>
	<div style="" class="image-container prevent-select" :style="{...props.style, boxSizing: 'border-box', backgroundColor: 'black'}" :title="JSON.stringify(obj)">
		<span></span>
		<svg
			width="100%"
			height="100%"
			xmlns="http://www.w3.org/2000/svg"
			@mouseenter="state.hover = true"
			@mouseover="state.hover = true"
			@mouseleave="state.hover = false"
			@mouseout="state.hover = false"
			@click="
				() => {
					if (img && !imageStore().getImage(img)) {
						store.dataStore().unpackAssets();
					}
				}
			"
		>
			<image v-if="config.backgroundImage && imageStore().getImage(config.backgroundImage) && state.hover == false" :href="imageStore().getImage(config.backgroundImage)" height="100%" width="100%" />
			<image v-else-if="config.backgroundImageHover && imageStore().getImage(config.backgroundImageHover) && state.hover == true" :href="imageStore().getImage(config.backgroundImageHover)" height="100%" width="100%" />
			<image v-if="imageStore().getImage(img)" :href="imageStore().getImage(img)" height="100%" width="100%" />
			<rect v-if="img && !imageStore().getImage(img)" x="0" width="100%" height="100%" rx="5" stroke-width="1" class="svg-rect-container" fill="green" />
			<!-- <text v-if="img && !imageStore().getImage(img) && obj?.$?.id && state.hover == true" x="50%" y="50%" font-size="50%" text-anchor="middle" dominant-baseline="central" fill="blue" stroke="blue">
					{{ obj.$.id }}
				</text> -->
			<foreignObject width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" v-if="img && !imageStore().getImage(img) && obj?.$?.id && state.hover == true">
				<div xmlns="http://www.w3.org/1999/xhtml" style="width: 100%; height: 100%; white-space: pre-wrap; line-height: 1; font-size: 8px; word-wrap: break-word; text-align: center; vertical-align: middle">{{ obj.$.id }}</div>
			</foreignObject>
			<text v-if="img && !imageStore().getImage(img) && (state.hover == false || !obj?.$?.id)" x="50%" y="50%" font-size="70%" text-anchor="middle" dominant-baseline="central" fill="blue" stroke="blue">Click</text>
		</svg>
	</div>
</template>

<style scoped>
img {
	width: 100%;
	height: 100%;
}

.image-container {
	width: 32px;
	height: 32px;
	/* width: 100%; height: 100%; */
	align-self: center;
}
svg {
	cursor: pointer;
}
</style>
