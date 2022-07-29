package utils

import (
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"
	"sync"
)

func HttpBuildQuery(params map[string]string) string {
	var uri url.URL
	q := uri.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	queryStr := q.Encode()
	return queryStr
}

func HttpGetChan(url string, response chan string, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	response <- string(bytes)
	<-limiter
}

func HttpGet(url string) (result string) {
	// result := orderedmap.NewOrderedMap[string, string]()

	wg := &sync.WaitGroup{}
	// 控制并发数为10
	limiter := make(chan bool, 10)
	defer close(limiter)

	// 函数内的局部变量channel, 专门用来接收函数内所有goroutine的结果
	responseChannel := make(chan string, 20)
	// 为读取结果控制器创建新的WaitGroup, 需要保证控制器内的所有值都已经正确处理完毕, 才能结束
	wgResponse := &sync.WaitGroup{}
	// 启动读取结果的控制器
	go func() {
		// wgResponse计数器+1
		wgResponse.Add(1)
		// 读取结果
		for response := range responseChannel {
			// 处理结果
			// result = append(result, response)
			result = response
		}
		// 当 responseChannel被关闭时且channel中所有的值都已经被处理完毕后, 将执行到这一行
		wgResponse.Done()
	}()

	// 计数器+1
	wg.Add(1)
	limiter <- true
	// 这里在启动goroutine时, 将用来收集结果的局部变量channel也传递进去
	go HttpGetChan(url, responseChannel, limiter, wg)

	// 等待所以协程执行完毕
	wg.Wait() // 当计数器为0时, 不再阻塞
	// fmt.Println("所有协程已执行完毕")

	// 关闭接收结果channel
	close(responseChannel)

	// 等待wgResponse的计数器归零
	wgResponse.Wait()

	// 返回聚合后结果
	return result
}

//raw
func CollectHttpGetRaw(urls []string) []string {
	var result []string

	wg := &sync.WaitGroup{}
	// 控制并发数为10
	limiter := make(chan bool, 10)
	defer close(limiter)

	// 函数内的局部变量channel, 专门用来接收函数内所有goroutine的结果
	responseChannel := make(chan string, 20)
	// 为读取结果控制器创建新的WaitGroup, 需要保证控制器内的所有值都已经正确处理完毕, 才能结束
	wgResponse := &sync.WaitGroup{}
	// 启动读取结果的控制器
	go func() {
		// wgResponse计数器+1
		wgResponse.Add(1)
		// 读取结果
		for response := range responseChannel {
			// 处理结果
			result = append(result, response)
		}
		// 当 responseChannel被关闭时且channel中所有的值都已经被处理完毕后, 将执行到这一行
		wgResponse.Done()
	}()

	for _, url := range urls {
		// 计数器+1
		wg.Add(1)
		limiter <- true
		// 这里在启动goroutine时, 将用来收集结果的局部变量channel也传递进去
		go HttpGetRaw(url, responseChannel, limiter, wg)
	}

	// 等待所以协程执行完毕
	wg.Wait() // 当计数器为0时, 不再阻塞
	// fmt.Println("所有协程已执行完毕")

	// 关闭接收结果channel
	close(responseChannel)

	// 等待wgResponse的计数器归零
	wgResponse.Wait()

	// 返回聚合后结果
	return result
}

//raw
func HttpGetRaw(url string, response chan string, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	response <- string(bytes)
	<-limiter
}
