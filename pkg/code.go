package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func SID(n int) string {
	rand.Seed((time.Now().UnixNano() + int64(time.Now().Second()-time.Now().Hour()+time.Now().Day())) / int64(time.Now().Year()))
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func VerifyCode(min int, max int) int32 {
	rand.Seed(time.Now().UnixNano())

	return int32(min + rand.Intn(max-min))
}
func ChangePassword() string {
	code := VerifyCode(324569, 985234)
	return fmt.Sprintf("%v%v%v", SID(25), code, SID(15))
}
