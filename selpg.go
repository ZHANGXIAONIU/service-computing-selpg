package main

import (
    "flag"
    "fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
)

type selpgArgs struct {
	start int
	end   int
	pagelength   int
	page_seperator bool
	print_dest string
	inFilename string
}

var progname = "selpg command"


func initSelpg(saAddr *selpgArgs) {
	flag.IntVar(&(saAddr.start), "s", -1, "start page of your file.")
	flag.IntVar(&(saAddr.end), "e", -1, "end page of your file.  must greater than or equal to start page.")
	flag.IntVar(&(saAddr.pagelength), "l", 10, "number of lines in one page. must greater than 0. default value is 72")
	flag.BoolVar(&(saAddr.page_seperator), "f", false, "use [-f=true] to seperate pages by \\f")
	flag.StringVar(&(saAddr.print_dest), "d", "", "specify destionation of output. default destination is stdout")

	flag.Parse()

	if len(flag.Args()) > 1 {
		printError("cannot read more than one file immediately!")
	}
	if len(flag.Args()) == 0 {
		saAddr.inFilename = ""
	} else {
		saAddr.inFilename = flag.Args()[0]
	}

	if saAddr.start < 1 {
		printError("start page need greater than 0!")
	}
	if saAddr.end < saAddr.start {
		printError("end page need grater than or equal to start page!")
	}
	if saAddr.pagelength < 1 {
		printError("page length need greater than 0!")
	}
}

func printError(err string) {
	fmt.Fprintf(os.Stderr, err+"\n"+
		"\nUSAGE: %s -s start_page -e end_page [ -f=true|false | -l lines_per_page ] [ -d dest ] [ in_filename ]\n", progname)
	os.Exit(1)
}

func runCommand() {
	var args selpgArgs
	initSelpg(&args)

	fin := os.Stdin
	var err error
	if args.inFilename != "" {
		fin, err = os.Open(args.inFilename)
		if err != nil {
			printError("could not open input file \"" + args.inFilename + "\"!")
		}
	}

	fout := os.Stdout
	var cmd *exec.Cmd
	if args.print_dest != "" {
		tmpStr := fmt.Sprintf("%s", args.print_dest)
		cmd = exec.Command("sh", "-c", tmpStr)
		if err != nil {
			printError("could not open pipe to \"" + tmpStr + "\"!")
		}
	}

	var line string
	pageCnt := 1
	inputReader := bufio.NewReader(fin)
	rst := ""
	if args.page_seperator == false {
		lineCnt := 0

		for true {
			line, err = inputReader.ReadString('\n')
			if err != nil {
				break
			}
			lineCnt++
			if lineCnt > args.pagelength {
				pageCnt++
				lineCnt = 1
			}
			if pageCnt >= args.start && pageCnt <= args.end {
				if args.print_dest == "" {
					fmt.Fprintf(fout, "%s", line)
				} else {
					rst += line
				}
			}
		}
	} else {
		for true {
			c, _, erro := inputReader.ReadRune()
			if erro != nil {
				break
			}
			if c == '\f' {
				pageCnt++
			}
			if pageCnt >= args.start && pageCnt <= args.end {
				if args.print_dest == "" {
					fmt.Fprintf(fout, "%c", c)
				} else {
					rst += string(c)
				}
			}
		}
	}

	if args.print_dest != "" {
		cmd.Stdin = strings.NewReader(rst)
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			printError("print error!")
		}
	}

	if pageCnt < args.start {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", progname, args.start, pageCnt)
	} else {
		if pageCnt < args.end {
			fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, args.end, pageCnt)
		}
	}

	fin.Close()
	fout.Close()
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
}


func main() {
	runCommand()
}
