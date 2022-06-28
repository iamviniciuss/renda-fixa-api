package infra

type HttpService interface {
	Get(path string, callback func(map[string]string, []byte, QueryParams) (interface{}, error))
	Post(path string, callback func(map[string]string, []byte, QueryParams) (interface{}, error))
	ListenAndServe(port string) error
}

type QueryParams interface {
	GetParam(key string) []byte
	AddParam(key string, value string)
}
