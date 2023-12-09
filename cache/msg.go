package cache

import (
	"encoding/json"
	"time"
)

// Message represents a structure for caching messages.
type Message struct {
	Detail string `json:"detail"`
	Device string `json:"device"`
	Time   string `json:"time"`
}

// CreateMessage creates a new Message with the provided details and device.
func CreateMessage(detail, device string) (Message, error) {
	msg := Message{
		Device: device,
		Detail: detail,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
	}
	return msg, nil
}

// AddMessageToList adds a Message to a list in Redis, associating it with the provided key.
func AddMessageToList(key string, message Message) error {
	// Convert the message to JSON.
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Push the JSON-formatted message to the Redis list.
	_, err = RedisClient.LPush(key, messageJSON).Result()
	if err != nil {
		return err
	}

	// Set an expiration time for the key in Redis (24 hours in this case).
	_, err = RedisClient.Expire(key, 24*time.Hour).Result()
	return err
}

// GetMessagesFromSortedSet retrieves messages from a Redis list associated with the provided key.
func GetMessagesFromSortedSet(key string) ([]Message, error) {
	// Retrieve JSON-formatted messages from the Redis list.
	messageJSONs, err := RedisClient.LRange(key, 0, 0).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal each JSON-formatted message and append it to the messages slice.
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

// GetMessagesFromSortedSetLimit retrieves a limited number of messages from a Redis list associated with the provided key.
func GetMessagesFromSortedSetLimit(key string, limit int) ([]Message, error) {
	// Retrieve JSON-formatted messages with a specified limit from the Redis list.
	messageJSONs, err := RedisClient.LRange(key, 0, int64(limit)).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal each JSON-formatted message and append it to the messages slice.
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
