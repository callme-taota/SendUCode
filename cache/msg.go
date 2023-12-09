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
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = RedisClient.LPush(key, messageJSON).Result()
	if err != nil {
		return err
	}

	_, err = RedisClient.Expire(key, 24*time.Hour).Result()
	return err
}

func GetMessagesFromSortedSet(key string) ([]Message, error) {
	messageJSONs, err := RedisClient.LRange(key, 0, 0).Result()
	if err != nil {
		return nil, err
	}

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
	messageJSONs, err := RedisClient.LRange(key, 0, int64(limit)).Result()
	if err != nil {
		return nil, err
	}

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
