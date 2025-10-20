package cascade

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func loadEnv(prefix string, cfg any) error {
	rv := reflect.ValueOf(cfg)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("cfg must be a pointer to a struct")
	}

	v := rv.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		envTag := fieldType.Tag.Get("env")

		if field.Kind() == reflect.Struct {
			nestedPrefix := prefix
			if envTag != "" {
				nestedPrefix = prefix + "_" + strings.ToUpper(envTag)
			}
			if err := loadEnv(nestedPrefix, field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		if envTag == "" {
			continue
		}

		envName := strings.ToUpper(prefix + "_" + envTag)
		if val, ok := os.LookupEnv(envName); ok {
			switch field.Kind() {
			case reflect.String:
				field.SetString(val)
			case reflect.Int:
				var i int
				if _, err := fmt.Sscanf(val, "%d", &i); err != nil {
					return fmt.Errorf("invalid int for %s: %w", envName, err)
				}
				field.SetInt(int64(i))
			case reflect.Bool:
				field.SetBool(val == "true" || val == "1")
			}
		}
	}
	return nil
}
