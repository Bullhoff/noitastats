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

<template v-if="dataStore().sorted?.WorldStateXml?.changed_materials">
	<div class="container">
		<table>
			<tr v-for="item in dataStore().sorted.WorldStateXml.changed_materials">
				<td>{{ item[0] }}</td>
				<td>&#x1F852;</td>
				<td>{{ item[1] }}</td>
			</tr>
		</table>
		<!-- 
      <div v-if="false">
        &#x1F80A;
        &#x1F852;
        &#x1F81E;
        &#x1F8A1;
        &#x1F872;
        &#x1F8A7;
        &#x1F846;
        &#x1F862;
      </div> 
    -->
	</div>
</template>

<style scoped>
.container {
	position: relative;
	width: fit-content;
	margin: 5px 5px 5px 5px;
}
table {
	border-spacing: 0;
	border-collapse: collapse;
}
table tr td {
	padding: 0 15px 0 15px;
	border: 1px solid rgb(82, 82, 82);
	width: fit-content;
	text-align: left;
	line-height: 1;
}
</style>
