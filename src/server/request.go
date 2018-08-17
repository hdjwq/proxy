package server

import (
	"net/http"
	"io"
	"errors"
	"fmt"
	"strings"
)

var  client  = &http.Client{}
func request(uri string,method string,body io.Reader,header http.Header) (response *http.Response,err error){
	 request,reqErr:=http.NewRequest(method,uri,body)
	 if reqErr!=nil {
		 err=errors.New(fmt.Sprintf("构建request失败:%v",reqErr))
		 return
	 }
	for key,item:=range header{
		str:=fmt.Sprint(item);
		str=strings.Trim(str,"[]")
		request.Header.Set(key,str);
	}
    response,err=client.Do(request);
	if err!=nil {
		err=errors.New(fmt.Sprintf("获取请求失败:%v",err))
	}
	return

}

