package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/urfave/cli"
)
// {"_id":"01","title":"如果小静想在朋友圈写日记，她最多可以写多长？","typecode":"01","typename":"单选","comments":"注释信息",
// "options":[{"value":"1","code":"A","content":"2000字"},{"value":"0","code":"B","content":"100字"},
// {"value":"0","code":"C","content":"500字"},{"value":"0","code":"D","content":"1000字"}],"examid":"001001"}

type Option struct{
	Value string `json:"value"`
	Code string `json:"code"`
	Content string `json:"content"`
}

type Questions struct {
	ID string  `json:"_id"`
	Examid string `json:"examid"`
	Typecode string `json:"typecode"`
	Typename string `json:"typename"`
	Comments string `json:"comments"`
	Title  string  `json:"title"`
	Options []Option `json:"options"`
}

func readq() {
	filepath := "q.md"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	var num = 0
	var qnum = 1
	var qbuf = make([]Questions, 0)
	var aquest Questions
	aquest.Examid = "001001"
	aquest.Typecode = "01"
	aquest.Typename = "单选"
	aquest.Comments = "注释信息"

	var op Option

	for {
		num++
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		line = strings.TrimSpace(line)
		// abuf[num] = line
		fmt.Printf("num: %d line: %s\n", num, line)
		switch (num) {
		case 1 :
			aquest.Title = line
		case 2, 3, 4, 5:
			op.Value = "0"
			idList := strings.Split(line , ".")
			op.Code = idList[0]
			op.Content = idList[1]
			aquest.Options = append(aquest.Options, op)
		case 6 :
			// fmt.Println(aquest.Options[0].Content)
			anlist := strings.Split(line , "：")
			// fmt.Println(anlist[1])
			switch (anlist[1]) {
			case "A":
				aquest.Options[0].Value = "1"
			case "B":
				aquest.Options[1].Value = "1"
			case "C":
				aquest.Options[2].Value = "1"
			case "D":
				aquest.Options[3].Value = "1"
			default:
				fmt.Println("no match")
			}
			break
		}

		if (num == 7) {
			aquest.ID = strconv.Itoa(qnum)
			qnum++
			num = 0
			qbuf = append(qbuf, aquest)
			aquest.Options = nil
		}
	}

	file2, err := os.OpenFile("q.json", os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("open file failed,err:",err)
        return
    }

    defer file2.Close()

	for _, val := range qbuf {
		// fmt.Println(val)
		bytes,_ := json.Marshal(val)
		stringData := string(bytes)+"\r\n"
		// fmt.Println(stringData)
		// file2.Write([]byte(stringData))    //写入字节切片数据
		file2.WriteString(stringData) //直接写入字符串数据
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		readq()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
