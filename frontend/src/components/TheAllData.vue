<script setup>
import {reactive, onMounted, computed} from 'vue';
import {storeToRefs} from 'pinia';
import * as store from '../stores/store';
import {storeStore, dataStore} from '../stores/store';

const emit = defineEmits([]);
const props = defineProps({});
const refs = reactive({});
const state = reactive({});
function buttonClick(e) {}

onMounted(() => {});

var treeViewObj = {
	DATA: {
		raw: dataStore().$state.raw,
		sorted: dataStore().$state.sorted,
		spellList: store.dataStore().spellList,
		perkList: store.dataStore().perkList,
		sessionList: store.dataStore().sessionList,
		xmlAndLua: store.dataStore().xmlAndLua(),
		killedBy: dataStore().killedBy,
		everything: {
			config: storeStore().config,
			windows: store.windowStore().$state.windows,
			dataStore: store.dataStore(),
			windowStore: store.windowStore(),
			storeStore: store.storeStore(),
			_: {
				localStorage: dataStore().getLocalStorage,
				DataEntityList: dataStore().DataEntityList,
				DataEntityEntity: dataStore().DataEntityEntity(),
				DataEntityListNested: dataStore().DataEntityListNested(),
				getSpellsSummary: dataStore().getSpellsSummary,
				getPerksSummary: dataStore().getPerksSummary,
				SessionSummary: dataStore().SessionSummary(dataStore().raw.Gameid),
				killedBy: dataStore().killedBy,
			},
		},
	},
};
function treeViewClick({e, key, obj}) {
	//console.log('treeViewClick', e)
	openWindow({e, title: key, obj, component: 'TreeView'});
}
function openWindow({component, e, title, obj} = {}) {
	let width = component == 'TreeView' ? 'fit-content' : '90vw';
	store.windowStore().createNewWindow({
		component,
		e,
		prop: {
			head: {text: title},
			body: {
				obj,
				style: {container: {width: width}},
			},
		},
	});
}
function contextMenu(obj) {
	let {e} = obj;
	let items = {
		'Open in new window': () => openWindow(obj),
	};
	store.contextMenuStore().onContextMenu(e, items);
}
</script>

<template>
	<div class="container" :ref="(el) => (refs.allData = el)">
		<!-- <Container>
				<template v-slot:head>
					<span @contextmenu.prevent="contextMenu({component: 'Table', e:$event, title:'spellList', obj:store.dataStore().spellList})">
						TEST
					</span>
				</template>
				<Table 
					:table="{
						sortBy:{key:'id',ascending:false},
					}"
					:obj="{
						row1: {id: 'a', subobj: {moo1: 'abc1', moo2: 'bcd1'}},
						row2: {id: 'b', subobj: {moo1: 'abc2', moo2: 'bcd2'}},
						row3: {id: 'c', subobj: {moo1: 'abc3', moo2: 'bcd3'}},
						row4: {id: 'd', subobj: {moo1: 'abc4', moo2: 'bcd4'}},
					}"
				/>
			</Container> -->

		<Container>
			<template v-slot:head>
				<span @contextmenu.prevent="contextMenu({component: 'Table', e: $event, title: 'spellList', obj: store.dataStore().spellList})"> Spells </span>
			</template>
			<Table
				:table="{
					sortBy: {key: 'id', ascending: false},
					/* showColumns: ['id', 'mana', 'lifetime', 'lifetime_add'], */
				}"
				:obj="store.dataStore().spellList"
			/>
		</Container>

		<Container>
			<template v-slot:head>
				<span @contextmenu.prevent="contextMenu({component: 'Table', e: $event, title: 'perkList', obj: store.dataStore().perkList})"> Perks </span>
			</template>
			<Table
				:table="{
					sortBy: {key: 'id', ascending: false},
					/* showColumns: ['id', 'ui_icon'], */
				}"
				:obj="store.dataStore().perkList"
			/>
		</Container>

		<Container>
			<template v-slot:head>
				<span @contextmenu.prevent="contextMenu({component: 'Table', e: $event, title: 'sessionList', obj: store.dataStore().sessionList})"> Sessions </span>
			</template>
			<Table
				:table="{
					sortBy: {key: 'id', ascending: false},
					/* showColumns: ['id', 'ui_icon'], */
				}"
				:obj="store.dataStore().sessionList"
			/>
		</Container>

		<TreeView :obj="treeViewObj" :elements="[{type: 'button', text: String.fromCodePoint('0x1F5D7'), onClick: treeViewClick, placement: 'left', rows: 'objects'}]" />
	</div>
</template>

<style scoped>
.container {
	margin: 5px;
	display: flex;
	flex-direction: column;
	gap: 5px;
}
</style>
