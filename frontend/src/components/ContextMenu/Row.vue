<template >
	<template v-if="true || prop.key == 'BLACK_HOLE'">
		<Clickable>
			<div v-if="rowComputed" style="position:relative; display:flex; flex-direction: row; width:fit-content;">
				<div v-for="(value, key, index) in rowComputed" style="position:relative; width:fit-content; border: 1px solid red; padding: 0 2px 0 2px;">
					{{ value }}
				</div>
			</div>
		</Clickable>
	</template>
</template>


<script setup>
import _ from 'lodash';
import Clickable from './Clickable.vue'

import {
	ref, reactive, toRefs, watch, watchEffect, onMounted, onBeforeMount, onUpdated, nextTick,
	onDeactivated, onRenderTriggered, computed
} from 'vue'

const emit = defineEmits([])
const props = defineProps(reactive({
	prop: {
		type: Object,
		default: {key:null, value:null, style:{}, type: null},
	},
	key: {	default: 0},
	value: {	default: 'V'},
	config: {	default: {}},
}))
const refs = reactive({})
const containerRef = ref(null)
const treeObj = reactive({})
const p = reactive(Object.assign({}, props))

const rowComputed = computed(()=>{
	let resObj = {}
	let resArr = []
	if(!props.prop.value) return
	for (let i = 0; i < props.prop.config.row.length; i++) {
		let col = [...props.prop.config.row[i]]
		let res = getObjectValue(props.prop.value, col)
		if(!res) return null
		for (let y = 0; y < res.length; y++) {
			if(res[y] == undefined) return
		}
		resArr.push(res)
	}
	return resArr
})
function getObjectValue(obj = {}, path = []){
	if(!obj) return
	if(!obj[path[0]] && Array.isArray(obj) && obj[0]) 
		return getObjectValue(obj[0], path)
	
	let k = path.shift()
	if(!path || path.length > 0) {
		if(!obj[k] && obj[0]) return getObjectValue(obj[0], path)
		else return getObjectValue(obj[k], path)
	}
	return obj[k]
}


onBeforeMount(async() => {})
onMounted(async () => {})
onUpdated(async ()=>{})



</script>



<style scoped>

</style>