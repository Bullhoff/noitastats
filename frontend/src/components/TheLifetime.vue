<script setup>
import {ref, reactive, onMounted, onBeforeMount, onUpdated, nextTick, computed} from 'vue';
//import _ from 'lodash';
import Image from './Image.vue';
import {dataStore} from './../stores/store.js';
import * as store from './../stores/store.js';

const emit = defineEmits([]);
const props = defineProps({});

const state = reactive({
	//radioSelectSource: 'fromInput',	// 'fromScan'
	filterInput: {},
	sortBy: {key: 'lifetime', ascending: false},
});
var tableSpellDropdown = reactive({
	//showFilterButton: false,
	sortBy: {key: 'lifetime', ascending: false},
	showFilter: true,
	showColumns: ['id', 'lifetime'],
	excludeColumnValues: {lifetime: ''},
	filterInput: {},
});

function reset(to = 0) {
	for (const [key, value] of Object.entries(store.lifetimeStore().allItems)) {
		value.amount = 0;
	}
	if (to == 'inventory') {
		for (let i = 0; i < store.lifetimeStore().InventoryLifeTimeAffecting.length; i++) {
			store.lifetimeStore().allItems[store.lifetimeStore().InventoryLifeTimeAffecting[i].$.id].amount += 1;
		}
	}
}

var resultSet = new Set();

function recursion(obj) {
	let defaultArgs = {path: [], start: 0, sum: 0, items: [], r: 10, index: 0};
	let obj2 = Object.assign(defaultArgs, JSON.parse(JSON.stringify(obj)));
	let {path, items, r, start, index} = JSON.parse(JSON.stringify(obj2));
	if (index == r) {
		return;
	}
	for (let i = start; i < items.length; i++) {
		let amount = path.filter((x) => x == items[i]).length;
		if (amount < store.lifetimeStore().allItems[items[i]].amount) {
			path[index] = items[i];
			resultSet.add(path.join('__'));
			recursion({path, items, r, start: i, index: index + 1});
		}
	}
}

function calculateOptions() {
	resultSet = new Set();

	var perks = {};
	var spells = {};
	for (const [key, value] of Object.entries(store.lifetimeStore().allItems)) {
		if (value.type == 'perk') perks[key] = value;
		else spells[key] = value;
	}
	recursion({items: Object.keys(spells), r: store.lifetimeStore().result.count_spells});
	let arr = Array.from(resultSet);

	for (let i = 0; i < arr.length; i++) {
		arr[i] = {path: arr[i].split('__')};
		// ADD PERKS
		for (const [key, value] of Object.entries(perks)) {
			for (let y = 0; y < value.amount; y++) {
				arr[i].path = [key, ...arr[i].path];
			}
		}
		let sum = 0;
		for (let y = 0; y < arr[i].path.length; y++) {
			if (!arr[i].path[y]) continue;
			sum += store.lifetimeStore().allItems[arr[i].path[y]].lifetime;
		}

		let aim = 1 + store.lifetimeStore().spellSelected.lifetime; //+ perkBonus	//1+target + perkBonus

		arr[i] = {
			[sum + '+1+' + store.lifetimeStore().spellSelected.lifetime]: sum + 1 + store.lifetimeStore().spellSelected.lifetime,
			sum_raw: sum,
			sum: sum + aim,
			...arr[i],
		};
	}
	let sorted = arr.sort((a, b) => Math.abs(a['sum']) - Math.abs(b['sum']));
	for (let i = 0; i < 20; i++) {
		store.lifetimeStore().resultArray[i] = sorted[i] ? sorted[i] : undefined;
	}
}

onBeforeMount(async () => {});
onMounted(() => {
	reset('inventory');
});

var style = {width: '32px', height: '32px'};
</script>

<template>
	<!-- 
