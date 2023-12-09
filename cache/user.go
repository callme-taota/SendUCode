package cache

import (
	"github.com/go-redis/redis"
	"senducode/tolog"
	"senducode/utils"
)

const reverseLookupHash = "reverse_lookup"
const forwardHash = "user_sessions"

type User struct {
	ID      string `json:"id"`
	Session string `json:"session"`
}

func AddUser(userID string) (string, error) {
	session, err := utils.CreateSession(userID)
	if err != nil {
		return "", err
	}

	err = RedisClient.HSet(forwardHash, userID, session).Err()
	if err != nil {
		return "", err
	}

	err = RedisClient.HSet(reverseLookupHash, session, userID).Err()
	if err != nil {
		return "", err
	}

	return session, err
}

func GetUserSessionByID(userID string) (string, error) {
	session, err := RedisClient.HGet(forwardHash, userID).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return session, nil
}

func GetUserIDByUserSession(session string) (string, error) {
	userid, err := RedisClient.HGet(reverseLookupHash, session).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return userid, nil
}

func DeleteUserBySession(session string) (bool, error) {
	userid, err := GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("DeleteUserBySession : %e", err)
		return false, err
	}
	err = RedisClient.HDel(reverseLookupHash, session).Err()
	if err != nil {
		tolog.Log().Errorf("DeleteUserBySession : %e", err)
		return false, err
	}
	err = RedisClient.HDel(forwardHash, userid).Err()
	if err != nil {
		tolog.Log().Errorf("DeleteUserBySession : %e", err)
		return false, err
	}
	err = RedisClient.Del(session).Err()
	if err != nil {
		tolog.Log().Errorf("DeleteUserBySession : %e", err)
		return false, err
	}
	return true, nil
}
