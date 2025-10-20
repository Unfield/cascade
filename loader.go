package cascade

import (
	"fmt"
	"reflect"
)

type Loader struct {
	file      string
	envPrefix string
	withFlags bool
	drivers   []FileDriver
}

type Option func(*Loader)

func WithFile(path string) Option {
	return func(l *Loader) {
		l.file = path
	}
}

func WithCustomFileDriver(driver FileDriver) Option {
	return func(l *Loader) {
		l.drivers = append(l.drivers, driver)
	}
}

func WithEnvPrefix(prefix string) Option {
	return func(l *Loader) {
		l.envPrefix = prefix
	}
}

func WithFlags() Option {
	return func(l *Loader) {
		l.withFlags = true
	}
}

func NewLoader(opts ...Option) *Loader {
	l := &Loader{}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Loader) Load(cfg any) error {
	ensureStructFields(cfg)

	if l.file != "" {
		if err := l.loadFile(cfg); err != nil {
			return fmt.Errorf("load file: %w", err)
		}
	}

	if err := loadEnv(l.envPrefix, cfg); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	if l.withFlags {
		if err := loadFlags(cfg); err != nil {
			return fmt.Errorf("load flags: %w", err)
		}
	}

	return nil
}

func ensureStructFields(cfg any) {
	rv := reflect.ValueOf(cfg)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return
	}
	v := rv.Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct && field.CanSet() && field.IsZero() {
			field.Set(reflect.New(field.Type()).Elem())
			ensureStructFields(field.Addr().Interface())
		}
	}
}
