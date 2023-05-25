
package main

import (
	"fmt"
    "log"
	"os"
	"os/exec"
	"strings"
	"reflect"
	"path/filepath"
	"sync"
	"strconv"
	"io/ioutil"
)
import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"context"
)

import (	
	_ "time"
	_ "encoding/xml"
	json "encoding/json" 
	"noitastats/utils"
	"github.com/fsnotify/fsnotify"	
	"errors"
)

func _() {
	fmt.Println("qol")
	log.Println("qol")
	reflect.TypeOf("str")
	strings.ReplaceAll("str", "\\", "/")
	os.Getenv("userprofile")
	_, _ = os.Stat("")
	exec.Command("cmd", "/c", "start", "c:/m.txt").Start()	//.Run()
	var mutex = &sync.Mutex{}
	mutex.Lock()
	mutex.Unlock()
	strconv.Atoi("123")
	filepath.Walk("c:/aaaa", func(path string, info os.FileInfo, err error) error {
		if err != nil { log.Fatalf(err.Error())  }
		if(!info.IsDir()){}
		return nil
	})
	_, _ = json.Marshal(`{"key":"val"}`)
	ioutil.ReadFile("url")
}
func _(ctx context.Context) {
	runtime.Quit(ctx)
}
func _() {
	utils.Stringify(map[string]interface{}{})
}



func bla(v interface{}) string{
	return fmt.Sprintf("%v", v)
}

