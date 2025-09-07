package cascade

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

func loadFlags(cfg any) error {
	flagSet := flag.NewFlagSet("config", flag.ContinueOnError)

	if err := registerFlags(flagSet, cfg, ""); err != nil {
		return err
	}

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	return nil
}

func registerFlags(flagSet *flag.FlagSet, cfg any, prefix string) error {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			if err := registerFlags(flagSet, field.Addr().Interface(), prefix); err != nil {
				return err
			}
			continue
		}

		flagTag := fieldType.Tag.Get("flag")
		if flagTag == "" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			flagSet.StringVar(field.Addr().Interface().(*string), flagTag, field.String(), "config flag")
		case reflect.Int:
			flagSet.IntVar(field.Addr().Interface().(*int), flagTag, int(field.Int()), "config flag")
		case reflect.Bool:
			flagSet.BoolVar(field.Addr().Interface().(*bool), flagTag, field.Bool(), "config flag")
		}
	}

	return nil
}
