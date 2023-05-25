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
	json "encoding/json" 
	xml "encoding/xml"
	"io/ioutil"
	"io"
	"time"
)
import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"context"
)

import (
	"noitastats/utils"
	_ "github.com/fsnotify/fsnotify"	
	_ "github.com/googollee/go-socket.io"
	_ "errors"
	
)
import (
	_ "archive/zip"
	_ "database/sql" 
	_ "modernc.org/sqlite"	
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
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	io.ReadAll(r)
	time.Now()
	xml.NewEncoder(os.Stdout)
}
func _(ctx context.Context) {
	runtime.Quit(ctx)
}
func _() {
	var ye = map[string]interface{}{}
	utils.Stringify(ye)
}

/* func ToStr[T any](v T) string{
	return fmt.Sprintf("%v", v)
} */

type PlayerStruct struct {
	GameStatsComponent struct  {
		Name string `xml:"name,attr"`
		Stats_filename string `xml:"stats_filename,attr"`	// "??STA/sessions/20230319-023419_kills.xml"
	}
}

var Paths map[string]interface{}
var watching = map[string]interface{}{}
var Gameid string
var LastDailyRunPlayed LastDailyRunPlayed_Struct
var PlayerXml = Entity_Struct{}
var WorldStateXml = Entity_Struct{}
var StatsXml = map[string]Stats_Struct{}
var SessionsXml = map[string]Stats_Struct{}
var AvailabeFiles = map[string]FileStruct{}

var Templates = map[string]interface{}{}
var EntityTemplate string
var EntityList = map[string]Entity_Struct{}
var DataEntityList = map[string]Entity_Struct{}

type FileStruct struct {
	Name string 
	ModTime time.Time 
	Size int64 
}
var mutex = &sync.Mutex{}

func GetPath(path string) string {
	path = strings.ReplaceAll(path, `\`, `/`)
	first := path[0:1]
	if(first == "/"){
		path, _ = filepath.Abs(path)
	}
	if(first == "%"){
		arr := strings.Split(path, "/")
		if(arr[0] == "%NoitaStats%"){
			arr[0] = Config.Paths.NoitaStats
			path = strings.Join(arr, `/`)
		}
		if(arr[0] == "%Nolla_Games_Noita%"){
			arr[0] = Config.Paths.Nolla_Games_Noita
			path = strings.Join(arr, `/`)
		}
	}
	path = utils.FixPath(path)
	return path
}
func DefaultsInit() {
	AvailabeFiles = map[string]FileStruct{}
	utils.CreateDirectory(GetPath(Config.Paths.NoitaStats_sync_to))
	//initVariables()
	//EntityTemplate = GetTemplate()
	NoitaInit()
	runtime.EventsEmit(globalContext, "server-event", map[string]interface{}{"AvailabeFiles": AvailabeFiles} )
	runtime.EventsEmit(globalContext, "server-event", map[string]map[string]Stats_Struct{"StatsXml": StatsXml})
	runtime.EventsEmit(globalContext, "server-event", map[string]map[string]Stats_Struct{"SessionsXml": SessionsXml})
	runtime.EventsEmit(globalContext, "server-event", map[string]interface{}{
		"LastDailyRunPlayed": LastDailyRunPlayed,
		"Stats": StatsXml, 
		"Config": Config, 
		//"EntityTemplate": EntityTemplate, 
	} )
}

func GetTemplate() string{
	entityTemplatePath := GetPath(Config.Paths.NoitaStats_templatejson)
	EntityTemplate = utils.File(entityTemplatePath)
	//fmt.Println("*GetTemplate 1*", EntityTemplate)
	if(EntityTemplate == ""){
		EntityTemplate = getDefaultEntityTemplate()
		utils.WriteFile(entityTemplatePath, EntityTemplate)
	}
	return EntityTemplate
}
func GetTemplates() map[string]interface{} {
	templates := map[string]interface{}{}
	return templates
	/* path := GetPath(Config.Paths.Templates)

	//fmt.Println("GetTemplates 1", path)
	//path := bla(Paths["cwd"]) + "\\noitadata\\templates\\"
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		//fmt.Println("GetTemplates 2", path)
		if err != nil {
            //fmt.Fatalf(err.Error())
        }
		filepathArr := strings.Split(path, "\\")
		//foldername :=  filepathArr[len(filepathArr)-2]
		filename :=  filepathArr[len(filepathArr)-1]
		filename = strings.Split(filename, ".")[0]

		//fmt.Println("GetTemplates 3", info)
		if(!info.IsDir()){
			//fmt.Println("GetLayouts", filename)
			//fmt.Println("**info", info)
			//fmt.Println("**path", path)

			//layouts["test"] = "abc"
			templates[filename] = utils.File(path)
			//str := utils.File(path)
		}
		return nil
	})
	return templates */
}

func getDefaultEntityTemplate() string{
	template :=  
	`[
		{"path":["-tags"], "config":{"name":null, "omitIfEmpty":false }},
		"Comment",
		{"path":["AbilityComponent", "-gun_level"], "config":{"name":"gun_level", "omitIfEmpty":true }}, 
		{"path":["AbilityComponent", "-reload_time_frames"], "config":{"name":null, "omitIfEmpty":true }}, 
		{"path":["AbilityComponent", "-mana_max"], "config":{"name":null, "omitIfEmpty":true }}, 
		{"path":["AbilityComponent", "-mana_charge_speed"], "config":{"name":null, "omitIfEmpty":true }}, 
		{"path":["AbilityComponent", "-stat_times_player_has_edited"], "config":{"name":null, "omitIfEmpty":true }}, 
		{"path":["AbilityComponent", "-stat_times_player_has_shot"], "config":{"name":null, "omitIfEmpty":true }},
		{"path":["*", "-stat_times_player_has_shot"], "config":{"name":"stat_times_player_has_shot", "omitIfEmpty":true }},

		{"_path":["-tags"], "config":{"name":null, "omitIfEmpty":false }},
		{"disabled":true, "path":["AbilityComponent", "-stat_times_player_has_shot"], "config":{"name":"moose", "omitIfEmpty":false }}
	]
	`
	return template
}
