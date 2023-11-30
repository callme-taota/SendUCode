package cache

import (
	"encoding/json"
	"time"
)

type Message struct {
	Detail string `json:"detail"`
	Device string `json:"device"`
	Time   string `json:"time"`
}

func CreateMessage(detail, device string) (Message, error) {
	msg := Message{
		Device: device,
		Detail: detail,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
	}
	return msg, nil
}

func AddMessageToList(key string, message Message) error {
	// 将消息序列化为 JSON 字符串
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// 使用 LPush 命令将消息添加到 List 的头部
	_, err = RedisClient.LPush(key, messageJSON).Result()
	if err != nil {
		return err
	}

	// 设置过期时间，例如设置为一天
	_, err = RedisClient.Expire(key, 24*time.Hour).Result()
	return err
}

func GetMessagesFromSortedSet(key string) ([]Message, error) {
	// 使用 ZRANGE 命令获取有序集合的所有元素
	messageJSONs, err := RedisClient.LRange(key, 0, 0).Result()
	if err != nil {
		return nil, err
	}

	// 解析 JSON 字符串为 Message 结构
	var messages []Message
	for _, messageJSON := range messageJSONs {
		var message Message
		err := json.Unmarshal([]byte(messageJSON), &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func GetMessagesFromSortedSetLimit(key string, limit int) ([]Message, error) {
	// 使用 ZRANGE 命令获取有序集合的所有元素
	messageJSONs, err := RedisClient.LRange(key, 0, int64(limit)).Result()
	if err != nil {
		return nil, err
	}

	// 解析 JSON 字符串为 Message 结构
	var messages []Message
	for _, messageJSON := range messageJSONs {
		var message Message
		err := json.Unmarshal([]byte(messageJSON), &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
