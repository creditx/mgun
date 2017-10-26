package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var file string
	flag.StringVar(&file, "f", "/path/to/config/file.yaml", "configuration yaml file")
	flag.Parse()

	if len(file) > 0 {
		bytes, err := ioutil.ReadFile(file)
		if err == nil {

			kill := GetKill()
			victim := NewVictim()
			gun := GetGun()
			reporter := GetReporter()

			err := yaml.Unmarshal(bytes, kill)
			if err == nil {
				err = yaml.Unmarshal(bytes, victim)
				if err == nil {
					err = yaml.Unmarshal(bytes, reporter)
					if err == nil {
						err = yaml.Unmarshal(bytes, gun)
						if err == nil {
							kill.SetVictim(victim)
							kill.SetGun(gun)
							kill.Prepare()
							kill.Start()
						} else {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("config file not found")
	}
}
