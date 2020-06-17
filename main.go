package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var file *os.File

func main() {

	var delSpase = regexp.MustCompile(`[[:space:]]`) //это для того чтобы убрать из url лишние символы
	var delPunct = regexp.MustCompile(`[[:punct:]]`) // и сделать имя файла более приличным

	if len(os.Args) != 3 {
		fmt.Print("не получилось", os.Args[0])
		os.Exit(1)
	}
	urlPath := flag.String("url", "none", "string")
	resultPath := flag.String("resultPath", "none", "string")
	flag.Parse()

	os.Mkdir(*resultPath, os.ModePerm)

	//if urlPath == "1" {		//это мне надо было для быстрой проверки программы
	//	urlPath = "C:/Users/Admin/Desktop/url.txt"
	//}

	urlStr, err := ioutil.ReadFile(*urlPath)
	if (err != nil) || (*urlPath == "none") {
		fmt.Println("неверное расположение файла с url") /////////////////////////////////
		log.Fatal(err)
		//res:="ошибка чтения файла";

	}

	//if resultPath == "1" {
	//	resultPath = "C:/Users/Admin/Desktop/result/"
	//}

	str := string(urlStr)

	urllines := strings.Split(string(str), "\n") // разделяем содержимое файла на строки, одна строка одна ссылка

	for index, _ := range urllines {
		urllines[index] = delSpase.ReplaceAllString(urllines[index], "") //убираем пробелы
		nameFile := delPunct.ReplaceAllString(urllines[index], "")       //убираем пунктуацию и слэши
		nameFile = strings.Replace(nameFile, "https", "", -1)            // убираем hhtps www
		nameFile = strings.Replace(nameFile, "www", "", -1)
		fmt.Println("name file = ", nameFile) // просто печатем

		resp, err := http.Get(urllines[index]) // делаем urllines[i] ссылкой
		if err != nil {
			fmt.Println("неверный url адрес")
			fmt.Println(err)
			return
		}
		getHtml(resp, file) //вызов функции resp - ссылка ; file - наш созданный файл

	}

}

func getHtml(res *http.Response, file2 *os.File) { // эта функция получает ссылку и файл
	defer res.Body.Close()

	for true {

		bs := make([]byte, 1014) // получаем посимвольно содержимое html документа
		n, err := res.Body.Read(bs)
		file2.WriteString(string(bs[:n])) //записываем это все в файл

		if n == 0 || err != nil {
			break
		}
	}
}
