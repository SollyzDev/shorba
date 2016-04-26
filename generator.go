package shorba

import (
    "time"
)

func GenerateString() string {
    return "This is string"
}

func GenerateInt() int {
    return 12
}

func GenerateBool() bool {
    return true
}


func GenerateTime() time.Time {
    return time.Now()
}