package request

import (
	"encoding/json"
)

type AIRequest interface {
	MarshalJSON() ([]byte, error)
	Send(url string, apiKey string) (key string, res AIResponse, err error)
}

// OpenAIRequest 定义向 OpenAI API 发起聊天请求的请求体类型
// Model 是要使用的模型的名称，例如 "gpt-3.5-turbo"
// Messages 是一个包含对话历史的消息列表，每个消息都有一个角色和内容
// Temperature 是采样温度，控制输出的随机性，范围从 0 到 2
// TopP 是核采样概率，控制输出的多样性
// MaxTokens 是生成的最大令牌数
// FrequencyPenalty 是频率惩罚，减少重复的可能性
// PresencePenalty 是存在惩罚，鼓励模型引入新话题
type OpenAIRequest struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float32   `json:"temperature,omitempty"`
	TopP             float32   `json:"top_p,omitempty"`
	MaxTokens        int       `json:"max_tokens,omitempty"`
	FrequencyPenalty float32   `json:"frequency_penalty,omitempty"`
	PresencePenalty  float32   `json:"presence_penalty,omitempty"`
}

// Message 定义消息结构体，包含角色和内容
// Role 是消息的角色，如 "system", "user", "assistant"
// Content 是消息的内容
// ReasoningContent 是思考过程中的内容
type Message struct {
	Role             string `json:"role"`
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content,omitempty"`
}

// MarshalJSON 实现 json.Marshal 接口，自定义 JSON 序列化
func (r *OpenAIRequest) MarshalJSON() ([]byte, error) {
	type Alias OpenAIRequest
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	})
}

func (r *OpenAIRequest) Send(url string, apiKey string) (key string, res AIResponse, err error) {
	// 发送请求并获取响应
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + apiKey,
	}
	body, err := r.MarshalJSON()
	if err != nil {
		return "", AIResponse{}, err
	}
	resp, err := HTTPRequest("POST", url, body, headers)
	if err != nil {
		return "", AIResponse{}, err
	}
	// 解析响应
	var response OpenAIResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return "", AIResponse{}, err
	}
	return "OpenAI", AIResponse{OpenAI: response}, nil
}

// OpenAIResponse 定义 OpenAI API 聊天请求的响应体类型
// ID 是响应的唯一标识符
// Object 是响应的对象类型
// Created 是响应的创建时间戳
// Model 是使用的模型的名称
// Choices 是一个包含多个选择的列表，每个选择都有一个消息和完成原因
// Usage 是使用的令牌统计信息
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice 定义响应中每个选择的结构
// Index 是选择的索引
// Message 是选择的消息
// FinishReason 是选择的完成原因
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage 定义响应中使用的令牌统计信息
// PromptTokens 是提示令牌的数量
// CompletionTokens 是完成令牌的数量
// TotalTokens 是总令牌的数量
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type AIResponse struct {
	OpenAI OpenAIResponse
}
