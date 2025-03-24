package typed

type GlobalConfig struct {
	HostPortConfig HostPortConfig `json:"hostPortConfig"`
	LogConfig      LogConfig      `json:"logConfig"`
}

type HostPortConfig struct {
	IsGlobal bool `json:"isGlobal"`
	Port     int  `json:"port"`
}

type LogConfig struct {
	LogLevel int `json:"logLevel"`
}
