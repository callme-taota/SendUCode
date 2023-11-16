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
		tolog.Log().Context(fmt.Sprintln("jsonReader", err)).Type(tolog.ToLogStatusError).PrintLog().Write()
		return nil, err
	}

	// 拼接文件路径
	absPath := filepath.Join(rootDir, filePath)

	fmt.Println("读取JSON文件:", absPath)

	// 读取文件内容
	fileContent, err := os.ReadFile(absPath)
	if err != nil {
		tolog.Log().Context(fmt.Sprintln("jsonReader", err)).Type(tolog.ToLogStatusError).PrintLog().Write()
		return nil, err
	}

	// 解析 JSON
	var jsonData map[string]interface{}
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		tolog.Log().Context(fmt.Sprintln("jsonReader", err)).Type(tolog.ToLogStatusError).PrintLog().Write()
		return nil, err
	}

	return jsonData, nil
}

func JSONConvertToMapString(originalMap interface{}) map[string]string {
	convertedMap := make(map[string]string)
	// 遍历原始映射
	for key, value := range originalMap.(map[string]interface{}) {
		// 使用类型断言检查值的类型
		switch v := value.(type) {
		case int:
			// 将 int 转换为字符串
			convertedMap[key] = fmt.Sprintf("%d", v)
		case float64:
			// 将 float64 转换为字符串
			convertedMap[key] = fmt.Sprintf("%f", v)
		case string:
			// 字符串类型直接复制
			convertedMap[key] = v
		default:
			// 其他类型可以根据需要进行处理
			// 这里可以添加额外的类型转换规则
			fmt.Printf("Unsupported type for key %s\n", key)
		}
	}

	return convertedMap
}
