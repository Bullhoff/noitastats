<template >
	<span  
			style=""
			:style="{textAlign: 'left', ...prop.value.style}" 
			:title="(prop.value.tooltip)?prop.value.tooltip:null"
			@click.prevent="{(prop.value.onClick)?prop.value.onClick():''; }"
			@click.left.prevent="{(prop.value.onClickLeft)?prop.value.onClickLeft():''; }"
			@click.middle.prevent="{(prop.value.onClickMiddle)?prop.value.onClickMiddle():''; }"
			@click.right.prevent="{(prop.value.onClickRight)?prop.value.onClickRight():''; }"
		>
			{{ (prop.value.value)? prop.value.value : prop.key }}
	</span>
</template>


<script setup>
import _ from 'lodash';
import Clickable from './Clickable.vue'
//import * as utils from './../utils.js'

import {
	ref, reactive, toRefs, watch, watchEffect, onMounted, onBeforeMount, onUpdated, nextTick,
	onDeactivated, onRenderTriggered, computed
} from 'vue'

const emit = defineEmits(['opening-sub'])
const props = defineProps(reactive({
	prop: {
		type: Object,
		default: {key:null, value:null, style:{}, type: null},
	},
}))

const refs = reactive({})
const containerRef = ref(null)
const treeObj = reactive({})
const p = reactive(Object.assign({}, props))


onBeforeMount(async() => {})
onMounted(async () => {})
onUpdated(async ()=>{})



</script>



<style scoped>


</style>