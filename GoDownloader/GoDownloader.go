/*
 * @Author: JeRyWu 1365840492@qq.com
 * @Date: 2022-09-28 22:00:42
 * @LastEditors: JeRyWu 1365840492@qq.com
 * @LastEditTime: 2022-09-28 23:43:22
 * @FilePath: /GoDownloader/GoDownloader.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"bufio"
	"os"
)

func download(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Errow", err)
		return "error"
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Errow", err)
		return "error"
	}
	fmt.Println("文件大小为", len(data))
	return string(data)
}
func render(filepath, data string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(data)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
func main() {
	var url, filepath, data string
	fmt.Scanln(&url)
	fmt.Scanln(&filepath)
	data = download(url)
	render(filepath, data)
}
