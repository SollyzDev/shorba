package generators

import (
    "time"
    "math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func String() string {
    b := make([]byte, 50)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

func LongTitle() string {
    b := make([]byte, 120)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

func ShortTitle() string {
    b := make([]byte, 26)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

func Bool() bool {
    return true
}
