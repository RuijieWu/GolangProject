<!--
 * @Author: JeRyWu 1365840492@qq.com
 * @Date: 2022-09-26 21:37:57
 * @LastEditors: JeRyWu 1365840492@qq.com
 * @LastEditTime: 2022-09-26 21:40:58
 * @FilePath: \GolangProject\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
本项目相关的我已有的基础知识：
    有通过Python的requests模块进行HTTP请求的学习经验
    了解HTTP协议的基础知识
    有到指针位置到结构体之前的C/C++知识基础
我计划的实习学习路线：
    速通Golang的基础知识 → 完成任务书前置任务 → 对完成任务所需的基础知识进行补完 → 用伪代码描述我计划实现的下载器 → 用Golang对构想进行实现
P.S. 
    我想试一试自己现在能力的边界在哪，想了解自己如今的水平，以制定自己学习的路线与方向。实习期的前两天我没有马上开始任务的研究，是因为那两天我参加了联创的熬测，但只完成了必做题的basic部分与一道基于我想研究安全方向而出的附加题，最后因码力不足没有通过，但我有很大的收获。在进入华科前，我在计算机上的学习是完全出于兴趣与爱好的，身边完全没有能交流技术的伙伴，因此心中一直对自己学习的方向有些迷茫，不知道自己短板在哪，也不知道接下来怎么做效率才能尽可能高。通过熬测这样在高压环境下注意力高度集中进行学习的独特体验，我能清楚地感受到自己在计算机领域不少应有的通识没有掌握，以及编程水平还停留在只比初中信竞高一点的水平。我认为这对我的帮助是巨大的。因此，在接下来这段同样是完全没有体验过的实习经历中，无论自己是否能够成功通过，我坚信这次实习肯定也能为我带来全新的体验与成长的可能性。当然我肯定希望自己能够成功通过就是了 ：P
    此外，有些代码我不一定了解其中包含的所有相关知识，仅仅是了解到这段代码能实现我需要的功能就缝过来了，私密马赛。
    
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

2022.9.27 Day 2
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

2022.9.28 Day3
用Golang实现了task1，但似乎有点小问题，狂暴DEBUGing。

草windows什么毛病突然连出一堆问题还找不到解决办法

好了换Ubuntu虚拟机了，这下解决了
