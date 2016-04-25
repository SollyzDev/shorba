package shorba

import (
    "gopkg.in/mgo.v2"
    "fmt"
    "reflect"
    "strings"
)

var db *mgo.Database

func Connect(host string, username string, passwor string, dbName string) {
    session, err := mgo.Dial(host)
    if err != nil {
        panic(err)
    }
    session.SetMode(mgo.Monotonic, true)
    db = session.DB(dbName)
}

// getMap converts a struct to map that contains field_name and field_type
func getMap(model interface{}) (map[string]interface{}, error) {
    modelReflect := reflect.ValueOf(model)
    if modelReflect.Kind() == reflect.Ptr {
        modelReflect = modelReflect.Elem()
    }
    // accept only a struct
    if modelReflect.Kind() != reflect.Struct {
        return nil, fmt.Errorf("Only accepts structs")
    } 
    modelReflectType := modelReflect.Type()
    values := make(map[string]interface{}, modelReflect.NumField())
    for i := 0; i < modelReflect.NumField(); i++ {
        field := modelReflectType.Field(i)
        if tag := field.Tag.Get("bson"); tag != "" {
            t := strings.Split(tag, ",")
            values[t[0]] = modelReflect.Field(i).Type()
        }
    }
    return values, nil
}

// Populate generates a dummy data for a certain collection
// based on the struct defention, it creates (n) of documents
func Populate(collection string, model interface{}, n int) error {
    // convert interface to a map
    _, err := getMap(model)
    if err != nil {
        return err
    }
    return nil
}