package mainservice

type ConfigMethod func(config interface{}, paths...string) error
