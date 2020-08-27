//go:generate goversioninfo -icon=./favicon.ico
package main

import (
	Config "Tracking/config"
	"Tracking/logger"
	"Tracking/server"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"os"
)

func readjson(file string, i interface{}) interface{} {
	f, _ := ioutil.ReadFile(file)
	_ = json.Unmarshal(f, i)
	return i

}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	logger.LogTOfile("./log.txt")
	defer func() {
		if err := recover(); err != nil {
			logger.PrintFatal(fmt.Sprintf("%+v", err))
			select {}
		}
	}()
	config := Config.Config{
		DeviceCachePath: os.Getenv("LOCALAPPDATA") + "\\Oculus\\DeviceCache.json",
		Url:             "/SensorStatus",
		Host:            "127.0.0.1:23333",
		DeviceId:        map[string]string{},
	}

	watch, _ := fsnotify.NewWatcher()

	if !IsExist("./config.cfg") {
		Config.NewConfig(config)
	} else {
		readjson("./config.cfg", &config)
		fmt.Printf("Config: %+v\n", config)
	}

	_ = watch.Add(config.DeviceCachePath)

	NewDevice := new(server.AllDeviceInfo)
	readjson(config.DeviceCachePath, &NewDevice)

	go server.InitEcho(NewDevice, config)
	go listen(*watch, NewDevice, config)

	logger.PrintInfo("init ok")
	select {}
}

func listen(watch fsnotify.Watcher, target *server.AllDeviceInfo, config Config.Config) {
	defer func() {
		if err := recover(); err != nil {
			logger.PrintFatal(fmt.Sprintf("%+v", err))
		}
	}()

	for {
		select {
		case ev := <-watch.Events:
			{

				if ev.Op&fsnotify.Write == fsnotify.Write {
					target.Mutex.Lock()
					readjson(config.DeviceCachePath, target)
					target.Mutex.Unlock()
					//log.Println("写入文件 : ", ev.Name)
				}
			}
		case err := <-watch.Errors:
			{
				log.Println("error : ", err)
				return
			}
		}
	}
}
