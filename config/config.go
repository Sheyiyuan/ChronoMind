package config

type GlobalConfig struct {
	HostPortConfig HostPortConfig `json:"hostPortConfig"`
	LogConfig      LogConfig      `json:"logConfig"`
	AiApiConfig    AiApiConfig    `json:"aiApiConfig"`
}

type HostPortConfig struct {
	IsGlobal   bool     `json:"isGlobal"`
	Port       int      `json:"port"`
	Mode       int      `json:"mode"`
	Whitelists []string `json:"whitelists"`
	Blacklists []string `json:"blacklists"`
	Admins     []string `json:"admins"`
}

type LogConfig struct {
	LogLevel int `json:"logLevel"`
}

type AiApiConfig struct {
	AiWorkConfig    AiWorkConfig    `json:"aiWorkConfig"`
	AiSpeakerConfig AiSpeakerConfig `json:"aiSpeakerConfig"`
}

type AiWorkConfig struct {
	ApiKey    string `json:"apiKey"`
	ApiUrl    string `json:"apiUrl"`
	ModelName string `json:"modelName"`
	MaxTokens int    `json:"maxTokens"`
}

type AiSpeakerConfig struct {
	ApiKey           string  `json:"apiKey"`
	ApiUrl           string  `json:"apiUrl"`
	ModelName        string  `json:"modelName"`
	MaxTokens        int     `json:"maxTokens"`
	Temperature      float32 `json:"temperature"`
	FrequencyPenalty float32 `json:"frequencyPenalty"`
	PresencePenalty  float32 `json:"presencePenalty"`
	InitPrompt       string  `json:"initPrompt"`
}
