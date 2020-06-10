package main

import (
	"fmt"
	_ "io"
	"io/ioutil"
	"net/http"
	_ "net/http"
	"os"
	_ "os"
	"strings"
)

var file *os.File

func main() {
	urlStr, err := ioutil.ReadFile("url.txt")
	if err != nil {
		fmt.Println("err")
	}
	str := string(urlStr)
	fmt.Println(str)
	for i := 0; i < len(string(str)); i++ {

		urllines := strings.Split(string(str), "\n")
		println("line №", i, " = ", urllines[i])
	}

	//urlFile, err := os.Open("url.txt")
	//if err!=nil{
	//	fmt.Println(err)
	//	os.Exit(2);
	//}
	//defer urlFile.Close()
	//stat,err:=urlFile.Stat()
	//if err!=nil {
	//	return
	//}
	//
	//data:=make([]byte,stat.Size())
	//
	//strUrl := string(data);
	//fmt.Println("str url = ",strUrl)

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("first")
	if err != nil {
		fmt.Println("can't create file", err)
		os.Exit(1)
	}

	defer file.Close()

	getHtml(resp, file)

}

func getHtml(res *http.Response, file2 *os.File) {
	defer res.Body.Close()

	for true {

		bs := make([]byte, 1014)
		n, err := res.Body.Read(bs)
		file2.WriteString(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}
