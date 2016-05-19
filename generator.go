package shorba

import (
    "time"
    "math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateString() string {
    b := make([]byte, 100)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

func GenerateInt() int {
    return rand.Intn(100)
}

func GenerateBool() bool {
    return true
}


func GenerateTime() time.Time {
    return time.Now()
}
