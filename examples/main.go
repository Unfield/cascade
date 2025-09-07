package main

import (
	"fmt"
	"log"

	"github.com/Unfield/Cascade"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port" toml:"port" env:"PORT" flag:"port"`
		Host string `yaml:"host" toml:"host" env:"HOST" flag:"host"`
	}
	Security struct {
		EnableTLS bool   `yaml:"enable_tls" toml:"enable_tls" env:"ENABLE_TLS" flag:"enable-tls"`
		CertFile  string `yaml:"cert_file" toml:"cert_file" env:"CERT_FILE" flag:"cert-file"`
	}
}

func main() {
	cfg := Config{}
	// defaults
	cfg.Server.Port = 8080
	cfg.Server.Host = "0.0.0.0"

	loader := Cascade.NewLoader(
		Cascade.WithFile("config.yaml"), // or config.toml
		Cascade.WithEnvPrefix("APP"),
		Cascade.WithFlags(),
	)

	if err := loader.Load(&cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server running on %s:%d (TLS: %v)\n",
		cfg.Server.Host, cfg.Server.Port, cfg.Security.EnableTLS)
}
