package request

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Sheyiyuan/ChronoMind/logos"
	"io"
	"net/http"
)

// HTTPRequest 发起一个HTTP请求
// method 是HTTP请求方法，如 "GET", "POST" 等
// url 是请求的URL
// body 是请求的主体内容，如果没有可以传入 nil
// headers 是请求头，可以传入自定义的请求头信息
func HTTPRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	// 创建一个新的HTTP请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建一个HTTP客户端
	client := &http.Client{}
	// 发起请求
	logos.Info("发起HTTP请求: %s %s", method, url)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logos.Info("HTTP请求返回状态码: %d", resp.StatusCode)
	// 确保在函数结束时关闭响应体
	defer func(Body io.ReadCloser) {
		logos.Trace("关闭HTTP响应体 %s %s", method, url)
		_ = Body.Close()
	}(resp.Body)

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP请求失败: %d, %s", resp.StatusCode, string(respBody)))
	}

	return respBody, nil
}
