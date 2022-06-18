package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

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

// 返回一个32位md5加密后的字符串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
