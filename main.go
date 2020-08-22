package main

import (
	"Tracking/config"
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
	defer func() { select {} }()

	config := Config.Config{
		DeviceCachePath: os.Getenv("LOCALAPPDATA") + "\\Oculus\\DeviceCache.json",
		Url:             "/SensorStatus",
		Host:            "127.0.0.1:23333",
		DeviceId:        map[string]int{},
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
}

func listen(watch fsnotify.Watcher, target *server.AllDeviceInfo, config Config.Config) {
	for {
		select {
		case ev := <-watch.Events:
			{
				//if ev.Op&fsnotify.Create == fsnotify.Create {
				//	log.Println("创建文件 : ", ev.Name)
				//}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					readjson(config.DeviceCachePath, target)
					//log.Println("写入文件 : ", ev.Name)
				}
				//if ev.Op&fsnotify.Remove == fsnotify.Remove {
				//	log.Println("删除文件 : ", ev.Name)
				//}
				//if ev.Op&fsnotify.Rename == fsnotify.Rename {
				//	log.Println("重命名文件 : ", ev.Name)
				//}
			}
		case err := <-watch.Errors:
			{
				log.Println("error : ", err)
				return
			}
		}
	}
}

//func getEnv(){
//	if runtime.GOOS == "windows"{
//		os.Getenv("LOCALAPPDATA")
//	}
//}
