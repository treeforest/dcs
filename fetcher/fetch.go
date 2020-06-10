package fetcher

import (
	"net/http"
	"fmt"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"log"
	"time"
	"net/url"
)

// 每隔10ms进行一次访问
var rateLimit = time.Tick(1000*time.Millisecond)

// 代理地址
var proxyLocation = "http://127.0.0.1:45002"

// http 代理访问
func HttpProxyFetch(URL string) ([]byte, error) {
	<- rateLimit

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyLocation)
	}

	transport := &http.Transport{Proxy:proxy}
	client := &http.Client{Transport:transport}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: get  url: %s", URL)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := client.Do(req)
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 模拟浏览器访问
func WebFetch(URL string)([]byte, error) {
	<-rateLimit

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: get  url: %s", URL)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := client.Do(req)

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 原生的方式访问
func BaseFetch(URL string) ([]byte, error) {
	resp, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d\n", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err := r.Peek(1024)

	if err != nil {
		log.Printf("fetch error: %v\n", err)
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(bytes, "")
	return e
}