




var counters = {
	parseLayout: 0, 
	parseCategory: 0, 
	getToKey: 0, 
}

function iterateLayouts(){
	for (const [listKey, listValue] of Object.entries(summaryLayouts)) {
	  
	}
}

var layoutFresh
var gameid

export function init(templatename, template, allData){
	gameid = allData.noitaDataSorted?.GameData?.Gameid
	layoutFresh = template
	//summary3 = parseLayout(layout, allData)
	//allData.templateData[templatename] = 
	return parseLayout(template, allData)
}
export function blaa(allData, summary3, templates){
	//return {}
	gameid = allData.noitaDataSorted?.GameData?.Gameid
	//return blaa2(allData, summary3)
	//allData.noitaDataSorted
	console.log('blaa')
	let layout = {
		GameData: {
			"GameId": ["noitaDataSorted", "GameData", "Gameid"], 
			"world_seed": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-world_seed"], 
			//"death_pos.x": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-death_pos.x"], 
			//"death_pos.y": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-death_pos.y"], 
			"Death pos": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", [{"key":"-death_pos.x"}, {"text":", "}, {"key":"-death_pos.y"}]], 
			"playtime_str": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-playtime_str"], 
			"projectiles_shot": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-projectiles_shot"], 

			"enemies_killed": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-enemies_killed"], 
			"dead": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-dead"], 
			"death_count": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-death_count"], 
			"killed_by": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-killed_by"], 
			"killed_by_extra": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-killed_by_extra"], 

			"gold": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-gold"], 
			"gold_all?": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-gold_all"], 
			"healed": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-healed"], 
			"places_visited": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-places_visited"], 

			"items": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-items"], 
			"streaks": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-streaks"], 
			"hp": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-hp"], 
			
			"kills": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"kills", "Stats", "-kills"], 
			
		},

		Stats: {
			"*": ["noitaDataRaw", "Stats", {"key": "$KEY", "type":"object"}, {"template": "$Stats"}],
		},
		"$Stats":{ 
			"*": ["Stats", [
				{"text":"Deaths: "}, {"key":"-deaths"}, {"text":", "}, 
				{"text":"kills: "}, {"key":"-kills"}, {"text":", "}, 
				{"text":"player_kills: "}, {"key":"-player_kills"}, {"text":", "}, 
				{"text":"player_projectile_count: "}, {"key":"-player_projectile_count"}
			]]
		},
		Stats_: {
			"*": ["noitaDataRaw", "Stats", {"key": "$KEY", "type":"object"}, {"template": "$Stats_"}],
		},
		"$Stats_": {
			"Deaths": ["Stats", "-deaths"],
			"kills": ["Stats", "-kills"],
			"player_kills": ["Stats", "-player_kills"],
			"player_projectile_count": ["Stats", "-player_projectile_count"],
		},

		Perks: {
			"*": ["noitaDataSorted", {"key": "Perks", "type":"array"}, {"template": "$Perks"}],
		},
		"$Perks": {
			"name": ["-name"],
			"tags": ["-tags"],
			"effect": ["GameEffectComponent","-effect"],
			"_tags": ["GameEffectComponent","-_tags"],
			"teleportation_probability": ["GameEffectComponent","-teleportation_probability"],

			"description": ["UIIconComponent","-description"],
			"icon_sprite_file": ["UIIconComponent","-icon_sprite_file"],
			"name": ["UIIconComponent","-name"],
			//"effect": ["UIIconComponent","-effect"],
		},

		
		Wands: {
			"*": ["noitaDataSorted", {"key": "Wands", "type":"array"}, {"template": "$Wands"}]
		}, 
		"$Wands": {
			"Tags": [ "-tags"	],
			"Wand level": [	"AbilityComponent", "-gun_level"	],
			"image file": [	"SpriteComponent", "-image_file"	],
			"damage_multiplier ": [ "HitboxComponent", "-damage_multiplier"],
			"mana": [ "AbilityComponent", "-mana"],
			"mana_max": [ "AbilityComponent", "-mana_max"],
			"mana_charge_speed": [ "AbilityComponent", "-mana_charge_speed"],
			"reload_time_frames": ["AbilityComponent", "-reload_time_frames"],
			"stat_times_player_has_shot": [ "AbilityComponent", "-stat_times_player_has_shot"],
			"ui_name": [ "AbilityComponent", "-ui_name"],
			"stat_times_player_has_edited": ["AbilityComponent", "-stat_times_player_has_edited"],
			"sprite_file": [ "AbilityComponent", "-sprite_file"],
			
			"speed_multiplier": [ "AbilityComponent", "gunaction_config", "-speed_multiplier"],
			"bounces": [ "AbilityComponent", "gunaction_config", "-bounces"],
			"blood_count_multiplier": [ "AbilityComponent", "gunaction_config", "-blood_count_multiplier"],
			"fire_rate_wait": [ "AbilityComponent", "gunaction_config", "-fire_rate_wait"],
			"action_mana_drain": [ "AbilityComponent", "gunaction_config", "-action_mana_drain"],
			"spread_degrees": [ "AbilityComponent", "gunaction_config", "-spread_degrees"],
			"action_max_uses": [ "AbilityComponent", "gunaction_config", "-action_max_uses"],
			
			"inventory_slot.x": [ "ItemComponent", "-inventory_slot.x"],
			"inventory_slot.y": [ "ItemComponent", "-inventory_slot.y"],
			"spawn_pos.x": [ "ItemComponent", "-spawn_pos.x"],
			"spawn_pos.y": [ "ItemComponent", "-spawn_pos.y"],
			"spawn_pos": [ "ItemComponent", [{"key":"-spawn_pos.x"}, {"text":", "}, {"key":"-spawn_pos.y"}]],
			"Spells": {
				"*": [ {"key": "Entity", "type":"array"}, {"template": "$Spells"}]
			}
		},
		Spells: {
			"*": ["noitaDataSorted", {"key": "Spells", "type":"array"}, {"template": "$Spells"}], 
		},
		"$Spells": {
			"ID": ["ItemActionComponent", "-action_id"],
			"Tags": ["-tags"], 
			"Image file": ["SpriteComponent", "0", "-image_file"],
			"Air friction": ["VelocityComponent", "-air_friction"],
			"gravity_x": ["VelocityComponent", "-gravity_x"],
			"gravity_y": ["VelocityComponent", "-gravity_y"],
			"liquid_drag": ["VelocityComponent", "-liquid_drag"],
			"affect_physics_bodies": ["VelocityComponent", "-affect_physics_bodies"],
			"displace_liquid": ["VelocityComponent", "-displace_liquid"],
			"mass": ["VelocityComponent", "-mass"],
			"terminal_velocity": ["VelocityComponent", "-terminal_velocity"],
			"damage_multiplier": ["HitboxComponent", "-damage_multiplier"],
			"spawn_pos.x": ["ItemComponent", "-spawn_pos.x"],
			"spawn_pos.y": ["ItemComponent", "-spawn_pos.y"],
			"uses_remaining": ["ItemComponent", "-uses_remaining"],
			"item_pickup_radius": ["ItemComponent", "-item_pickup_radius"],
			"can_go_up": ["SimplePhysicsComponent", "-can_go_up"],
			"inventory_slot.x": ["ItemComponent", "-inventory_slot.x"],
			"inventory_slot.y": ["ItemComponent", "-inventory_slot.y"],
		},

		Test: {
			//"*": ["noitaDataSorted", {"key": "Spells", "type":"array"}, {"template": "$Spells"}], 
			//"*": ["noitaDataSorted", "Spells", {"template": "$Spells"}], 
			"*": ["noitaDataSorted", "Wands", {"template": "$Wands"}]
		},
		Test2: {
			"ID": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "ItemActionComponent", "-action_id"],
			"Tags": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "-tags"], 
			"Image file": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "SpriteComponent", "0", "-image_file"],
			"Air friction": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-air_friction"],
			"gravity_x": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_x"],
			"gravity_y": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_y"],
		},
		Test3: {
			"bla*":{
				"ID": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "ItemActionComponent", "-action_id"],
				"Tags": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "-tags"], 
				"Image file": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "SpriteComponent", "0", "-image_file"],
				"Air friction": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-air_friction"],
				"gravity_x": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_x"],
				"gravity_y": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_y"],
			}
		},
		
	}
	templates['test'] = layout
	console.log('**************summary*************', allData.noitaDataSorted)
	layoutFresh = layout
	//summary3 = parseLayout(layout, allData)
	allData.templateData['test'] = parseLayout(layout, allData)
	console.log('summary3', summary3)
	//getToKey2(layout, summary)
	return allData.templateData['test']
	//return summary3
}


