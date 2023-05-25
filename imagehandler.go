package main

import (
	"fmt"
	"log"
	"strings"
	"os"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
	"noitastats/utils"
	//"reflect"
	"path/filepath"
	"encoding/base64"
	"io/ioutil"
	"net/http"
)
func _() {
	fmt.Println("")
}
type SpriteStruct struct  {
	//XMLName xml.Name `xml:"Sprite"`
	Filename string `xml:"filename,attr"`
}
var Images map[string]string

func ImageInit() error{
	Images = map[string]string{}
	paths := []string{
		os.Getenv("userprofile") + "/appdata/locallow/Nolla_Games_Noita" + "/data/ui_gfx/", 
		os.Getenv("userprofile") + "/appdata/locallow/Nolla_Games_Noita" + "/data/items_gfx/", 
		//os.Getenv("userprofile") + "/appdata/locallow/Nolla_Games_Noita" + "/data/ui_gfx/gun_actions/", 
	}
	for i := 0; i < len(paths); i++ {
		path := utils.FixPath(paths[i])
		if(path == ""){
			continue
		}
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			if(!info.IsDir() && !strings.Contains(path, ".xml")){
				getImage(path)
			}
			return nil
		})
	}
	return nil
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}


func getImage(url string) string{
	if _, err := os.Stat(url); err != nil {
		return ""
	}
	bytes, err := ioutil.ReadFile(url)
	if err != nil {
		bytes = GetImageFilenameFromXml(url)
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += toBase64(bytes)
	url = strings.Replace(url, os.Getenv("userprofile") + "\\appdata\\locallow\\Nolla_Games_Noita\\", "", 1)
	url = strings.ReplaceAll(url, "\\", "/")
	
	return base64Encoding
}

func GetImageFilenameFromXml(url string) []byte{
	splitted := strings.Split(url, ".")
	if(splitted[len(splitted)-1] != "xml"){
		splitted[len(splitted)-1] = "xml"
		url = strings.Join(splitted, ".")
		str := utils.File(url)
		if(str == ""){
			//return ""
		}
		Sprite := SpriteStruct{}
		utils.XmlUnmarshal(&Sprite, str)
		url = GetPath(Config.Paths.Nolla_Games_Noita) + "\\" + Sprite.Filename
		url = utils.FixPath(url)
		file := ReadFile(url)
		return file
	}
	return nil
}


func ReadFile(url string) []byte{
	bytes, err := ioutil.ReadFile(url)
	if err != nil {
		//fmt.Println("ReadFile err", err)
	}
	return bytes
}

func Base64Encode(bytes []byte) string{
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}