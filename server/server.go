package server

import (
	Config "Tracking/config"
	"Tracking/logger"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

type AllDeviceInfo struct {
	Devices []map[string]interface{} `json:"devices"`
	Mutex   sync.Mutex
}

type DataPack struct {
	Data map[string]interface{}
}

var page string
var config Config.Config
var target *AllDeviceInfo

func InitEcho(t *AllDeviceInfo, Setting Config.Config) {
	defer func() {
		if err := recover(); err != nil {
			logger.PrintFatal(fmt.Sprintf("%+v", err))
		}
	}()

	e := echo.New()
	config = Setting
	target = t
	api := "/api" + config.Url

	InitPage(api)
	e.Static("/static", "static")

	e.GET(config.Url, GetIndex)
	e.GET(api, StatusApi)

	e.Logger.Fatal(e.Start(config.Host))
}

func GetIndex(c echo.Context) error {
	return c.HTML(http.StatusOK, page)
}

func InitPage(api string) {
	t, _ := template.ParseFiles("static/2SensorStatus.html")

	var htmlpageIO = bytes.Buffer{}
	if err := t.Execute(&htmlpageIO, api); err != nil {
		log.Println("err:", err.Error())
	}
	page = fmt.Sprint(htmlpageIO.String())

}

func StatusApi(c echo.Context) error {
	defer func() {
		if err := recover(); err != nil {
			logger.PrintFatal(fmt.Sprintf("%+v", err))
		}
	}()

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	var dict = make(map[string]interface{})
	var data DataPack

	if config.DeviceCachePath == "" {
		println("not valid path set!")
		os.Exit(0)
	}
	defer target.Mutex.Unlock()
	target.Mutex.Lock()

	for _, v := range target.Devices {

		switch v["subtype"] {
		case "unknown":
			if v["type"] != "sensor" {
				continue
			}

			for Uid, indexID := range config.DeviceId {
				if Uid == v["id"] {
					dict[indexID] = v["visibility"].(map[string]interface{})
					dict[indexID].(map[string]interface{})["connectionState"] = v["connectionState"]
				}
			}
		case "ltouch":
			dict["ltouch"] = map[string]interface{}{"connectionState": v["connectionState"], "batteryState": v["batteryState"]}
		case "rtouch":
			dict["rtouch"] = map[string]interface{}{"connectionState": v["connectionState"], "batteryState": v["batteryState"]}
		}
	}

	data.Data = dict

	return c.JSON(http.StatusOK, data)
}
