# service-computing-selpg
## 测试部分
First, generate three files, inputFile.txt, outputfile.txt and errorfile.txt<br>
inputFile.txt file has 10 pages, each page has 10 lines, of course, I specified selpg each line is also 10 pages (default is 72)<br>

Test:<br>
1.<br>
input: $ ./selpg -s 1 -e 1 inputfile.txt<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test1.png)<br>
As Figure, the success of the first page of the 10 lines of text output to the screen.<br>


2.<br>
input: $ ./selpg -s 1 -e 1 < inputfile.txt<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test2.png)<br>
The results of this test are similar to test 1.<br>

3.<br>
input: $ python testout.py | ./selpg -s 1 -e 1 < inputfile.txt<br>output:<br>

![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test3.png)<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test3(1).png)<br>
The meaning of this command is to enter the first page of the testout.py file as input to the screen.<br>
My testout.py file is empty, so there is no output on the screen, but the instructions are executed successfully.<br>


4.<br>
input: $ ./selpg -s 5 -e 10 inputfile.txt >outputfile.txt<br>
output:<br>![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test4(1).png)<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test4(2).png)<br>
This command writes 5 to 10 pages of inputfile.txt to outputfile.txt. As shown in the figure, outputfile.txt has had the<br>
content of inputfile.txt.<br>


5.<br>
input: $ ./selpg -s 1 -e 8 inputfile.txt 2>errorfile.txt<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test5(1).png)<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test5(2).png)<br>
This command is to inputfile.txt in the 1 to 8 pages of content written to screen, where the wrong<br>
information write to errorfile.txt. As shown in the figure, screen has had the corresponding<br>
content, and there is no error message so the errorfile.txt only have the information of executed command.<br>

6.<br>
input: $ ./selpg -s 4 -e 6 inputfile.txt >outputfile.txt 2>errorfile.txt<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test6(1).png)<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test6(2).png)<br>
This command is to inputfile.txt in the 4 to 6 pages of content written to outputfile.txt, where the wrong<br>
information write to errorfile.txt. As shown in the figure, outputfile.txt file has had the corresponding<br>
content, and there is no error message so the errorfile.txt only have the information of executed command.<br>

7.<br>
input: $ ./selpg -s 2 -e 4 inputfile.txt >outputfile.txt 2>/dev/null<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test7.png)<br>
The outputfile.txt does not change compared to the previous one, and the command has been executed successfully,<br>
so there is no error message.<br>

8.<br>
input: $ ./selpg -s 1 -e 3 inputfile.txt >/dev/null<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test8.png)<br>
The result is same as the test 7.<br>

9.<br>
input: $ ./selpg -s 2 -e 5 -l 7 inputfile.txt<br>
output:<br>

![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test11.png)<br>
This command means that 2 to 5 pages of the inputfile.txt file are output to the screen, but the number of lines<br>
<per page is 7. That is, the first inputfile.txt file re-page and then output. Of course, the contents of the<br>
process of this document is the same<vr>
10.<br>
input: $ ./selpg -s 1 -e 4 -f inputfile.txt<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test12.png)<br>
The meaning of this command is based on page breaks to delimit, inputfile.txt 1 to 4 pages output to the screen. <br>

Here is to say that I defined the page break for the "\ f", because the inputfile.txt file does not "\ f" appears, <br>
so all the files are considered in the first page. After the implementation of this command there is an error message.<br>
11.<br>
input: $ ./selpg -s 1 -e 3 inputfile.txt > outputfile.txt 2>errorfile.txt &<br>
output:<br>
![image](https://github.com/Tendernesszh/service-computing-selpg/blob/master/testpicture/test14.png)<br>


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
 