export function parseLayout(layout = {}, summary){
	console.log('parseLayout 1', {layout, summary})
	let res = {}
	if(Array.isArray(layout)) res = getToKey("keyname", layout, summary)
	else 
	for (const [key, value] of Object.entries(layout)) {
		counters.parseLayout += 1
		//console.log('parseLayout...', key, 'VALUE:',JSON.stringify(value), Array.isArray(value), typeof value)
		
		let opt = value['$']
		if(key == '$') {}
		else if(key[0] == '$'){}
		//else if(Array.isArray(layout)) res[key] = getToKey("keyname", layout, summary)
		else if(Array.isArray(value)) res[key] = getToKey("keyname", value, summary)
		else if(typeof value == 'object') res[key] = parseCategory(value, summary)
		if(value['$'] && value['$'].array) res[key] = Object.values(res)
	}
	console.log('parseLayout DONE', res, counters)
	return res
}
function parseCategory(category = {}, summary){
	//console.log('-----------parseCategory 1', counters.parseCategory, category, summary)
	let res = {}
	for (const [key, value] of Object.entries(category)) {
		counters.parseCategory += 1
		//console.log('**********parseCategory1', counters.parseCategory, 'KEY:',  key, 'VALUE:',value, Array.isArray(value), typeof value)
		if(key == '$') {}
		else if(key == '*') res = getToKey(key, value, summary)
		//else if(Array.isArray(value)) res = getToKey(key, value, summary)
		else if(Array.isArray(value)) res[key] = getToKey(key, value, summary)
		else if(typeof value == 'object') res[key] = parseCategory(value, summary)
	}
	//console.log('parseCategory DONE', typeof res, res, 'counters:', counters)
	return res
}
function handleTemplate(keyName, row2, summary, template){
	let templ = layoutFresh[template]
	return parseCategory(templ, summary)
}
function getToKey(keyName, row2, summary){
	let row = JSON.parse(JSON.stringify(row2))
	counters.getToKey += 1

	
	//if(row2?.[0]?.[0]?.key)
	//	console.log('************row2[0]**************', row2[0])
	//console.log('*getToKey 1.ROW:', row, 'SUMMARY:',summary, 'keyName:', keyName, 'counters:',counters)
	let resObj = {}
	row = (typeof row == 'string')? [row] : row
	let k = row[0]
	row.shift()

	if(k == 'Wands')
		console.log('--------', k, typeof summary[k])
//	console.log('getToKey 2.ROW:', row, 'SUMMARY:',summary, "k:",k)
	//if(k[0] == '$'){		if(k == '$i'){}	}
	//console.log('*', k, typeof summary[k])
	if(summary && typeof summary[k] == 'object' && k == 'Wands'){	// && k == 'Wands'
		// {"template": "$Stats"}
		/* let arr = []
		for (let i = 0; i < summary[k].length; i++) {
			let res = getToKey(keyName, row, summary[k][i])
			if(res) arr.push(res)
		}
		return getToKey(keyName, [], arr) */

		let obj = []
		for (const [key, value] of Object.entries(summary[k])) {
			let res = getToKey(keyName, row, value)
			if(res) obj[key] = res
		}
		return getToKey(keyName, [], obj)
	}
	if(typeof k == 'object' && summary){

		
		//console.log('*object*', k)
		//let [k, v] = Object.entries(key)[0]
		//if(v == 'array'){}
		let {obj, type, key, template, special_key, div_key} = k
		if(false){}

		if(type == 'OP'){
			let val = getToKey(keyName, [key], summary[key])
			//let divval = getToKey(keyName, [div_key], summary[div_key])
			console.log('***********OP RES1', val, 'div_key')
			let res = val/100/1
			console.log('***********OP RES2', res)
			return getToKey(keyName, [], res)
		}
		//else if(){}
		//if(!obj && !type && !template)
		//	console.log('************k**************', k, row, summary),
		else if(Array.isArray(k) && row.length < 1){
			let str = ''
			for (let i = 0; i < k.length; i++) {
				if(k[i].text){str+=k[i].text}
				if(k[i].key){str+=summary[k[i].key]}
				//if(k[i].key){str+=from.obj[k[i].key]}
			}
			//console.log('************Array.isArray**************', str)
			return getToKey(keyName, [], str)
			//return addEntry({from: {obj: str, path: from.path}, to})
		}
		else if(template){
			let templ = layoutFresh[template]
			return parseCategory(JSON.parse(JSON.stringify(layoutFresh[template])), summary)
			let res = handleTemplate(keyName, row, summary, template)
			//console.log('************res**************', res)
			return res
			//return getToKey(keyName, [], res)
		}
		else if(type == 'array'){
			resObj[key] = []
			let arr = []
			let res
			//console.log('getToKey.0..', "summary[key]", key, summary, )
			if(!summary[key]) return
			for (let i = 0; i < summary[key].length; i++) {
				//console.log('*for*', i, special_key, row, summary[key][i])
			//	res = getToKey(keyName, row, summary[key][i])
			//	if(res) arr.push(res)
				//console.log('getToKey.1..', i, "summary[key]", summary[key], summary[key][i])
				if(special_key){
					//console.log('******getToKey.1..', i, special_key, layoutFresh[special_key], key, summary[key][i])
					
					res = getToKey(keyName, row, summary[key][i])
					let ii = getValue(JSON.parse(JSON.stringify(layoutFresh[special_key])), summary[key][i], i)
					//if(res) resObj[key].push(res)
					//console.log('******getToKey.2..', i, ii)
					if(res) arr[ii] = res
				}else{
					res = getToKey(keyName, row, summary[key][i])
					if(res) resObj[key].push(res)
					if(res) arr.push(res)
				}
				
//				console.log('getToKey.1..', i, {"summary[key][i]":summary[key][i]})
			//	let res = getToKey(keyName, row, summary[key][i])
				//console.log('getToKey.2..', i, keyName, row, {res, "summary[key][i]":summary[key][i]})
				
			}
			//console.log('getToKey..',counters, arr)
			return getToKey(keyName, [], arr)
			//return arr
			//return getToKey(keyName, [], arr)
			//return {[keyName]: resObj[key]}
			//return getToKey(row, resObj)
		}
		else if(type == 'object'){
			//console.log('*******object*******', k, gameid)
			if(key == '$Gameid'){return getToKey(keyName, row, summary[gameid])}
			if(key == '$KEY'){
				//console.log('************$KEY**************', k, gameid, row)
				let obj = {}
				for (const [key, value] of Object.entries(summary)) {
					//console.log('**for**', key, value)
					let res = getToKey(keyName, row, summary[key])
					if(res) obj[key] = res
				}
				return getToKey(keyName, [], obj)
			}
			//obj[key]
		}
	} 
	/* else if(summary && typeof summary[k] == 'object' && ){
		console.log('ööööö', k, Array.isArray(summary[k]), typeof summary[k], summary[k],  )
		if(Array.isArray(summary[k])){
			let arr = []
			if(!summary[k]) return
			for (let i = 0; i < summary[k].length; i++) {
				console.log('åå', k, i, summary[k][i]  )
				let res = getToKey(keyName, row, summary[k][i])
				if(res) arr.push(res)
			}
			return getToKey(keyName, [], arr)
		} 
		else {
			let obj = []
			for (const [key, value] of Object.entries(summary[k])) {
				let res = getToKey(keyName, row, value)
				if(res) obj[key] = res
			}
			return getToKey(keyName, [], obj)
		}
		
	} */
	else if(k &&  summary){
		//console.log('ööööö', k, summary, )
		return getToKey(keyName, row, summary[k])
	} 
	else {
		//console.log('getToKey DONE1', counters, summary)
		//console.log('getToKey DONE2', k, summary)
		if(k == '*') summary = Object.values(summary)
		return summary
		//return {[keyName]: summary}
		//return summary[k]
	}

	//row = (typeof row == 'string')? [row] : row
	/* if(!row || row.length == 0){
		return summary[key]
	} else {
		return getToKey(row, summary[key])
	} */
}
function getValue(path, obj, i = 0){
	//console.log('**getValue 1**', i, path, obj)
	if(!obj) return i
	//console.log('**getValue 2**', i, path, obj)

	let k = path[0]
	path.shift()

	//if(path.length>0) return getValue(path, obj[k])
	if(k) return getValue(path, obj[k], i)
	//else {
		//console.log('**getValue 3**', i, k, path, obj)
		//return obj[k]
	return obj
	//}
}






