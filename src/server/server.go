package server

import (
    "fmt"
	"encoding/json"
    "net/http"
    "strings"
)
var serverConfig *Server;
type Server struct {
	Port string `json:"port"`
	Target string `json:"target"`
	Proxy string `json:"proxy"`
	Static string `json:"static"`
}
type handle struct {

}
func (h *handle)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	  uri:=r.RequestURI;
	if strings.HasPrefix(uri,serverConfig.Proxy) {
		isProxy(w,r,serverConfig);
		return
	}else {
        isStatic(w,r,serverConfig)
	}

}
func Start(str [] byte)  {
	serverConfig=&Server{}
	err:=json.Unmarshal(str,serverConfig);
	if err!=nil {
		fmt.Println("config.json文件错误:%v",err)
	}
	handler:=&handle{};
    httpServer:=&http.Server{
    	 Addr:serverConfig.Port,
    	 Handler:handler,
	}
	fmt.Printf("server run addr:localhost%s\n",serverConfig.Port)
	httpErr:=httpServer.ListenAndServe();
	if httpErr!=nil {
		fmt.Println("服务启动错误:%v",httpErr)
	}
}
