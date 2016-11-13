package generators

import (

)

func FirstName() string {
    return ""
}

func LastName() string {
    return ""
}

func FullName() string {
    return FirstName() + " " + MaleName()
}

func MaleName() string {
    return ""
}

func FemaleName() string {
    return ""
}
