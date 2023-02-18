package main

// @path: stress_testing/stress_testing.go
// @description: stress testing of each service
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/myzhan/boomer"
)

// This is a tool like Apache Benchmark a.k.a "ab".
// It doesn't implement all the features supported by ab.

var client *http.Client
var postBody []byte

var verbose bool

var method string
var url string
var timeout int
var postFile string
var contentType string

var disableCompression bool
var disableKeepalive bool

func worker() {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	request.Header.Set("Content-Type", contentType)

	startTime := time.Now()
	response, err := client.Do(request)
	elapsed := time.Since(startTime)

	if err != nil {
		if verbose {
			log.Printf("%v\n", err)
		}
		boomer.RecordFailure("http", "error", 0.0, err.Error())
	} else {
		boomer.RecordSuccess("http", strconv.Itoa(response.StatusCode),
			elapsed.Nanoseconds()/int64(time.Millisecond), response.ContentLength)

		if verbose {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("%v\n", err)
			} else {
				log.Printf("Status Code: %d\n", response.StatusCode)
				log.Println(string(body))
			}

		} else {
			io.Copy(ioutil.Discard, response.Body)
		}

		response.Body.Close()
	}
}

func main() {
	method = "GET"
	url = "https://www.baidu.com/"
	timeout = 10
	postFile = ""
	contentType = "text/plain"
	disableCompression = false
	disableKeepalive = false
	verbose = false
	/*
			log.Printf(`HTTP benchmark is running with these args:
		method: %s
		url: %s
		timeout: %d
		post-file: %s
		content-type: %s
		disable-compression: %t
		disable-keepalive: %t
		verbose: %t`, method, url, timeout, postFile, contentType, disableCompression, disableKeepalive, verbose)
	*/
	if url == "" {
		log.Fatalln("--url can't be empty string, please specify a URL that you want to test.")
	}

	if method != "GET" && method != "POST" {
		log.Fatalln("HTTP method must be one of GET, POST.")
	}

	if method == "POST" {
		if postFile == "" {
			log.Fatalln("--post-file can't be empty string when method is POST")
		}
		tmp, err := ioutil.ReadFile(postFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		postBody = tmp
	}

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost: 2000,
		DisableCompression:  disableCompression,
		DisableKeepAlives:   disableKeepalive,
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(timeout) * time.Second,
	}

	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn:     worker,
	}
	boomer.Run(task)
}