Divide By 2		DIVIDE_2
Divide By 3		DIVIDE_3
Divide By 4		DIVIDE_4
Divide By 10	DIVIDE_10
	 -->
	<div class="container">
		<fieldset style="width: fit-content; padding-top: 1ch">
			<legend>Spells in inventory</legend>
			<div style="display: flex; flex-direction: row">
				<span v-for="item in store.lifetimeStore().InventoryLifeTimeAffecting">
					<Image :style="style" :image="item.$.img" />
				</span>
			</div>
		</fieldset>

		<fieldset style="width: fit-content; padding-top: 1ch">
			<legend>Target spell</legend>
			<table>
				<tr>
					<td>
						<Image :style="style" :image="dataStore().LuaData.gun_actions?.[store.lifetimeStore().spellSelected.id]?.sprite" />
					</td>
					<td>Lifetime: <input type="number" v-model="store.lifetimeStore().spellSelected.lifetime" style="width: 7ch" @change="" /></td>
					<td>
						<Dropdown
							:config="{arrowLeft: true}"
							:style="{
								content: {},
							}"
						>
							<template v-slot:head>
								{{ store.lifetimeStore().spellSelected.id }}
							</template>

							<Table
								@click="
									(obj) => {
										(store.lifetimeStore().spellSelected.id = obj.id), (store.lifetimeStore().spellSelected.lifetime = Number(obj.lifetime));
									}
								"
								:table="tableSpellDropdown"
								@table-state="Object.assign(tableSpellDropdown, $event)"
								:style="{
									item: {
										cursor: 'pointer',
									},
								}"
								:obj="store.dataStore().spellList"
							/>
						</Dropdown>
					</td>
				</tr>
			</table>

			<!-- 
				:table="{
										showFilterButton: false,
										sortBy: state.sortBy, 
										showFilter: true, 
										showColumns: ['id', 'lifetime', ],
										excludeColumnValues: {lifetime:''},
										filterInput: state.filterInput, 
									}"
			 --></fieldset>

		<div>
			<fieldset style="width: fit-content; padding-top: 1ch">
				<legend>Spells</legend>
				<table class="table-amount">
					<tr>
						<th></th>
						<th></th>
						<th></th>
						<th></th>
						<th>
							<span class="prevent-select" style="display: inline-grid; grid-template-columns: 50% 50%; width: 100%; margin: 1px 1px 1px 1px; line-height: 1">
								<span style="border: 1px solid red; border-radius: 5px; text-align: center; cursor: pointer" @click="reset('inventory')" title="Reset to inventory">I</span>
								<span style="border: 1px solid red; border-radius: 5px; text-align: center; cursor: pointer" @click="reset()" title="Reset to 0">0</span>
							</span>
						</th>
					</tr>

					<tr v-for="(value, key) in store.lifetimeStore().allItems" :style="[value.type == 'perk' ? {backgroundColor: 'rgba(10,10,10,0.5)'} : {}]">
						<td><Image :style="style" :image="value.image" style="" /></td>
						<td>
							<p style="">{{ value.name }}</p>
						</td>
						<td>
							<p>{{ value.lifetime > 0 ? '+' + value.lifetime : value.lifetime }}</p>
						</td>
						<td><Image :style="style" :image="value.image" /></td>
						<td><input type="number" min="0" v-model="value.amount" style="width: 5ch" /></td>
						<td>
							<p>{{ value.amount * value.lifetime }}</p>
						</td>
						<td>
							<p>{{ value?.type }}</p>
						</td>
					</tr>
					<tr>
						<td></td>
						<td></td>
						<td></td>
						<td></td>
						<td>{{ store.lifetimeStore().result.count }}</td>
					</tr>
				</table>
			</fieldset>
		</div>

		<fieldset style="width: fit-content; padding-top: 1ch">
			<legend>Find combinations</legend>
			<div>
				<div style="display: flex; flex-direction: row; width: fit-content">
					<button style="grid-column: 3 / span 3" @click="calculateOptions()">Calculate</button>
					<p v-if="store.lifetimeStore().result.count > 55" style="color: red">This will take a while</p>
				</div>
				<table>
					<tr>
						<th>Sum</th>
						<th>Result</th>
						<th colspan="100%">Combo</th>
					</tr>
					<tr v-for="row in store.lifetimeStore().resultArray">
						<template v-if="row?.sum != undefined">
							<td>{{ row?.sum_raw }}</td>
							<td :style="[row?.sum == 0 ? {backgroundColor: 'green'} : {}]">{{ row?.sum - 1 }}</td>
							<td v-if="row?.path" v-for="item in row.path">
								<Image :style="style" :image="store.lifetimeStore().allItems[item].image" />
							</td>
						</template>
					</tr>
				</table>
			</div>
		</fieldset>
	</div>
</template>

<style scoped>
table {
	border-spacing: 0;
	border-collapse: collapse;
}
.table-amount th {
	border: 0px solid rgb(0, 0, 0);
}

table th,
td {
	border: 2px solid rgb(0, 0, 0);
	padding: 0 5px 0 5px;
	text-align: left;
	/* width:fit-content; */
}

.container {
	display: flex;
	flex-direction: column;
}

.items {
	display: inline-grid;
	grid-template-columns: fit-content(5ch) fit-content(15ch) fit-content(15ch) fit-content(5ch) fit-content(5ch) fit-content(5ch) fit-content(20ch);
	/* grid-template-columns: fit-content(5ch) repeat( 3, minmax(0px, 1fr) ); */
	/* grid-template-columns: auto-fill auto-fill auto-fill auto; */
	text-align: left;
	gap: 0 10px;
	align-items: center;
}
.item p {
	border: 1px solid red;
	text-align: left;
	/* padding: 0 12px 0 2px; */
}
input {
	background-color: black;
	color: white;
}
</style>
