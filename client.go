package httpclient

import (
	"errors"
	"github.com/luckydog8686/logs"
	"golang.org/x/xerrors"
	"net/http"
	"reflect"
)

func NewHttpClient(addr string, outs []interface{}, requestHeader http.Header) error {
	for _,v := range outs{
		vtype := reflect.TypeOf(v)
		if vtype.Kind() != reflect.Ptr{
			return errors.New("expected handler to be a pointer")
		}
		etype := vtype.Elem()
		if etype.Kind()!= reflect.Struct{
			return errors.New("handler should be a struct")
		}
		val := reflect.ValueOf(v)
		for i:=0;i<etype.NumField();i++{

		}
	}
}

func makeFunc(f reflect.StructField)(reflect.Value, error)  {
	ftyp := f.Type
	if ftyp.Kind() != reflect.Func {
		return reflect.Value{}, errors.New("handler field not a func")
	}
	return reflect.Value{},nil
}