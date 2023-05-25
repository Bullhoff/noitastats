<script setup>
import {reactive, onMounted, computed, nextTick} from 'vue';
//import * as app from '../../../wailsjs/go/main/App'
//import * as runtime from '../../../wailsjs/runtime/runtime.js'
//import { storeToRefs } from 'pinia'
//import * as store from "./../../stores/store.js"; //imageStore()
import {refStore, dataStore} from './../../stores/store.js'; //imagesStore

onMounted(() => {});
async function emitHandler(e) {}
</script>

<template>
	<div class="container">
		<div>
			<TreeView
				:obj="{
					PlayerSummary: dataStore().PlayerSummary,
					SessionSummary: dataStore().SessionSummary(),
					WorldStateSummary: dataStore().WorldStateSummary,
				}"
				@emit-handler="emitHandler"
			/>
		</div>
	</div>
</template>

<style scoped>
.container {
	position: unset;
	width: 100%;
	display: flex;
	flex-direction: column;
}
</style>
