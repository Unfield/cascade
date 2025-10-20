# 🌊 Cascade

[![Go Reference](https://pkg.go.dev/badge/github.com/Unfield/cascade.svg)](https://pkg.go.dev/github.com/Unfield/cascade)
[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.24.4-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/Unfield/cascade)](https://goreportcard.com/report/github.com/Unfield/cascade)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/Unfield/cascade?label=version&color=green)

**Cascade** is a modular configuration loader for Go.
It merges configuration from multiple sources in a cascading order:

**File < Environment < Command-line flags**

---

## Features

- 📂 Load from **YAML** or **TOML** files
- 🌍 Override with **environment variables**
- 🖥️ Override with **command-line flags**
- ⚡ Define **defaults** in your Go structs
- 🔌 Modular & reusable across projects

---

## Quick Example

```go
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
    cfg.Server.Port = 8080 // default
    cfg.Server.Host = "0.0.0.0"

    loader := cascade.NewLoader(
        cascade.WithFile("config.yaml"), // or config.toml
        cascade.WithEnvPrefix("APP"),
        cascade.WithFlags(),
    )

    if err := loader.Load(&cfg); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Server running on %s:%d (TLS: %v)\n",
        cfg.Server.Host, cfg.Server.Port, cfg.Security.EnableTLS)
}
```

---

## Priority Order

1. **Defaults** (in Go struct)
2. **Config file** (YAML/TOML)
3. **Environment variables** (`APP_PORT=9000`)
4. **Command-line flags** (`--port=7000`)

---

## License

MIT
