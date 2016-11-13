package generators

import (
    "math/rand"
)

func Int() int {
    return rand.Intn(100)
}

func Age() int {
    return rand.Intn(60)
}
