package utils

import (
	"math/rand"
	"time"
)

// ランダムなn字の文字列と現在時間をくっつけた文字列を返却する
func GetUniqueString(n int) string {
	strTime := time.Now().Format("20060102150405")
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b) + strTime
}
