package home

import (
	"flag"
	"github.com/gookit/color"
)

//定义各个参数
var (
	help     bool
	url      string
	thread   int
	wordlist string
	status   string
	output   string
)

//定义参数调用方法

func Help() bool       { return help }
func Url() string      { return url }
func Thread() int      { return thread }
func Wordlist() string { return wordlist }
func Status() string   { return status }
func Output() string   { return output }

func init() {
	flag.BoolVar(&help, "h", false, "Help")
	flag.StringVar(&url, "u", "", "URL for Scan")
	flag.IntVar(&thread, "t", 30, "Thread")
	flag.StringVar(&wordlist, "wd", "", "wordlist")
	flag.StringVar(&status, "s", "*", "Change StatusCode")
	flag.StringVar(&output, "0", "", "put log to file")

	flag.Parse()
}
func Usage() {
	color.Blue.Println(`
				____  __                          
			   / __ \/ /_  __  _______            
			  / /_/ / __ \/ / / / ___/            
			 / ____/ /_/ / /_/ (__  )             Version: 0.1
			/_/   /_.___/\__,_/____/			  Blog: https://ctfking.com
`)
	color.Info.Println(`
Usage: Pbus <-u URL> <-wd Wordlist> <-t Threads> <-s StatusCode> <-o Output log>

Params:`)
	flag.PrintDefaults()
}
