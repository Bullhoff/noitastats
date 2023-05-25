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
	"io/ioutil"
	"time"
	"errors"
)
import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"context"
)
import (
	"noitastats/utils"
)

func _() {
	fmt.Println("qol")
	log.Println("qol")
	reflect.TypeOf("str")
	strings.ReplaceAll("str", "\\", "/")
	os.Getenv("userprofile")
	_, _ = os.Stat("")
	var mutex = &sync.Mutex{}
	mutex.Lock()
	mutex.Unlock()
	strconv.Atoi("123")
	filepath.Walk("c:/aaaa", func(path string, info os.FileInfo, err error) error {		
		if err != nil { log.Fatalf(err.Error());  }		
		if(!info.IsDir()){}		
		return nil;	
	})
	_, _ = json.Marshal(`{"key":"val"}`)
	ioutil.ReadFile("url")
	time.Sleep(time.Duration(time.Second * 1))
	errors.New("err")
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

var globalContext context.Context
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	globalContext = a.ctx
	runtime.LogPrint(a.ctx, "message string")
	runtime.EventsOn(ctx, "client-event", callback) 
}

//func (b *App) domready(ctx context.Context) (prevent bool) {	return true}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	done := make(chan bool, 1)
	x, y := runtime.WindowGetPosition(a.ctx)
	fmt.Println("WindowGetPosition", x, y)
	width, height := runtime.WindowGetSize(a.ctx)
	fmt.Println("WindowGetSize", width, height)
	fmt.Println("ctx", ctx)
	fmt.Println("a", a)

	runtime.EventsOnce(globalContext, "client", func(optionalData ...interface{}) {
		config, ok := optionalData[0].(map[string]interface{})["config"] 
		if(ok){
			str := utils.Stringify(config)
			conf := ConfigStruct{}
			json.Unmarshal([]byte(str), &conf)
			
			remember_pos := conf.Window.Remember_pos
			remember_size := conf.Window.Remember_size
			fmt.Println("remember_pos", remember_pos)
			fmt.Println("remember_size", remember_size)
			if(remember_pos){
				conf.Window.Pos_x = x	//Config.Window.Pos_x
				conf.Window.Pos_y = y	//Config.Window.Pos_Y
			}
			if(remember_size){
				conf.Window.Size_x = width 	//Config.Window.Size_x
				conf.Window.Size_y = height	//Config.Window.Size_Y
			}
			jsonstr, _ := json.MarshalIndent(conf, "", "\t")
			fmt.Println("DONE", string(jsonstr))
			utils.WriteFile(GetPath(Config.Paths.NoitaStats_configjson), string(jsonstr))
		} 
	})
	runtime.EventsEmit(globalContext, "server-event", "exiting" )
	time.Sleep(1 * time.Second)
	done <- true
	<-done
	return false
	
}


func (a *App) SaveJson(path, filename string, data interface{}) error {
	jsonstr, errJson := json.MarshalIndent(data, "", "\t")
	if(errJson != nil){
		return errJson
	}
	path = GetPath(path) + filename
	errWrite := utils.WriteFile(path, string(jsonstr))
	if(errWrite != nil){
		return errWrite
	}
	return nil
}


func (a *App) GetNoitaDataFile(dataurl string) string{	
	path := GetPath(Config.Paths.Nolla_Games_Noita) + dataurl
	exists, _ := utils.PathExists(path)
	fmt.Println("GetNoitaDataFile 1", exists, path)
	if(!exists){
		return ""
	}
	res, _ := ioutil.ReadFile(path)
	return string(res)
}




func (a *App) GetNoitaData(gameId string) map[string]interface{}{
	return getSpecificGame(gameId)
	if(gameId == "latest"){
		return getSpecificGame(gameId)
		//return NoitaData
	} else {
		return getSpecificGame(gameId)
	}
	
}
func (a *App) GetAvailabeFiles() map[string]FileStruct{
	return AvailabeFiles
}



func (a *App) ExecCommand(p map[string]string) {
	// entry, ok := p["opt"]
	opt, _ := p["opt"]
	fmt.Println("ExecCommand", utils.Stringify(p))

	if(opt == "open_config"){
		exec.Command("cmd", "/c", "start", GetPath(Config.Paths.NoitaStats_configjson)).Start()	//.Run()
	}
}

