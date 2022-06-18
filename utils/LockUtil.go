package utils

import (
	"dousheng/database"
	sn "github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sync"
	"time"
)

var node *sn.Node

func Init() (err error) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	tm, err := time.Parse(layout, "Feb 4, 2014 at 6:05pm (PST)")
	if err != nil {
		return
	}
	sn.Epoch = tm.UnixNano() / 1000000
	node, err = sn.NewNode(64)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

//密码加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

//密码验证
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

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
