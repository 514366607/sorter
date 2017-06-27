package main

import (
	"./bsort"
	"./qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "; outfile=", *outfile, "; algorithm = ", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		var t1 = time.Now()
		switch *algorithm {
		case "qsort":
			qsort.Qsort(values)
		case "bsort":
			bsort.BubbleSort(values)
		default:
			panic(" Sorting algorithm " + *algorithm + " is either unknown or unsupported.")
		}
		var t2 = time.Now()
		fmt.Println(" The sorting process costs ", t2.Sub(t1), " to complete.")

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

//读取文件并返回数组切片
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Filed to open the input file ", infile, err)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF { //文件还没有到结尾，返回错误
				err = err1
				return
			}

			//文件读取到最后一行了
			break
		}

		if isPrefix {
			fmt.Println("A too long line , seems unexpected.")
			return
		}

		var str string = string(line) //以字符串形式接收
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Filed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	var maxkey int = len(values) - 1
	for key, value := range values {
		var str string = strconv.Itoa(value)
		file.WriteString(str)
		if key < maxkey {
			file.WriteString("\n")
		}
	}

	return nil
}
