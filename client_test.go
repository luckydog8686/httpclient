package httpclient

import (
	"github.com/luckydog8686/logs"
	"reflect"
	"testing"
	"time"
)

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
	/*
	var s SS
	NewHttpClient("http://127.0.0.1:80",[]interface{}{&s.Internal},nil)
	 */
}

func TestDoPost(t *testing.T) {

}

func TestMakeFunc(t *testing.T) {
	var ss SS
	Struct2Client(&ss.Internal)
	time.Sleep(time.Second)
	res,err := ss.Ping("haha")
	logs.Info(reflect.TypeOf(err))
	logs.Info(reflect.TypeOf(nil))
	if err != nil{
		t.Fatal(err)
	}
	t.Log(res)
}