package util

import (
	"encoding/json"
	"math/rand"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// StringToStruct 字符串转结构体
func StringToStruct(str string, class interface{}) (interface{}, error) {
	b := []byte(str)
	cls := &class
	err := json.Unmarshal(b, cls)
	if err == nil {
		return cls, nil
	} else {
		return nil, err
	}
}

// StructToString 结构体转字符串
func StructToString(s *interface{}) ([]byte, error)	{
	byteArray, err := json.Marshal(&s)

	if err == nil {
		return byteArray, nil
	} else {
		return nil, err
	}
}
