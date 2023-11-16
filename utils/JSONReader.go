package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"senducode/tolog"
)

func JSONReader(filePath string) (map[string]interface{}, error) {

	// 获取项目根目录路径
	rootDir, err := os.Getwd()
	if err != nil {
		tolog.WriteLog(fmt.Sprintln("jsonReader",err),"error")
		return nil, err
	}

	// 拼接文件路径
	absPath := filepath.Join(rootDir, filePath)

	fmt.Println("读取JSON文件:", absPath)

	// 读取文件内容
	fileContent, err := os.ReadFile(absPath)
	if err != nil {
		tolog.WriteLog(fmt.Sprintln("jsonReader",err),"error")
		return nil, err
	}

	// 解析 JSON
	var jsonData map[string]interface{}
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		tolog.WriteLog(fmt.Sprintln("jsonReader",err),"error")
		return nil, err
	}

	return jsonData, nil
}