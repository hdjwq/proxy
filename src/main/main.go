package main

import (
	"os"
	"fmt"
	"path"
	"io/ioutil"
	"server"
)

func main()  {
	 file,err:=os.Open(path.Join(os.Args[0],"../config.json"))
	if err!=nil {
		fmt.Println("打开文件失败:%v",err);
		return
	}
	str,err:=ioutil.ReadAll(file)
	if err!=nil {
		fmt.Println("读取文件失败:%v",err);
		return
	}
	server.Start(str)
}
