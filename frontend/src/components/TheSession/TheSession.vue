<script setup>
import {reactive, onMounted, computed, nextTick} from 'vue';

//import * as app from '../../../wailsjs/go/main/App'
//import * as runtime from '../../../wailsjs/runtime/runtime.js'
//import { storeToRefs } from 'pinia'
//import * as store from "./../../stores/store.js"; //imageStore()
import {storeStore, dataStore} from './../../stores/store.js'; //imagesStore

import TheGameInfo from './TheGameInfo.vue';
import Entities from './../Entities.vue';
import TheFungalShifts from './TheFungalShifts.vue';

onMounted(() => {});
</script>

<template>
	<div class="container" style="">
		<span>
			<table>
				<tr v-if="storeStore()?.config?.state?.session?.subviews?.wands?.open && dataStore().sorted?.PlayerXml?.Wands">
					<td>Wands</td>
					<td><Entities :obj="dataStore().sorted.PlayerXml.Wands" :prop="{class: 'wands', size: dataStore().sorted?.PlayerXml?.Wands?.length}" /></td>
				</tr>
				<tr v-if="storeStore()?.config?.state?.session?.subviews?.items?.open && dataStore().sorted?.PlayerXml?.Items">
					<td>Items</td>
					<td><Entities :obj="dataStore().sorted.PlayerXml.Items" :size="4" :prop="{class: 'pocket', size: 4}" /></td>
				</tr>

				<tr v-if="storeStore()?.config?.state?.session?.subviews?.spells?.open && dataStore().sorted?.PlayerXml?.Spells && dataStore().sorted?.PlayerXml?.Spells">
					<td>Spells</td>
					<td>
						<Entities v-if="dataStore().sorted.PlayerXml.Spells.length <= 16" :obj="dataStore().sorted.PlayerXml.Spells" :prop="{class: 'spells'}" />
						<Entities v-else :obj="dataStore().getSpellsSummary" :prop="{class: 'spells'}" />
					</td>
				</tr>
				<tr v-if="storeStore()?.config?.state?.session?.subviews?.perks?.open && dataStore().sorted?.PlayerXml?.Perks && dataStore().sorted?.PlayerXml?.Perks">
					<td>Perks</td>
					<td>
						<Entities v-if="dataStore().sorted.PlayerXml.Perks.length < 20" :obj="dataStore().sorted.PlayerXml.Perks" :prop="{class: 'perks'}" />
						<Entities v-else :obj="dataStore().getPerksSummary" :prop="{class: 'perks'}" />
					</td>
				</tr>
			</table>
			<TheGameInfo v-if="storeStore()?.config?.state?.session?.subviews?.game_info?.open" />
			<TheFungalShifts v-if="storeStore()?.config?.state?.session?.subviews?.fungal_shifts?.open" />
		</span>
	</div>
</template>

<style scoped>
.container {
	overflow-x: auto;
	overflow-y: hidden;
	width: 100%;
	height: 100%;
}
</style>
