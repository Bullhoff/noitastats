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
	_ "modernc.org/sqlite"	//go get modernc.org/sqlite
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


var Config ConfigStruct

type ConfigStruct struct {
	General GeneralStruct `json:"general"`
	Paths PathsStruct `json:"paths"`
	Window WindowStruct `json:"window"`
	Highlight []HighlightStruct `json:"highlight"`
	State StateStruct `json:"state"`
}
type GeneralStruct struct {
	Debug bool `json:"debug"`
} 
type StateStruct struct {
	Sidemenu StateStructSub `json:"sidemenu"` 
	Session StateStructSub `json:"session"` 
	Alldata StateStructSub `json:"alldata"` 
	Lifetime_calculator StateStructSub `json:"lifetime_calculator"` 
	Debug StateStructSub `json:"debug"` 
}
type StateStructSub struct {
	Open bool `json:"open"` 
	X int `json:"x,omitempty"` 
	Y int `json:"y,omitempty"` 
	Width int `json:"width,omitempty"` 
	Height int `json:"height,omitempty"` 
	Sort_by string `json:"sort_by,omitempty"` 
	Display_filestatus bool `json:"display_filestatus,omitempty"`
	Display_input bool `json:"display_input,omitempty"`
	Display_extra bool `json:"display_extra,omitempty"`
	Subviews map[string]StateStructSub `json:"subviews,omitempty"` 
}
type Session_listStruct struct {
	//Display bool `json:"display"`
	Sort_by string `json:"sort_by,omitempty"`
	Display_filestatus bool `json:"display_filestatus,omitempty"`
	Display_input bool `json:"display_input,omitempty"`
	Display_extra bool `json:"display_extra,omitempty"`
	
} 

type WindowStruct struct {
	Always_on_top bool `json:"always_on_top"`
	Window_title string `json:"window_title"`
	Remember_pos bool `json:"remember_pos"`
	Pos_x int `json:"pos_x"`
	Pos_y int `json:"pos_y"`
	Remember_size bool `json:"remember_size"`
	Size_x int `json:"size_x"`
	Size_y int `json:"size_y"`
} 

type HighlightStruct struct {
	Active bool `json:"active"`
	Color string `json:"color"`
	String string `json:"string"`
}

type SessionJsonStruct struct {
	Favorite bool `json:"favorite"`	
	Comment string `json:"comment"`	
}
type SessionAdditionalStruct struct {
	Type string `json:"type"`
	Icon string  `json:"icon"`
	Text string  `json:"text"`
}

type NoitaStats_sub struct {
	Moose2 string
	Moose3 string
}
type PathsStruct struct {
	Noita string
	NoitaStats string 
	Nolla_Games_Noita string 
	NoitaStats_configjson string `json:"noitaStats_configjson"`
	NoitaStats_templatejson string `json:"noitaStats_templatejson"`
	NoitaStats_sync_to string `json:"noitaStats_sync_to"`
	Noita_data_folder string `json:"noita_data_folder"`
	Save00 string `json:"save00"`
	Playerxml string `json:"playerxml"`
	Worldstatexml string `json:"worldstatexml"`
	Stats string `json:"stats"`
	Sessions string `json:"sessions"`
}
func InitConfig(){
	fmt.Println("InitConfig")
	cwd, _ := os.Getwd() 
	// absPath, _ := filepath.Abs("../mypackage/data/file.txt")
	utils.CreateDirectory(cwd+`/data/`)
	config := utils.File(cwd+`/data/config.json`)	// GetPath(Config.Paths.NoitaStats_configjson)
	if(config == ""){
		DefaultConfig()
		jsonstr, _ := json.MarshalIndent(Config, "", "\t")
		utils.WriteFile(GetPath(Config.Paths.NoitaStats_configjson), string(jsonstr))
	} else {
		json.Unmarshal([]byte(config), &Config)
	}
}

func DefaultConfig() {	
	fmt.Println("DefaultConfig")
	
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	noitaInstallPath, _ := FindPath([]string{"C:/Program Files/Steam/steamapps/common/Noita/", "C:/Steam/steamapps/common/Noita/", "C:/Program Files (x86)/Steam/steamapps/common/Noita/"})
	
	Config = ConfigStruct{	
		General: GeneralStruct{
			//Disable_everything_except_filesync: false,
		},
		Paths: PathsStruct{
			Noita: noitaInstallPath, 
			Nolla_Games_Noita: utils.FixPath(os.Getenv("userprofile") + "/appdata/locallow/Nolla_Games_Noita/"),
			NoitaStats: utils.FixPath(cwd+`/data/`),

			NoitaStats_configjson: `%NoitaStats%/config.json`  ,
			NoitaStats_templatejson: `%NoitaStats%/template_entity.json` , 
			NoitaStats_sync_to: `%NoitaStats%/noitasaves/`  ,

			Noita_data_folder: "%Nolla_Games_Noita%/data/"  ,
			Save00: "%Nolla_Games_Noita%/save00/"  ,
			Playerxml: "%Nolla_Games_Noita%/save00/player.xml"  ,
			Worldstatexml: "%Nolla_Games_Noita%/save00/world_state.xml"  ,
			Stats: "%Nolla_Games_Noita%/save00/stats/"  ,
			Sessions: "%Nolla_Games_Noita%/save00/stats/sessions/"  ,
		},
		Window: WindowStruct{
			Always_on_top: false, 
			Window_title: "NoitaStats", 
			Remember_pos: true,
			//Pos_x: {}.(int), 
			//Pos_y: "",
			Remember_size: true,
			//Size_x: "", 
			//Size_y: "",
		},
		State: StateStruct{
			Sidemenu: StateStructSub{
				Open: true, 
				Sort_by: "datetime", 	
				Display_filestatus: true, 	
				Display_input: true,	
				Display_extra: true,	
				
			}, 
			Session: StateStructSub{
				Open: true, 
				Subviews: map[string]StateStructSub{
					"game_info": StateStructSub{Open: true}, 
					"wands": StateStructSub{Open: true}, 
					"spells": StateStructSub{Open: true}, 
					"items": StateStructSub{Open: true}, 
					"perks": StateStructSub{Open: true}, 
					"fungal_shifts": StateStructSub{Open: true}, 
					"data": StateStructSub{}, 
				} ,
			}, 
			Alldata: StateStructSub{
				Open: true,
			},
		}, 
		Highlight: []HighlightStruct{
			HighlightStruct{
				Active: false, 
				Color: "Chocolate", 
				String: "Text to highlight",
			}, 
			HighlightStruct{
				Active: false, 
				Color: "#D2691E", 
				String: "Text to highlight",
			}, 
		}, 
	
	}
}