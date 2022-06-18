package utils

import (
	"dousheng/database"
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex

func Lock(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	startTime := time.Now().Unix()
	for {
		lock, err := database.RedisDB.SetNX(key, 1, 5*time.Second).Result()
		if err != nil {
			log.Println(err.Error())
		}
		if lock == true {
			return true
		} else {
			time.Sleep(5 * time.Millisecond)
		}
		endTime := time.Now().Unix()
		if endTime-startTime >= 5 {
			break
		}
	}
	return false
}

func UnLock(key string) {
	if _, err := database.RedisDB.Del(key).Result(); err != nil {
		log.Println(err.Error())
	}
	//fmt.Println("unlock", nums)
}