function blaa2(allData, summary3){
	let layout = {
		GameData: {
			"GameId": ["noitaDataSorted", "GameData", "Gameid"], 
			"world_seed": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-world_seed"], 
			//"death_pos.x": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-death_pos.x"], 
			//"death_pos.y": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-death_pos.y"], 
			"Death pos": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", [{"key":"-death_pos.x"}, {"text":", "}, {"key":"-death_pos.y"}]], 
			"playtime_str": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-playtime_str"], 
			"projectiles_shot": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-projectiles_shot"], 

			"enemies_killed": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-enemies_killed"], 
			"dead": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-dead"], 
			"death_count": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-death_count"], 
			"killed_by": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-killed_by"], 
			"killed_by_extra": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-killed_by_extra"], 

			"gold": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-gold"], 
			"gold_all?": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-gold_all"], 
			"healed": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-healed"], 
			"places_visited": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-places_visited"], 

			"items": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-items"], 
			"streaks": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-streaks"], 
			"hp": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"stats", "Stats", "stats", "-hp"], 
			
			"kills": ["noitaDataRaw", "Sessions", {"key": "$Gameid", "type":"object"},"kills", "Stats", "-kills"], 
			
		},

		Stats: {
			"*": ["noitaDataRaw", "Stats", {"key": "$KEY", "type":"object"}, {"template": "$Stats"}],
		},
		"$Stats":{ 
			"*": ["Stats", [
				{"text":"Deaths: "}, {"key":"-deaths"}, {"text":", "}, 
				{"text":"kills: "}, {"key":"-kills"}, {"text":", "}, 
				{"text":"player_kills: "}, {"key":"-player_kills"}, {"text":", "}, 
				{"text":"player_projectile_count: "}, {"key":"-player_projectile_count"}
			]]
		},
		Stats_: {
			"*": ["noitaDataRaw", "Stats", {"key": "$KEY", "type":"object"}, {"template": "$Stats_"}],
		},
		"$Stats_": {
			"Deaths": ["Stats", "-deaths"],
			"kills": ["Stats", "-kills"],
			"player_kills": ["Stats", "-player_kills"],
			"player_projectile_count": ["Stats", "-player_projectile_count"],
		},

		Perks: {
			"*": ["noitaDataSorted", {"key": "Perks", "type":"array"}, {"template": "$Perks"}],
		},
		"$Perks": {
			"name": ["-name"],
			"tags": ["-tags"],
			"effect": ["GameEffectComponent","-effect"],
			"_tags": ["GameEffectComponent","-_tags"],
			"teleportation_probability": ["GameEffectComponent","-teleportation_probability"],

			"description": ["UIIconComponent","-description"],
			"icon_sprite_file": ["UIIconComponent","-icon_sprite_file"],
			"name": ["UIIconComponent","-name"],
			//"effect": ["UIIconComponent","-effect"],
		},

		
		Wands: {
			"*": ["noitaDataSorted", {"key": "Wands", "type":"array"}, {"template": "$Wands"}]
		}, 
		"$Wands": {
			"Tags": [ "-tags"	],
			"Wand level": [	"AbilityComponent", "-gun_level"	],
			"image file": [	"SpriteComponent", "-image_file"	],
			"damage_multiplier ": [ "HitboxComponent", "-damage_multiplier"],
			"mana": [ "AbilityComponent", "-mana"],
			"mana_max": [ "AbilityComponent", "-mana_max"],
			"mana_charge_speed": [ "AbilityComponent", "-mana_charge_speed"],
			"reload_time_frames": ["AbilityComponent", "-reload_time_frames"],
			"stat_times_player_has_shot": [ "AbilityComponent", "-stat_times_player_has_shot"],
			"ui_name": [ "AbilityComponent", "-ui_name"],
			"stat_times_player_has_edited": ["AbilityComponent", "-stat_times_player_has_edited"],
			"sprite_file": [ "AbilityComponent", "-sprite_file"],
			
			"speed_multiplier": [ "AbilityComponent", "gunaction_config", "-speed_multiplier"],
			"bounces": [ "AbilityComponent", "gunaction_config", "-bounces"],
			"blood_count_multiplier": [ "AbilityComponent", "gunaction_config", "-blood_count_multiplier"],
			"fire_rate_wait": [ "AbilityComponent", "gunaction_config", "-fire_rate_wait"],
			"action_mana_drain": [ "AbilityComponent", "gunaction_config", "-action_mana_drain"],
			"spread_degrees": [ "AbilityComponent", "gunaction_config", "-spread_degrees"],
			"action_max_uses": [ "AbilityComponent", "gunaction_config", "-action_max_uses"],
			
			"inventory_slot.x": [ "ItemComponent", "-inventory_slot.x"],
			"inventory_slot.y": [ "ItemComponent", "-inventory_slot.y"],
			"spawn_pos.x": [ "ItemComponent", "-spawn_pos.x"],
			"spawn_pos.y": [ "ItemComponent", "-spawn_pos.y"],
			"spawn_pos": [ "ItemComponent", [{"key":"-spawn_pos.x"}, {"text":", "}, {"key":"-spawn_pos.y"}]],
			"Spells": {
				"*": [ {"key": "Entity", "type":"array"}, {"template": "$Spells"}]
			}
		},
		Spells: {
			"*": ["noitaDataSorted", {"key": "Spells", "type":"array"}, {"template": "$Spells"}], 
		},
		"$Spells": {
			"ID": ["ItemActionComponent", "-action_id"],
			"Tags": ["-tags"], 
			"Image file": ["SpriteComponent", "0", "-image_file"],
			"Air friction": ["VelocityComponent", "-air_friction"],
			"gravity_x": ["VelocityComponent", "-gravity_x"],
			"gravity_y": ["VelocityComponent", "-gravity_y"],
			"liquid_drag": ["VelocityComponent", "-liquid_drag"],
			"affect_physics_bodies": ["VelocityComponent", "-affect_physics_bodies"],
			"displace_liquid": ["VelocityComponent", "-displace_liquid"],
			"mass": ["VelocityComponent", "-mass"],
			"terminal_velocity": ["VelocityComponent", "-terminal_velocity"],
			"damage_multiplier": ["HitboxComponent", "-damage_multiplier"],
			"spawn_pos.x": ["ItemComponent", "-spawn_pos.x"],
			"spawn_pos.y": ["ItemComponent", "-spawn_pos.y"],
			"uses_remaining": ["ItemComponent", "-uses_remaining"],
			"item_pickup_radius": ["ItemComponent", "-item_pickup_radius"],
			"can_go_up": ["SimplePhysicsComponent", "-can_go_up"],
			"inventory_slot.x": ["ItemComponent", "-inventory_slot.x"],
			"inventory_slot.y": ["ItemComponent", "-inventory_slot.y"],
		},

		Test: {
			//"*": ["noitaDataSorted", {"key": "Spells", "type":"array"}, {"template": "$Spells"}], 
			//"*": ["noitaDataSorted", "Spells", {"template": "$Spells"}], 
			"*": ["noitaDataSorted", "Wands", {"template": "$Wands"}]
		},
		Test2: {
			"ID": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "ItemActionComponent", "-action_id"],
			"Tags": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "-tags"], 
			"Image file": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "SpriteComponent", "0", "-image_file"],
			"Air friction": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-air_friction"],
			"gravity_x": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_x"],
			"gravity_y": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_y"],
		},
		Test3: {
			"bla*":{
				"ID": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "ItemActionComponent", "-action_id"],
				"Tags": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "-tags"], 
				"Image file": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "SpriteComponent", "0", "-image_file"],
				"Air friction": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-air_friction"],
				"gravity_x": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_x"],
				"gravity_y": ["noitaDataSorted", {"key": "Spells", "type":"array"}, "VelocityComponent", "-gravity_y"],
			}
		},
		
	}
	console.log('**************summary*************', allData.noitaDataSorted)
	layoutFresh = layout
	//summary3 = parseLayout(layout, allData)
	//console.log('summary3', summary3)
	summary3 = getToKey2(layout, allData)
	return summary3
}


  function getToKey2(val, summary){
	console.log('****getToKey2****', val, summary)
	if(false){}

	else if(Array.isArray(val) || typeof val == 'string'){
		console.log('!!!!Array || string!!!!!', val)
		val = (typeof val == 'string')? [val] : val
		let k = val[0]
		val.shift()
		if(k){
			return getToKey2(val, summary[k])
		} 
		else {
			console.log('getToKey DONE1', summary)
			console.log('getToKey DONE2', {[keyName]: summary})
			//return {[keyName]: summary}
			//return summary[k]
			console.log('!!!!DONE0!!!!!', val, summary)
			return summary
		}
	}
	
	else if(typeof val == 'object'){
		console.log('!!!!object!!!!!', val)
		let ent = Object.entries(val)
		let res = {}
		for (const [key, value] of ent) {
			console.log('---object---', key, value)
			if(key == '$') continue
			res[key] = getToKey2(value, summary)
		}
		if(val['$'] && val['$'].array){
			res = Object.values(res)
		}
		return getToKey2(res, summary)
	}


	console.log('!!!!DONE1!!!!!', val, summary)
	return summary
}