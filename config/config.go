package config

import (
	"flag"
	"fmt"
	"log"
)

// Settings 全局配置
var Settings = NewConfigDefault()

// 默认配置
const (
	DefaultPort              = 8080
	DefaultProvider          = "ip2location"
	DefaultIp2LocationDbFile = "db/ip2location.bin"
	DefaultIp2RegionDbFile   = "db/ip2region.db"
)

// Config 应用配置
type Config struct {
	BindAddress       string // http server 监听地址
	Provider          string // IP位置转换服务提供商
	Ip2LocationDbFile string // IP2Location数据库文件位置
	Ip2RegionDbFile   string // Ip2Region数据库文件位置
}

func NewConfigDefault() *Config {
	return &Config{
		BindAddress:       fmt.Sprintf(":%d", DefaultPort),
		Provider:          DefaultProvider,
		Ip2LocationDbFile: DefaultIp2LocationDbFile,
		Ip2RegionDbFile:   DefaultIp2RegionDbFile,
	}
}

// Init 从命令行解析参数
func Init() {
	var port int
	var provider string
	flag.IntVar(&port, "p", DefaultPort, "-p <port>")
	flag.StringVar(&provider, "provider", DefaultProvider, "-provider <ip2region | ip2location>")
	flag.Parse()
	Settings.BindAddress = fmt.Sprintf(":%d", port)
	Settings.Provider = provider
	Settings.Ip2LocationDbFile = DefaultIp2LocationDbFile
	Settings.Ip2RegionDbFile = DefaultIp2RegionDbFile

	log.Printf("settings: %+v", Settings)
}
