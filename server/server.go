package server

import (
	Config "Tracking/config"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type AllDeviceInfo struct {
	Devices []map[string]interface{} `json:"devices"`
}

type DataPack struct {
	Data map[string]interface{}
}

var page string
var config Config.Config
var target *AllDeviceInfo

func InitEcho(t *AllDeviceInfo, Setting Config.Config) {
	e := echo.New()
	config = Setting
	target = t
	api := "/api/" + config.Url

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
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	var dict = make(map[string]interface{})
	var data DataPack

	if config.DeviceCachePath == "" {
		println("not valid path set!")
		os.Exit(0)
	}

	for _, v := range target.Devices {

		if v["type"] != "sensor" {
			continue
		}
		for Uid, Uintid := range config.DeviceId {
			if Uid == v["id"] {
				i := strconv.Itoa(Uintid)
				dict[i] = v["visibility"].(map[string]interface{})
			}
		}
	}

	data.Data = dict

	return c.JSON(http.StatusOK, data)
}
