Fake Code For Whole BreakPoint Resume Downloader
//ʲô�������Զ�մ����ҵ�һ����������α���� 
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
    printf("������%1f \% ,��ʣ��%1f \% ,��ǰ�ٶ�Ϊ%d MBPS ����Ҫ%d ��",downloaded/filesize��undownloaded/filesize,downloadspeed,undownloaded / downloadspeed)
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
    core = runtime.CPUnum //Ĭ���߳���
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