package server

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

func isProxy(w http.ResponseWriter,r *http.Request,config *Server)  {
	proxyUrl:=strings.TrimPrefix(r.RequestURI,config.Proxy);
	response,err:=request(config.Target+proxyUrl,r.Method,r.Body,r.Header);
	if err!=nil {
		fmt.Println(err.Error());
		w.WriteHeader(response.StatusCode);
		return
	}
	str,readErr:=ioutil.ReadAll(response.Body);
	if readErr!=nil {
		w.WriteHeader(500);
		return
	}
	_,wErr:=w.Write(str);
	if wErr!=nil {
		w.WriteHeader(500);
		return
	}
}
