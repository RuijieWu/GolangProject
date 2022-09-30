/*
 * @Author: JeRyWu 1365840492@qq.com
 * @Date: 2022-09-28 22:00:42
 * @LastEditors: JeRyWu 1365840492@qq.com
 * @LastEditTime: 2022-09-30 21:57:11
 * @FilePath: /GoDownloader/GoDownloader.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main
 import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
//	"net/url"
	"os"
//	"path/filepath"
//	 "runtime"
//	 "sync"
//	 "flag"
//	 "io"
	"log"
//	 "math"
//	 "strconv"
	"time"
 )
 //读Golang代码时学到的编程风格
 var (
	 core       int       // 协程数目
	 url        string    // 下载链接
	 filepath   string    // 保存路径
	 length     int       // 文件长度
	 size       int       // 片段大小
	 start      time.Time // 起始时间
	 end        time.Time // 完成时间
	 data       string    // 数据实体
 //	 wait sync.WaitGroup 没用上的同步线程的变量
 )
 
 //学到了利用flag包解析参数提供交互，但下载器未做完还没放进去
 //学习到这点的站点https://www.cnblogs.com/flash55/p/11160858.html
 /*func init() {
	 flag.IntVar(&core, "n", 5, "The number of goroutines")
	 flag.StringVar(&url, "u", "", "Target URL")
	 flag.StringVar(&filepath, "p", "", "The path where the file is saved")
	 flag.Parse()
	 if filepath == "" {
		 filepath= "C:\Users\"
	 }
 }*/
 //	通过NewRequest自定义请求头
 // 	学到了把error写进系统的报错日志中
 //  学到了带时间戳输入输出
 /*
	 func MutiDownload(num, start, end int) {
	 StartTime := time.Now()
	 req, err := http.NewRequest("GET", url, nil)
	 if err != nil {
		 log.Fatal(err)
	 }
	 req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	 log.Printf("[%d] start download(range bytes: %d-%d)", bum, start, end)
	 resp, err := http.DefaultClient.Do(req)
	 if err != nil {
		 log.Fatalln(err)
	 }
	 temp := fmt.Sprintf("%s.%d", path, num)
	 file, err := os.Create(temp)
	 if err != nil {
		 file.Close()
		 resp.Body.Close()
		 log.Fatalf("[%d] cannot create %s %v\n", num, temp, err)
	 }
	 n, err := io.Copy(file, resp.Body)
	 if err != nil {
		 file.Close()
		 resp.Body.Close()
		 log.Fatalf("[%d] cannot write %s %v\n", num, temp, err)
	 }
	 EndTime := time.Now()                        // 结束时间
	 duration := EndTime.Sub(StartTime).Seconds() // 用时
	 speed := float64(n) / duration / 1024 / 1024 // 计算平均下载速度
	 log.Printf("[%d] download successfully: %f MB/s\n", num, speed)
	 file.Close()
	 resp.Body.Close()
 }*/
 //获取目标文件大小
 /*	func getlen(url string) int64 {
		 response, err := http.Get(url)
		 if err != nil {
			 log.Println(err)
			 return 0
		 }
		 filesize := response.Header.Get("Content-Length")
		 return filesize
	 } */
 //直接下载目标文件
 func directdownload(url string) string {
	 response, err := http.Get(url)
	 if err != nil {
		 log.Println(err)
		 return "error"
	 }
	 fmt.Println("文件大小为", response.Header.Get("Content-Length"), " Bytes") //找了好久才找到访问响应头这个字段的方法，裂开了
	 data, err := ioutil.ReadAll(response.Body)
	 if err != nil {
		 log.Println(err)
		 return "error"
	 }
	 return string(data)
 }
 //实现文件的保存
 func render(filepath, data string) {
	 file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	 if err != nil {
		 fmt.Println("文件打开失败", err)
		 log.Println(err)
	 }
	 defer file.Close()
	 write := bufio.NewWriter(file)
	 write.WriteString(data)
	 write.Flush()
 }
 
 /*
 找到的实现组合断点续传文件的函数，但还没研究清楚
 func MergeFragments() {
	 dest, err := os.Create(path)
	 if err != nil {
		 dest.Close()
		 log.Fatalf("Create file %s %v\n", path, err)
	 }
	 defer dest.Close()
	 fmt.Println("Merging files...")
	 for i := 0; i < amount; i++ {
		 temp := fmt.Sprintf("%s.%d", path, i)
		 fmt.Printf("reading file %s\n", temp)
		 source, err := os.Open(temp)
		 if err != nil {
			 log.Fatalln("reading file error:", err)
		 }
		 _, err = io.Copy(dest, source)
		 if err != nil {
			 log.Fatalln(err)
		 }
		 source.Close()
		 os.Remove(temp)
	 }
	 fmt.Println("Merge file complete")
 }*/
 func main() {
	 fmt.Println("请输入下载目标url:")
	 fmt.Scanln(&url)
	 fmt.Println("请输入文件保存的绝对路径（例如/home/1.jpg):")
	 fmt.Scanln(&filepath)
	 // 获取文件长度
	 //length, _ = strconv.Atoi(resp.Header.Get("Content-Length"))
	 //size = int(math.Ceil(float64(length) / float64(amount))) // 文件片段长度*/
	 /*	没用上的根据目标url是否支持断点续传的流程控制
	 if checkrq(url) {
			 filesize := getlen(url)
			 cores := runtime.NumCPU()
			 for icount := 1; icount <= cores; icount++ {
				 start := (filesize / cores) * (icount - 1)
				 end := (filesize / cores) * icount
				 wait.Add(1)
				 go multidownload(start, end, icount)
			 }
			 wait.Wait()
		 } else {
			 fmt.Println("It cannot be downloaded by breakpoint resume method")
	 */
	 data = directdownload(url)
	 render(filepath, data)
	 //}
 }