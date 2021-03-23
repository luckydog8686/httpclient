package httpclient

import "testing"

type SS struct {
	Internal struct{
		Ping func(str string)(string,error)
	}
}

func (s *SS)Ping(str string)(string,error)  {
	return s.Internal.Ping(str)
}

type SSInterface interface {
	Ping(str string)(string,error)
}

func TestNewHttpClient(t *testing.T) {
	var s SS
	NewHttpClient("http://127.0.0.1:80",[]interface{}{&s.Internal},nil)
}