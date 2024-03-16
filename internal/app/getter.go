package app

import (
	"errors"
	"github.com/bh-qt/proxypool/log"

	"github.com/bh-qt/proxypool/internal/cache"
	"encoding/json"
	"github.com/ghodss/yaml"
	//"github.com/bh-qt/proxypool/pkg/tool"
	"github.com/bh-qt/proxypool/config"
	"github.com/bh-qt/proxypool/pkg/getter"
	"bufio"
	"os"
	"time"
	"fmt"
)

var Getters = make([]getter.Getter, 0)

func InitConfigAndGetters(path string) (err error) {
	err = config.Parse(path)
	if err != nil {
		return
	}
	if s := config.Config.SourceFiles; len(s) == 0 {
		return errors.New("no sources")
	} else {
		initGetters(s)
	}
	return
}

func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	json.Unmarshal([]byte(`{"url":"`+str+`"}`),&tempMap)
	return tempMap
}
func JsonToTgMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	json.Unmarshal([]byte(`{"channel":"`+str+`"}`),&tempMap)
	return tempMap
}
func JsonToyearMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	getYear := time.Now().Format("2006") //获取年
	getMonth := time.Now().Format("01") //获取月
	getDay  := time.Now().Day() //获取日
	json.Unmarshal([]byte(`{"url":"`+fmt.Sprintf(`%s%s/%s/%s%d.txt`, str, getYear,getMonth,getMonth,getDay)+`"}`),&tempMap)
	return tempMap
}

func initGetters(sourceFiles []string) {
	Getters = make([]getter.Getter, 0)
	for _, path := range sourceFiles {
		data, err := config.ReadFile(path)
		if err != nil {
			log.Errorln("Init SourceFile Error: %s\n", err.Error())
			continue
		}
		sourceList := make([]config.Source, 0)
		err = yaml.Unmarshal(data, &sourceList)
		if err != nil {
			log.Errorln("Init SourceFile Error: %s\n", err.Error())
			continue
		}
		for _, source := range sourceList {
			if source.Type == "clash" {
				file,err := os.Open("/root/proxypool/config/clash.txt")
				if err != nil {continue}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					text :=JsonToMap(scanner.Text())
					source.Options = text
					g, err := getter.NewGetter(source.Type,source.Options)
						if err == nil && g != nil {
							Getters = append(Getters, g)
							log.Debugln("init getter: %s %v", source.Type, scanner.Text())
						}
				}
			}else if source.Type == "subscribe1" {
				file,err := os.Open("/root/proxypool/config/subscribe.txt")
				if err != nil {continue}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					text :=JsonToMap(scanner.Text())
					source.Options = text
					g, err := getter.NewGetter(source.Type,source.Options)
						if err == nil && g != nil {
							Getters = append(Getters, g)
							log.Debugln("init getter: %s %v", source.Type, scanner.Text())
						}
				}
			}else if source.Type == "webfuzz1" {
				file,err := os.Open("/root/proxypool/config/webfuzz.txt")
				if err != nil {continue}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					text :=JsonToMap(scanner.Text())
					source.Options = text
					g, err := getter.NewGetter(source.Type,source.Options)
						if err == nil && g != nil {
							Getters = append(Getters, g)
							log.Debugln("init getter: %s %v", source.Type, scanner.Text())
						}
				}
			}else if source.Type == "tgchannel" {
				file,err := os.Open("/root/proxypool/config/tgchannel.txt")
				if err != nil {continue}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					text :=JsonToTgMap(scanner.Text())
					source.Options = text
					g, err := getter.NewGetter(source.Type,source.Options)
						if err == nil && g != nil {
							Getters = append(Getters, g)
							log.Debugln("init getter: %s %v", source.Type, scanner.Text())
						}
				}
			}else if source.Type == "webfuzznyr1" {
				file,err := os.Open("/root/proxypool/config/webfuzznyr.txt")
				if err != nil {continue}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					text :=JsonToyearMap(scanner.Text())
					source.Options = text
					//sType :=JsonToTypeMap()
					g, err := getter.NewGetter("subscribe",source.Options)
						if err == nil && g != nil {
							Getters = append(Getters, g)
							log.Debugln("init getter: %s %v", source.Type, scanner.Text())
						}
				}
			}
		}
	}
	log.Infoln("Getter count: %d", len(Getters))
	cache.GettersCount = len(Getters)
}
