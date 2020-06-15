package main

import (
	"fmt"
	_ "io"
	"io/ioutil"
	"net/http"
	_ "net/http"
	"os"
	_ "os"
	"regexp"
	_ "regexp"
	"strings"
)

var file *os.File

func main() {

	var reg = regexp.MustCompile(`[[:space:]]`)
	var name = regexp.MustCompile(`[[:punct:]]|[["https"]|[["www"]`)

	fmt.Print("Введите расположение файлов к url адресам или 1  = ")
	var urlPath, resultPath string
	fmt.Fscan(os.Stdin, &urlPath)
	if urlPath == "1" {
		urlPath = "C:/Users/Admin/Desktop/url.txt"
	}
	urlStr, err := ioutil.ReadFile(urlPath)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Print("Введите расположение папки результатов или 1  = ")
	fmt.Fscan(os.Stdin, &resultPath)
	if resultPath == "1" {
		resultPath = "C:/Users/Admin/Desktop/result/"
	}
	str := string(urlStr)

	for i := 0; i < len(string(str)); i++ {
		urllines := strings.Split(string(str), "\n")
		println("line №", i, "1", urllines[i])

		urllines[i] = reg.ReplaceAllString(urllines[i], "")
		nameFile := name.ReplaceAllString(urllines[i], "")

		fmt.Println("name file = ", nameFile)

		file, err := os.Create(resultPath + nameFile + ".html")
		if err != nil {
			fmt.Println("can't create file", err)
			os.Exit(1)
		}
		resp, err := http.Get(urllines[i])
		if err != nil {
			fmt.Println("урл не работает")
			fmt.Println(err)
			return
		}
		getHtml(resp, file)
		defer file.Close()

	}
	fmt.Println("urllines = ", urllines[1])

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
