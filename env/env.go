package env

type Provider interface {
	Init() error
	Get(key string) interface{}
}