func FindPath(paths []string) (string, error) {
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(paths[i]); !os.IsNotExist(err) {
			return strings.ReplaceAll(paths[i], "/", `\`), nil
		}
	}
	return "", errors.New("Path not found") 
}

func NoitaInit() { 
	//fmt.Println("NoitaInit", )
	var m sync.Mutex
	Save00 := GetPath(Config.Paths.Save00)
	if(Save00 != ""){
		go NotifyDirectoryChange(Save00 + `\`, &m)
		go NotifyDirectoryChange(GetPath(Config.Paths.Stats) + "\\", &m)
		go NotifyDirectoryChange(GetPath(Config.Paths.Sessions) + "\\", &m)
		go NotifyDirectoryChange(Save00 + `\world\`, &m)
		readXml(GetPath(Config.Paths.Playerxml), "", &m)
		readXml(GetPath(Config.Paths.Worldstatexml), "", &m)
		iteratePath(GetPath(Config.Paths.Stats), &m)
	}
	iteratePath(GetPath(Config.Paths.NoitaStats_sync_to), &m)
}
func iteratePath(path string, m *sync.Mutex) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
            log.Fatalf(err.Error())
        }
		if(!info.IsDir()){
			filepathArr := strings.Split(path, "\\")
			foldername :=  filepathArr[len(filepathArr)-2]
			filename :=  filepathArr[len(filepathArr)-1]
			NoitaStats_sync_to := GetPath(Config.Paths.NoitaStats_sync_to) 
			if(strings.Contains(bla(path), ""+NoitaStats_sync_to) && filename == "data.json"){

			}
			if(strings.Contains(bla(path), ""+NoitaStats_sync_to)){
				checkFile(path, "init", m)
			} else if(foldername == "layouts"){

			} else if(foldername == "stats" || foldername == "sessions" ){
				readXml(path, "init", m)
			} else if(strings.Contains(bla(path), GetPath(Config.Paths.Noita_data_folder))){
				readXml(path, "init", m)
			}
		}
		return nil
	})
}


func getSpecificGame(gameid string) map[string]interface{}{
	fmt.Println("getSpecificGame 1 ", gameid)
	data := map[string]interface{}{}
	arr:= [2]string{"player.xml", "world_state.xml"}
	for i := 0; i < len(arr); i++ {
		path := GetPath(Config.Paths.NoitaStats_sync_to) + "" + gameid + `\` + arr[i]
		str := utils.File(path)
		Entity := Entity_Struct{}
		utils.XmlUnmarshal(&Entity, str)
		if(arr[i]=="player.xml"){
			if(len(Entity.GameStatsComponent) == 0) {
				data["Gameid"] = ""
				data["PlayerXml"] = Entity_Struct{}
				continue
			}
			data["Gameid"] = getGameId(Entity.GameStatsComponent[0].Stats_filename)
			data["PlayerXml"] = Entity
		} else {
			data["WorldStateXml"] = Entity
		}
		
	}
	return data
}
func checkDataJsonFile(path string, m *sync.Mutex) error{
	filepathArr := strings.Split(path, `\`)
	foldername :=  filepathArr[len(filepathArr)-2]
	filename :=  filepathArr[len(filepathArr)-1]

	NoitaStats_sync_to := GetPath(Config.Paths.NoitaStats_sync_to)
	if(strings.Contains(bla(path), ""+NoitaStats_sync_to) && filename == "data.json"){
		path := NoitaStats_sync_to + foldername + `\data.json`
		path = utils.FixPath(path)
		jsonContent := utils.File(path)
		if(jsonContent == ""){
			return nil
		}
		
		data := SessionJsonStruct{} 
		
		utils.JsonUnmarshal(jsonContent, &data)
		runtime.EventsEmit(globalContext, "server-event", map[string]map[string]SessionJsonStruct{"SessionJson": {foldername: data}} )
	}
	return nil
}
func checkFile(path string, Op string, m *sync.Mutex) string{
	checkDataJsonFile(path, m)
	fileStat, err := os.Stat(path)
	if err != nil {
		return ""
	}
	if(fileStat.IsDir() || !strings.Contains(path, ".xml") || fileStat.Size() == 0){
		return ""
	}
	path_str := getPath_str(path)
	m.Lock()
	entry, ok := AvailabeFiles[path_str]
	if(ok && entry.Size == fileStat.Size()){
		m.Unlock()
		return ""
	}
	entry.Name = fileStat.Name()
	entry.ModTime = fileStat.ModTime()
	entry.Size = fileStat.Size()
	AvailabeFiles[path_str] = entry
	runtime.EventsEmit(globalContext, "server-event", map[string]map[string]interface{}{"AvailabeFilesSingle": {path_str: AvailabeFiles[path_str]}} )	// AvailabeFiles[path_str]
	m.Unlock()

	return entry.Name
	return ""
}


func readXml(path, Op string, m *sync.Mutex) error{
	filename := checkFile(path, Op, m)
	if(filename == ""){
		return nil
	}
	str := utils.File(path)
	if(str == ""){
		return nil
	}

	path_str := getPath_str(path)
	split_ := strings.Split(filename, "_")

	if(strings.Split(path_str, `\`)[0] == "data" ){
		m.Lock()
		xmlres, _ :=  DataEntityList[path_str]
		utils.XmlUnmarshal(&xmlres, str)
		path_str = strings.ReplaceAll(path_str, `\`, `/`)
		DataEntityList[path_str] =  xmlres
		m.Unlock()
	} else if(filename == "_last_daily_run_played.xml"){
		utils.XmlUnmarshal(&LastDailyRunPlayed, str)
		runtime.EventsEmit(globalContext, "server-event", map[string]LastDailyRunPlayed_Struct{"LastDailyRunPlayed": LastDailyRunPlayed} )
	} else if(filename == "player.xml" ){
		PlayerXml = Entity_Struct{}
		utils.XmlUnmarshal(&PlayerXml, str)
		Gameid = getGameId(PlayerXml.GameStatsComponent[0].Stats_filename)
		//m.Lock()
		if(Op != "init"){
			runtime.EventsEmit(globalContext, "server-event", map[string]interface{}{"Gameid": Gameid} )
			runtime.EventsEmit(globalContext, "server-event", map[string]Entity_Struct{"PlayerXml": PlayerXml} )
		}
		copyFile(path, m)
		//m.Unlock()
	} else if(filename == "world_state.xml"){
		//m.Lock()
		WorldStateXml = Entity_Struct{}
		utils.XmlUnmarshal(&WorldStateXml, str)
		if(Op != "init"){
			runtime.EventsEmit(globalContext, "server-event", map[string]Entity_Struct{"WorldStateXml": WorldStateXml} )
		}
		copyFile(path, m)
		//m.Unlock()
	} else if(strings.Split(filename, "_")[0] == "stats"){
		m.Lock()
		xmlres, _ :=  StatsXml[filename] //Entity_Struct{}
		utils.XmlUnmarshal(&xmlres, str)
		StatsXml[filename] =  xmlres
		if(Op != "init"){
			runtime.EventsEmit(globalContext, "server-event", map[string]Stats_Struct{"StatsXml": StatsXml[filename]} )
		}
		m.Unlock()
	} else if(len(split_)>1 && (split_[1] == "stats.xml" || split_[1] == "kills.xml")){ 
		m.Lock() 
		gameid := strings.Split(filename, "_")[0]
		xmlres, _ :=  SessionsXml[gameid]
		utils.XmlUnmarshal(&xmlres, str)
		SessionsXml[gameid] =  xmlres
		if(Op != "init"){
			runtime.EventsEmit(globalContext, "server-event", map[string]map[string]Stats_Struct{"SessionsXml": map[string]Stats_Struct{gameid: SessionsXml[gameid]}} )
		}
		m.Unlock()
	}  
	return nil
}





func NotifyDirectoryChange(path string, m *sync.Mutex) error{
	readXml(path, "init", m) 

	m.Lock()
	entry, ok := watching[path]
	if(ok){
		m.Unlock()
		return nil
	}
	entry = true
	watching[path] = entry
	m.Unlock()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("NotifyDirectoryChange ERROR", watcher)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				Op := fmt.Sprintf("%v", event.Op)
				if(Op != "REMOVE" && (strings.Contains(event.Name, ".autosave_player") || strings.Contains(event.Name, ".autosave_world_pixel_scenes") || strings.Contains(event.Name, ".autosave_world_state"))){
					copyFile(event.Name, m)
				}	
				if(Op == "REMOVE" && strings.Contains(event.Name, "player.xml")){
					runtime.EventsEmit(globalContext, "server-event", map[string]string{"EndingGame": Gameid} )
					Gameid = ""
				}
				if(Op != "REMOVE" && strings.Contains(event.Name, ".xml")){
					if(Op != "WRITE" && Op != "REMOVE"){
						fmt.Println("**FILECHANGE**", event.Op, event.Name)
					}
					if(strings.Contains(event.Name, "config.xml"))	{
					} else {
						go readXml(event.Name, Op, m)
					}
					
				} else if(Op == "REMOVE"){}
				
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}()
	err = watcher.Add(path)
	if err != nil {
		log.Println("Add failed:", path, err)
	}
	<-done
	return nil
}


func copyFile(originalFile string, m *sync.Mutex)  error{
	if(Gameid == ""){
		return nil
	}
	newFile := GetPath(Config.Paths.NoitaStats_sync_to) + "\\" + Gameid + "\\" + strings.Replace(originalFile, GetPath(Config.Paths.NoitaStats_sync_to), "", -1)
	newFile = strings.Replace(newFile, GetPath(Config.Paths.Save00) + `world`, "", -1)
	newFile = strings.Replace(newFile, GetPath(Config.Paths.Save00), "", -1)
	//createPaths()
	utils.CreateDirectory(GetPath(Config.Paths.NoitaStats_sync_to) + "\\" + Gameid)
	_, err := utils.Copy(originalFile, newFile )
	if(err != nil) {
		//log.Println("!!copyFile!! delete 4 err", err)
	}
	return nil
}


func getPath_str(path string) string{
	path = strings.Replace(path, GetPath(Config.Paths.Save00), "", 1)
	path = strings.Replace(path, GetPath(Config.Paths.NoitaStats_sync_to), "", 1)
	path = strings.Replace(path, GetPath(Config.Paths.Nolla_Games_Noita), "", 1)
	
	if(path == "\\world_state.xml" || path == "\\player.xml" ){	
		path = ""+ Gameid + "" + path
	}
	if(path == "world_state.xml" || path == "player.xml" ){	
		path = ""+ Gameid + "\\" + path
	}
	return path
}

func getGameId(str string) string{
	str = str[strings.LastIndex(str, "/")+1:]
	str = strings.Split(str, "_")[0]
	return str
}
