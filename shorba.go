package shorba

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"

	G "github.com/objectizer/shorba/generators"
)

var db *mgo.Database

// Connect connects to mongodb and initials db instance
func Connect(host string, username string, passwor string, dbName string) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	db = session.DB(dbName)
}

// getMap converts a struct to map that contains field_name and field_type
func getMap(model interface{}) (map[string]string, error) {
	modelReflect := reflect.ValueOf(model)
	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}
	// accept only a struct
	if modelReflect.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Only accepts structs")
	}
	modelReflectType := modelReflect.Type()
	values := make(map[string]string, modelReflect.NumField())
	for i := 0; i < modelReflect.NumField(); i++ {
		field := modelReflectType.Field(i)
		if tag := field.Tag.Get("bson"); tag != "" {
			t := strings.Split(tag, ",")
			values[t[0]] = modelReflect.Field(i).Type().String()
			//values[t[0]] = []byte(modelReflect.Field(i).Type())
		}
	}
	return values, nil
}

// Populate generates a dummy data for a certain collection
// based on the struct defention, it creates (n) of documents
func Populate(collection string, model interface{}, n int) error {
	// convert interface to a map
	modelMap, err := getMap(model)
	if err != nil {
		return err
	}
	// loop on n and generate random values for each field
	for x := 0; x < n; x++ {
		item := make(map[string]interface{})
		for key, val := range modelMap {
			switch val {
			case "string":
				item[key] = G.String()
			case "int":
				item[key] = G.Int()
			case "bool":
				item[key] = G.Bool()
			case "bson.ObjectId":
				item[key] = bson.NewObjectId()
			case "time.Time":
				item[key] = G.Time()
			default:
				//item[key] = GenerateString()
			}
			fmt.Printf("%s: %s => %v\n", key, val, item[key])
		}
	}
	return nil
}
