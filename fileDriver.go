package cascade

type FileDriver interface {
	CanHandle(path string) bool
	Unmarshal(data []byte, cfg any) error
}
