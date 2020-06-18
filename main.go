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
	"time"
)

var file *os.File

func main() {
	t0 := time.Now()

	var delSpase = regexp.MustCompile(`[[:space:]]`) //это для того чтобы убрать из url лишние символы
	var delPunct = regexp.MustCompile(`[[:punct:]]`) // и сделать имя файла более приличным

	if len(os.Args) != 3 {
		fmt.Print("не получилось", os.Args[0])
		os.Exit(1)
	}
	urlPath := flag.String("url", "none", "string")
	resultPath := flag.String("resultPath", "none", "string")
	flag.Parse()

	//urlPath := "./url.txt" // проверка в голанд
	//resultPath := "./datafiles/result/"

	os.Mkdir(*resultPath, os.ModePerm)

	urlStr, err := ioutil.ReadFile(*urlPath)
	if (err != nil) || (*urlPath == "none") {
		fmt.Println("неверное расположение файла с url") /////////////////////////////////
		log.Fatal(err)

	}
	ch := make(chan bool)
	i := 0
	str := string(urlStr)
	urllines := strings.Split(string(str), "\n") // разделяем содержимое файла на строки, одна строка одна ссылка

	for index, _ := range urllines {
		urllines[index] = delSpase.ReplaceAllString(urllines[index], "") //убираем пробелы

		nameFile := delPunct.ReplaceAllString(urllines[index], "") //убираем пунктуацию и слэши
		nameFile = strings.Replace(nameFile, "https", "", -1)      // убираем hhtps www
		nameFile = strings.Replace(nameFile, "http", "", -1)
		nameFile = strings.Replace(nameFile, "www", "", -1)
		//fmt.Println("name file = ", nameFile) // просто печатем
		go getHtml(nameFile, *resultPath, urllines, index, ch) //вызов функции resp - ссылка ; file - наш созданный файл go-
		i = index
	}

	for j := 0; j < i; j++ {
		<-ch
	}

	time.Sleep(500 * time.Millisecond)
	t1 := time.Now()
	fmt.Println("Время выполнения = ", t1.Sub(t0))
}

func getHtml(name string, resultP string, urlL []string, i int, ch chan bool) { // эта функция получает ссылку и файл

	fmt.Println(i, name)
	//	time.Sleep(100 * time.Millisecond);
	file, err := os.Create(resultP + name + ".html") // создаем файл resultPath -директория в которую запишется файk
	if (err != nil) || (resultP == "none") {         // и добавляем расширение html
		fmt.Println("can't create file", err)
		os.Exit(1)
	}
	defer file.Close()

	resp, err := http.Get(urlL[i]) // делаем urllines[i] ссылкой
	if err != nil {
		fmt.Println("неверный url адрес")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	for true {

		bs := make([]byte, 1014) // получаем посимвольно содержимое html документа
		n, err := resp.Body.Read(bs)
		file.WriteString(string(bs[:n])) //записываем это все в файл

		if n == 0 || err != nil {
			break
		}
	}
	ch <- true
}
