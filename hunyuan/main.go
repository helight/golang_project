package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

type ResponseBody struct {
	Result string `json:"result"`
}

func main() {
	url := "https://api.hunyuan.tencent.com/v1/respond"
	apiKey := "your_api_key" // 替换为你的 API 密钥

	// 创建请求体
	requestBody := RequestBody{
		Prompt: "你好，世界！",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 解析响应体
	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	fmt.Println("Response:", responseBody.Result)
}
