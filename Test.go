package main

import (
"fmt"
"runtime"
"strings"
"net/http"
"sync"
"io/ioutill"//经查询后会用到的所有包
)

func main()  {
    var url filepath string
	fmt.scanf("&s",&url);
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }//报错则取消下载   
    resp
}
func savefile(filename,filepath,content string){
//保存文件的函数
}