func (a *App) UpdateSessionJson(id string, data SessionJsonStruct) {
	NoitaStats_sync_to := GetPath(Config.Paths.NoitaStats_sync_to)
	path := NoitaStats_sync_to + id + `\data.json`
	utils.CreateDirectory(NoitaStats_sync_to + `` + id)
	utils.WriteFile(path, utils.StringifyIndent(data))
	//jsonContent := utils.File(path)
}
func (a *App) SaveComment(id, text string) {
	path := GetPath(Config.Paths.NoitaStats_sync_to) + id + `\data.json`
	jsonContent := utils.File(path)
	data := map[string]interface{}{}
	utils.JsonUnmarshal(jsonContent, &data)
	data["text"] = text
	utils.WriteFile(path, utils.Stringify(data))
}

func (a *App) GetDataXmlFiles() map[string]Entity_Struct{
	path := GetPath(Config.Paths.Noita_data_folder) + "/entities/"
	path = utils.FixPath(path)
	exists, _ := utils.PathExists(path)
	if(!exists){
		return map[string]Entity_Struct{}
	}
	dataEntityList := map[string]Entity_Struct{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
            log.Fatalf(err.Error())
        }
		if(!info.IsDir()){
			split := strings.Split(path, ".")
			if(split[len(split)-1] != "xml"){
				return nil
			}
			path_str := getPath_str(path)
			path_str = strings.ReplaceAll(path_str, `\`, `/`)

			xmlres, _ :=  dataEntityList[path_str]
			str := utils.File(path)
			utils.XmlUnmarshal(&xmlres, str)
			dataEntityList[path_str] =  xmlres
		}
		return nil
	})
	runtime.EventsEmit(globalContext, "server-event", map[string]map[string]Entity_Struct{"DataEntityList": dataEntityList} )
	return dataEntityList
}
func (a *App) GetImages() map[string]string{
	return Images
}

func (a *App) GetImage(imgsrc string) string{
	img := Images[imgsrc]
	if(img != ""){
		return img
	}
	url := GetPath(Config.Paths.Nolla_Games_Noita) + "\\" + imgsrc
	url = utils.FixPath(url)
	img = getImage(url)
	return img
}


func (a *App) DeleteTemplate(filename string) {
	//utils.Delete(bla(Paths["templates"])+"\\"+filename)
}
func (a *App) SaveTemplate(filename, data string, ) {
	//utils.WriteFile(bla(Paths["templates"])+"\\"+filename, data)
}

func (a *App) GetConfig(str string) string{
	return utils.File(GetPath(Config.Paths.NoitaStats_configjson))
}



func (a *App) UnpackAssets() error{ // map[string]string
	
	noitaPath := GetPath(Config.Paths.Noita)
	if(noitaPath == ""){
		fmt.Println("Noita install direcotry not found. Navigate to folder or unpack data.wak manually", noitaPath)
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:        "Unpack data.wak",
			Message:      "Cant find data.wak\nChange the 'noita' path in the config file or unpack according to the official instructions",
		})
		return nil
	}

	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Unpack data.wak",
		Message:       "This will run: \n"+ "\"cd " + noitaPath + " && noita.exe -wizard_unpak\" \n" +"Which will unpack data.wak (≈40 MB) to: \n"+GetPath(Config.Paths.Nolla_Games_Noita)+"data\\ \n\n aight?",
		DefaultButton: "Yes",
		CancelButton:  "No",
	})
	if(result == "Yes"){
		cmd := exec.Command("cmd", "/c", "cd",noitaPath, "&&", "noita.exe", "-wizard_unpak")
		cmd.Run()
		return nil
	}
	return nil
}


type Params struct {
	opt, str string
}

func (a *App) StartExplorer(p map[string]string) error{	// p *Params		// map[string]string	map[string]interface{}

	if(p["opt"] == "start_datafolder"){
		//exec.Command("explorer", bla(Paths["cwd"])).Start()
		exec.Command("explorer", GetPath(Config.Paths.NoitaStats)).Start()
	}
	if(p["opt"] == "save00"){
		exec.Command("explorer", GetPath(Config.Paths.Save00)).Start()
	}
	if(p["opt"] == "save_rec"){
		exec.Command("explorer", GetPath(Config.Paths.Nolla_Games_Noita)+`save_rec`).Start()
	}
	if(p["opt"] == "noitapath"){
		exec.Command("explorer", GetPath(Config.Paths.Noita)).Start()
	}
	if(p["opt"] == "startnoita"){
		exec.Command("explorer", "steam://rungameid/881100").Start()
	}
	if(p["opt"] == "showfile"){
		basePath := p["basePath"]
		path := p["path"]
		fmt.Println("showfile1", basePath, path)
		if(basePath != "" && path != ""){
			//fullpath := bla(Paths[basePath]) + path
			//fullpath = utils.Fix(fullpath)
			//fmt.Println("showfile2", fullpath)
			//exec.Command("explorer.exe", "/select,"+fullpath).Start()
			//fmt.Println("showfile3", "\\\\", "\\", string(filepath.Separator))
		}
	}
	return nil	
}


func (a *App) WindowAction(p map[string]interface{}){	// p map[string]string
	if(p == nil){
		return
	}
	opt := bla(p["opt"])
	value := bla(p["value"])
	x, _ := strconv.Atoi(bla(p["x"]))
	y, _ := strconv.Atoi(bla(p["y"]))

	if(opt == "WindowSetTitle"){		// Go: WindowSetTitle(ctx context.Context, title string)		// JS: WindowSetTitle(title: string)
		runtime.WindowSetTitle(a.ctx, "title string")
		fmt.Println("WindowAction", opt)
	}
	if(opt == "WindowGetSize"){		// Go: WindowGetSize(ctx context.Context) (width int, height int)		// JS: WindowGetSize() : Size
		x, y := runtime.WindowGetSize(a.ctx)
		fmt.Println("WindowAction", opt, bla(x), bla(y))
	}
	if(opt == "WindowSetSize"){		// Go: WindowSetSize(ctx context.Context, width int, height int)		// JS: WindowSetSize(size: Size)
		runtime.WindowSetSize(a.ctx, x, y)
		fmt.Println("WindowAction", opt, bla(x), bla(y))
	}
	if(opt == "WindowGetPosition"){		// Go: WindowGetPosition(ctx context.Context) (x int, y int)		// JS: WindowGetPosition() : Position
	}
	if(opt == "WindowSetPosition"){		// Go: WindowSetPosition(ctx context.Context, x int, y int)		// JS: WindowSetPosition(position: Position)
	}
	if(opt == "WindowSetAlwaysOnTop"){		// Go: WindowSetAlwaysOnTop(ctx context.Context, b bool)		// JS: WindowSetAlwaysOnTop(b: Boolen)
		b := p["b"].(bool)
		runtime.WindowSetAlwaysOnTop(a.ctx, b)
	}
	if(opt == "WindowFullscreen"){		// Go: WindowFullscreen(ctx context.Context)		// JS: WindowFullscreen()
	}
	if(opt == "WindowMinimise"){		// Go: WindowMinimise(ctx context.Context)		// JS: WindowMinimise()
	}
	if(opt == "WindowShow"){		// Go: WindowShow(ctx context.Context)		// JS: WindowShow()
		runtime.WindowShow(a.ctx)
	}
	if(opt == "WindowHide"){		// Go: WindowHide(ctx context.Context)		// JS: WindowHide()
		runtime.WindowHide(a.ctx)
	}
	if(opt == "WindowReloadApp"){		// Go: WindowReloadApp(ctx context.Context)		// JS: WindowReloadApp()
		runtime.WindowReloadApp(a.ctx)
	}
	if(opt == "WindowExecJS"){		// Go: WindowExecJS(ctx context.Context, js string)		//fmt.Println("WindowAction", opt, value)
		runtime.WindowExecJS(a.ctx, value)
	}
	if(opt == "WindowCenter"){		//Go: WindowCenter(ctx context.Context)		//JS: WindowCenter()
		runtime.WindowCenter(a.ctx)
	}
	
	

	
	

	if(opt=="Popup"){
		//path, _ := os.Getwd()
		//cmd := exec.Command(path+"\\editor.exe", "-project=xxxx")
		
		cmd := exec.Command(`cmd`, "/c", "cd", `C:\Users\D\Documents\11C\Programming\Golang\workspace\popup\build\bin\`, "&&", "start", "popup.exe", "--flagname=6771").Run()	//"/c",  
		fmt.Println("WindowAction", opt, cmd)
	}
	//return utils.File(bla(Paths["config"]))
}