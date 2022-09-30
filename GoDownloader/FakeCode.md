Fake Code For Whole BreakPoint Resume Downloader
//什么语言特性都沾点的乱的一批的下载器伪代码 
```
func directdownload(string url)
{
response = requests.get(url)
save(response.body)
}

func getlen() int
{
head = requests.head()
return head[Content-Length]
}

func multidownload(int start,end,num string url,filepath)
{
response = requests.get(url,Header[Range]="bytes="+start+"-"+end)
content=response.body
save(filepath,char(num)+".tmp",content)
}

func save(string filepath,filename,content)
{
targetfile = filepath+filename
os.create(targetfile)
targetfile.write(content)
}
func show()
{
    downloaded = getsize(file.tmp)
    filesize = requests.head(url)[Content-Length] 
    undownloaded = filesize - downloaded
    downloadspeed = getspeed()
    while(true)
    {
    if downloaded == filesize
    	return
    printf("已下载%1f \% ,还剩下%1f \% ,当前速度为%d MBPS 还需要%d 秒",downloaded/filesize，undownloaded/filesize,downloadspeed,undownloaded / downloadspeed)
    sleep(1)
    }
}

func combine(string filepath,targetfile)
{
for file in os.listdir(filepath)
    flash = read(file)
    targetfile.write(flash)
    file.close()
}

var fin core int

func main()
{
input -> url , filepath
if check(url)
    {
    core = runtime.CPUnum //默认线程数
    for (i=0;i<=core;i++)
	{
	go multidownload(url , filepath , i)
	go show()
	}
    combine(listdir(filepath))
    }
else
    {
    directdownload(url , filepath)
    }
    return 0;
}
```