package environment

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func Unmarshal(conf interface{}) error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	v := reflect.ValueOf(conf).Elem()
	for i := 0; i < v.NumField(); i++ {
		tf := v.Type().Field(i)
		key, ok := tf.Tag.Lookup("env")
		if ok {
			value, err := Getenv(key)
			if err != nil {
				return fmt.Errorf("environment does not exist: %v", key)
			}
			vf := v.FieldByName(tf.Name)
			if !vf.CanSet() {
				return fmt.Errorf("cannot set environment to field: %v", tf.Name)
			}
			switch vf.Kind() {
			case reflect.String:
				vf.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val, err := strconv.ParseInt(value, 0, vf.Type().Bits())
				if err != nil {
					return fmt.Errorf("cannot parse environment to int: %v", value)
				}
				vf.SetInt(val)
			case reflect.Bool:
				val, err := strconv.ParseBool(value)
				if err != nil {
					return fmt.Errorf("cannot parse environment to bool: %v", value)
				}
				vf.SetBool(val)
			default:
				return fmt.Errorf("type not supported: %v", v.Kind().String())
			}
		}
	}
	return nil
}

func Getenv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment does not exist: %v", key)
	}

	return val, nil
}
