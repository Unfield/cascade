package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func loadEnv(prefix string, cfg any) error {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			loadEnv(prefix, field.Addr().Interface())
			continue
		}

		envTag := fieldType.Tag.Get("env")
		if envTag == "" {
			continue
		}

		envName := prefix + "_" + envTag
		envName = strings.ToUpper(envName)

		if val, ok := os.LookupEnv(envName); ok {
			if field.Kind() == reflect.String {
				field.SetString(val)
			}
			if field.Kind() == reflect.Int {
				var i int
				fmt.Sscanf(val, "%d", &i)
				field.SetInt(int64(i))
			}
			if field.Kind() == reflect.Bool {
				if val == "true" || val == "1" {
					field.SetBool(true)
				} else {
					field.SetBool(false)
				}
			}
		}
	}
	return nil
}
