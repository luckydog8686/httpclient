package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/luckydog8686/logs"
	"io/ioutil"
	"net/http"
	"reflect"
)
/*
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
*/
func MakeFunc(f reflect.StructField)(reflect.Value, error)  {
	ftyp := f.Type
	if ftyp.Kind() != reflect.Func {
		return reflect.Value{}, errors.New("handler field not a func")
	}

	return reflect.Value{},nil
}

func Struct2Client(sc interface{})error  {
	scTyp := reflect.TypeOf(sc)
	if scTyp.Kind() != reflect.Ptr {
		return errors.New("expected handler to be a pointer")
	}
	typ := scTyp.Elem()
	if typ.Kind() != reflect.Struct {
		return errors.New("handler should be a struct")
	}

	vm := reflect.ValueOf(sc).Elem()
	for i:=0;i<typ.NumField();i++{
		fn,err:=MakeHttpPost(typ.Field(i))
		if err != nil {
			return err
		}
		vm.Field(i).Set(fn)
	}
	return nil
}



func MakeHttpPost(f reflect.StructField) (reflect.Value,error) {
	ftyp := f.Type
	if ftyp.Kind() != reflect.Func {
		return reflect.Value{}, errors.New("handler field not a func")
	}
	type Result struct {
		Data  interface{} `json:"data"`
		Error string `json:"error"`
	}
	DoPost:=func (args []reflect.Value)[]reflect.Value  {
		var nilError = reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())
		logs.Info("参数数量:",len(args))
		url := "http://localhost/ss/hello"
		var httpBody []byte
		logs.Info(args[0].Interface())
		out1:=reflect.New(ftyp.Out(0))
		httpBody,err := json.Marshal(args[0].Interface())
		logs.Info(string(httpBody))
		if err != nil {
			logs.Error(err)
			return []reflect.Value{out1,reflect.ValueOf(err)}
		}
		req,err := http.NewRequest("POST",url,bytes.NewBuffer(httpBody))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if resp != nil{
			defer resp.Body.Close()
		}
		if err != nil{
			logs.Error(err)
			return []reflect.Value{out1.Elem(),reflect.ValueOf(err)}
		}
		body, err:= ioutil.ReadAll(resp.Body)
		if err != nil{
			logs.Error(err)
			return []reflect.Value{out1.Elem(),reflect.ValueOf(err)}
		}
		//outE := out1.Interface()
		res := &Result{}
		logs.Info(string(body))
		err = json.Unmarshal(body,res)
		if err != nil{
			logs.Error(err)
			return []reflect.Value{out1.Elem(),reflect.ValueOf(err)}
		}
		logs.Info(res)
		logs.Info(res.Data)
		logs.Info(res.Error)
		logs.Info(reflect.TypeOf(res.Error))
		//errMsg := fmt.Sprintf("%v",res.Error)

		if res.Error==""{
			logs.Info("等于nil")
			return []reflect.Value{reflect.ValueOf(res.Data),nilError}
		}
		logs.Info("不等于nil")
		return []reflect.Value{reflect.ValueOf(res.Data),reflect.ValueOf(errors.New(res.Error))}
		/*
			var payload []byte
			if len(args)>0{
				payload,err := json.Marshal(args[0].Interface())，
				if err != nil{
					return
				}
			}
		*/
	}
	return reflect.MakeFunc(ftyp,DoPost),nil
}

type Response struct {
	Error error `json:"error"`
	Data json.RawMessage `json:"data"`
}
//参数包含请求的数据结构
//


func processError(err error)[]reflect.Value  {
	return nil
}