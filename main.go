package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//输入url
	fmt.Println("请输入url: ")
	var url string
	_, err := fmt.Scan(&url)
	if err != nil {
		log.Fatalf("输入错误：%s", err)
	}
	// 如果url后面没有 / 自动加上
	if url[len(url)-1] != '/' {
		url += "/"
	}

	// 打开字典
	wordlist, err2 := os.Open("dicc.txt")

	// 捕获打开字典异常
	if err2 != nil {
		log.Fatalf("Error when opening file: %s", err2)
	}
	defer wordlist.Close()

	fileScanner := bufio.NewScanner(wordlist)

	// 每一行字典请求
	for fileScanner.Scan() {
		var rurl string = url + fileScanner.Text()
		resp, err := http.Get(rurl)
		if err != nil {
			log.Fatalln(err)
		}
		statuscode := resp.StatusCode
		fmt.Println(statuscode, "    ", "/"+fileScanner.Text())

	}
	// 捕获读取字典异常
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("读取文件时报错： %s", err)
	}

}
