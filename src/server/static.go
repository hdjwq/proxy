package server

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)
func isStatic(w http.ResponseWriter,r *http.Request,config *Server)  {
	 uri:=r.RequestURI;
	 file,err:=os.Open(config.Static+uri);
	if err!=nil {
		if err==os.ErrNotExist {
			w.Write([] byte(fmt.Sprintf("未找到此文件!")))
			return
		}
		w.Write([] byte(fmt.Sprintf("打开文件错误!")))
		return
	}
	fileInfo,err:=file.Stat();
	if err!=nil {
		w.Write([] byte(fmt.Sprintf("系统异常!")))
		return
	}
	if fileInfo.IsDir() {
		w.Write([] byte(fmt.Sprintf("暂不支持文件夹预览!")))
		return
	}
	str,err:=ioutil.ReadAll(file);
	if err!=nil {
		w.Write([] byte(fmt.Sprintf("系统异常!")))
		return
	}
	w.Write(str)
}
