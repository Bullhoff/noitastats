package main

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"fmt"
	"noitastats/utils"
	"reflect"
	"encoding/json"
)


func callback(optionalData ...interface{}){
	//fmt.Println("client-event 1", reflect.TypeOf(optionalData), reflect.TypeOf(optionalData[0]),  "optionalData")
	if(optionalData[0] == "init"){
		fmt.Println("client-event 2", optionalData[0])
		InitConfig()
		DefaultsInit()
	}
	if(false && reflect.TypeOf(optionalData[0]) == reflect.TypeOf(map[string]interface{}{})){	// map[string]interface{}
		//fmt.Println("!!reflect.TypeOf!!", reflect.TypeOf(optionalData[0]))
		gun_actions, ok := optionalData[0].(map[string]interface{})["gun_actions"]
		if(ok && false){
			str, _ := json.Marshal(gun_actions)
			json.Unmarshal(str, &EntityList)

			for key, element := range EntityList {
				fmt.Println("Key:", key, 
					element.GunActions[0].Related_extra_entities, 
					element.GunActions[0].Custom_xml_file, 
					element.GunActions[0].Related_projectiles, 
				)
				if(element.GunActions[0].Related_projectiles != nil){
					result := Entity_Struct{}
					file := utils.File(`C:\Users\D\AppData\LocalLow\Nolla_Games_Noita\` + bla(element.GunActions[0].Related_projectiles[0]))
					utils.XmlUnmarshal( &result, file,)
					fmt.Println("result", key, result.ProjectileComponent)
					entry, _ := EntityList[key]
					entry.ProjectileComponent = result.ProjectileComponent
					EntityList[key] = entry
					
				}
			}
			runtime.EventsEmit(globalContext, "server-event", map[string]map[string]Entity_Struct{
				"EntityList": EntityList,
			})
			
		}
		
	}
}