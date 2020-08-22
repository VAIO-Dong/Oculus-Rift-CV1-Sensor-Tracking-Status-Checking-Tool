package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	DeviceCachePath string         `json:"path"`
	Url             string         `json:"url"`
	Host            string         `json:"host"`
	DeviceId        map[string]int `json:"device_id"`
}

func NewConfig(Config Config) {
	f, _ := os.Create("./config.cfg")

	fmt.Printf("write path into config\n")
	fmt.Println("if over, please inter:exit or EXIT")
	for i := 1; i <= 4; i++ {

		device_id := ""
		fmt.Printf("Enter the device id for the ID %d :", i)
		_, _ = fmt.Scanln(&device_id)

		if strings.ToUpper(device_id) == "EXIT" {
			println("exit")
			break
		}
		if len(device_id) != 14 {
			log.Fatalln("Is the format wrong?")
		}

		Config.DeviceId[device_id] = i
	}
	b, _ := json.Marshal(Config)
	_, _ = f.Write(b)
	os.Exit(0)
}
