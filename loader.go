package cascade

import "fmt"

type Loader struct {
	file      string
	envPrefix string
	withFlags bool
}

type Option func(*Loader)

func WithFile(path string) Option {
	return func(l *Loader) {
		l.file = path
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
	if l.file != "" {
		if err := loadFile(l.file, cfg); err != nil {
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
