package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func main() {
	//Golang发送post请求
	post := "{}"
	//api_url := ""
	api_url := "http://ubmp-gateway.ubmp.svc.cluster.local:8078/portal/v1/api/checkAccessToken"
	var jsonstr = []byte(post) //转换二进制
	now := time.Now().Second()
	buffer := bytes.NewBuffer(jsonstr)
	request, err := http.NewRequest("POST", api_url, buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")                                                                                                                                                                                                          //添加请求头
	request.Header.Set("Authorization", "Bearer ZH_00001:eyJ0eXBlIjoiSldUIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiJ0ZXN0IiwidGVuZW1lbnRDb2RlIjoiWkhfMDAwMDEiLCJleHAiOjE2NjkzNDY4NjAsImlhdCI6MTY2ODA1MDg2MCwiYWNjb3VudCI6IjE1MjM2MjAzOTEwIn0.d8S38_TVRNCATXj9lnFhRIX2EAIAzEzr3I3GWdQB8nM") //添加请求头
	//request.Header.Set("Authorization", "Bearer ZH_00001:c60f6fc1d4fb4a7091efdb9399682366") //添加请求头
	client := http.Client{}                                     //创建客户端
	resp, err := client.Do(request.WithContext(context.TODO())) //发送请求
	if err != nil {
		fmt.Printf("client.Do%v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
	}
	end := time.Now().Second()
	fmt.Println("response:" + strconv.FormatInt(int64(end-now), 10) + string(respBytes))

}
