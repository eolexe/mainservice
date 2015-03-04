package mainservice

type Service interface {
	Init() error
	Run() error
	Stop() error
	NewConfig() interface{}
}

