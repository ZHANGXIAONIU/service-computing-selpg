# service-computing-selpg
















































## 设计说明

selpgArgs是一个结构体，用于存储命令行的所有参数，构造如下:<br>
```
type selpgArgs struct {
	start int
	end   int
	pagelength   int
	page_seperator bool
	print_dest string
	inFilename string
}
```
<br>
函数initSelpg用于初始化参数，解析命令，并且判断输入是否合理，比如说起始页数是不是大于1，终止页是不是大于等于起始页,<br>
以及是不是只有一个读入文件被调用。<br>

```
flag.Parse()
```
......<br>
......<br>

```
if saAddr.start < 1 {
		printError("start page need greater than 0!")
	}
	if saAddr.end < saAddr.start {
		printError("end page need grater than or equal to start page!")
	}
	if saAddr.pagelength < 1 {
		printError("page length need greater than 0!")
	}
```
<br>
<br>
而函数runCommand执行解析后的命令。对于管道问题，我使用了os/exec 包来生成子进程并通过管道输入信息<br>
这里给出部分代码：<br>

```
fout := os.Stdout
	var cmd *exec.Cmd
	if args.printDest != "" {
		tmpStr := fmt.Sprintf("%s", args.printDest)
		cmd = exec.Command("sh", "-c", tmpStr)
		if err != nil {
			printError("could not open pipe to \"" + tmpStr + "\"!")
		}
	}
```
<br>
 ......<br>
 ......<br>
 

