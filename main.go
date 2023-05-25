package main

import (
	"embed"
	"context"
	//"fmt"
	"github.com/wailsapp/wails/v2"
	//"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
    _ "github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS
var app *App
func main() {
	createWindow()
}
func createWindow(){
	app = NewApp()
	err := wails.Run(&options.App{
		Title:  "noitastats",
		Width:  1024,
		Height: 768,
		Frameless:          false,
		AlwaysOnTop:        false,
		BackgroundColour: &options.RGBA{R: 25, G: 25, B: 25, A: 0},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  func(ctx context.Context){
			app.startup(ctx)
        },
		OnBeforeClose:      app.beforeClose,
		CSSDragProperty: "--wails-draggable",
        CSSDragValue: "drag",
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
            WebviewIsTransparent: false,
            WindowIsTranslucent: false,
			DisableWindowIcon: true,
			DisableFramelessWindowDecorations: false,
		},
		Linux: &linux.Options{
            WindowIsTranslucent: false,
        },
		Debug: options.Debug{
            OpenInspectorOnStartup: true,
        },

	})
	if err != nil {
		println("Error:", err.Error())
	}
}
