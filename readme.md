<!--
 * @Author: JeRyWu 1365840492@qq.com
 * @Date: 2022-09-26 21:37:57
 * @LastEditors: JeRyWu 1365840492@qq.com
 * @LastEditTime: 2022-09-26 21:40:58
 * @FilePath: \GolangProject\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
2022.9.26 Day 1
    赶完了这逆天学院的作业DDL，正式开始准备实习项目，先配置好了VScode的Golang环境，运行了一个HelloWorld，并阅读了Golang的菜鸟教程开始尽可能地学习，记录如下：
    Go语言要求{不能单独成行是不是和Python通过代码缩进来区分代码层次一样，是一种强制性地实现编程规范地手段？草{不单独成行的写法太折磨了。
    Golang是在声明变量语句的末端加上变量的数据类型，Go语言的开发者是抱有着怎样的想法决定这样构建Go语言的？相较于数据类型放在前面会有什么样的优势？例如代码可读性上的提升？
    为了提高完成本任务的效率先去了解了最接近的任务断点续传的原理，我理解的原理如下：
    task 1：调用相关借口像目标文件服务器发送http请求（一般是get方法），通过获取响应头的Content-Length字段得知下载文件的大小，获得目标文件的body内容后写入到本地文件中保存到指定路径。
    task 2：断点续传：我们能通过http报文的响应头Accept-Ranges字段了解到下载文件是否支持断点续传，当我们暂停下载后，可以在本地记录下载进度，然后请求待下载文件未下载完成的分段部分，下载完后将所有片段合并为一个文件完成下载。
    task 3 还未来得及了解但我设想的Go语言实现方式是：调用与本机网络有关的接口，获取本机吞吐量的信息，得到下载速度。每隔一段时间读取已写入的数据内容获得长度得到以下载的文件大小并通过计算得到百分比。
    Day1用了三个多小时的时间了解了Go语言的部分语法，并对前3个task的实现有了初步的计划。我打算在Day2学完task涉及的Go语言知识然后先完成task1，并开始学习task4与task5的知识内容。今天的记录到此结束。

    *******************************************************************************************************************

2022.9.27 Dat 2
    Go语言基本知识大概了解了，但还是不大能够很熟悉的运用。今天先把伪代码写了，然后在将伪代码用Golang实现。
    task 1
    input url,path
    response = requests.get("url")
    filesize = response.head[Content-length]
    file = response.body
    render(file,path)
advance
    task 2
    input url,path
    start = 0
    headresponse = requests.gethead("url")  //mark
    if headresponse[Accept-Ranges] == "byte"
        response = requests.get("url"[start:])
	filesize = response.head[Content-length]
        if getchar() == "pause"
	    downloadstop()
	    render(downloadedfile,path)
	    start = len(downloadedfile)
	    system("pause")
	else
	    goto mark

明天又是从早上到晚上的满课，逆天学校，进度没赶上，争取在Day3完成task1，2