package main

import (
	"encoding/json"
	"errors"
	"github.com/Sheyiyuan/ChronoMind/config"
	"github.com/Sheyiyuan/ChronoMind/logos"
	"io"
	"os"
)

func initCore() error {
	// 检查并创建必要的目录和文件
	if _, err := os.Stat("./conf/"); os.IsNotExist(err) {
		// 如果不存在，则创建该文件夹
		err := os.Mkdir("./conf/", 0755)
		if err != nil {
			logos.Fatal("初始化时，创建 conf/ 文件夹失败: %v", err)
		}
	}

	// 检查./conf/config.json是否存在
	if _, err := os.Stat("./conf/config.json"); os.IsNotExist(err) {
		// 如果不存在，则创建该文件
		file, err := os.Create("./conf/config.json")
		if err != nil {
			logos.Fatal("初始化时，创建 ./conf/config.json 配置文件失败: %v", err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				logos.Fatal("关闭 ./conf/config.json 配置文件失败: %v", err)
			}
		}(file)
	}

	// 检查并更新配置文件
	var globalConfig config.GlobalConfig

	// 读取配置文件
	file, err := os.Open("./conf/config.json")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logos.Fatal("无法关闭配置文件 ./conf/config.json: %v", err)
		}
	}(file)
	// 解码JSON配置
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&globalConfig)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}
	}

	// 检查并更新配置
	if globalConfig.HostPortConfig.Port < 1024 || globalConfig.HostPortConfig.Port > 65535 {
		globalConfig.HostPortConfig.Port = 8080
	}
	if globalConfig.LogConfig.LogLevel < 1 || globalConfig.LogConfig.LogLevel > 7 {
		globalConfig.LogConfig.LogLevel = 3
	}
	if globalConfig.AiApiConfig.AiWorkConfig.MaxTokens < 1 {
		globalConfig.AiApiConfig.AiWorkConfig.MaxTokens = 100
	}
	formattedJSON, err := json.MarshalIndent(globalConfig, "", "  ")
	if err != nil {
		return err
	}

	// 将格式化后的JSON字符串写入文件
	file, err = os.Create("./conf/config.json")
	if err != nil {
		logos.Fatal("初始化时，创建 ./conf/config.json 配置文件失败: %v", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logos.Fatal("无法关闭配置文件 ./conf/config.json: %v", err)
		}
	}(file)

	_, err = file.Write(formattedJSON)
	if err != nil {
		logos.Fatal("初始化时，写入 ./conf/config.json 配置文件失败: %v", err)
		return err
	}

	if _, err := os.Stat("./data/"); os.IsNotExist(err) {
		// 如果不存在，则创建该文件夹
		err := os.Mkdir("./data/", 0755)
		if err != nil {
			logos.Fatal("初始化时，创建 data/ 文件夹失败: %v", err)
		}
	}

	checkDataFolderExistence := func(dataAddress string) error {
		// 检查./data/文件夹中是否存在dataAddress文件夹
		if _, err := os.Stat(dataAddress); os.IsNotExist(err) {
			err := os.Mkdir(dataAddress, 0755)
			if err != nil {
				return err
			}
		}
		return nil
	}

	err = checkDataFolderExistence("./data/log/")
	if err != nil {
		logos.Fatal("创建日志文件夹 ./data/log/ 失败: %v", err)
		return err
	}

	// 创建日志文件
	file, err = os.OpenFile("./data/log/ChronoMind.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logos.Error("创建日志文件./data/ChronoMind.log 失败: %v", err)
		return err
	}
	return nil
}
