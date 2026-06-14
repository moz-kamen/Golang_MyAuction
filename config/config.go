package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	EthClient EthClientConfig `mapstructure:"eth_client"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type EthClientConfig struct {
	RpcUrl         string        `mapstructure:"rpc_url"`
	WebsocketUrl   string        `mapstructure:"websocket_url"`
	Timeout        time.Duration `mapstructure:"timeout"`
	AuctionAddress string        `mapstructure:"auction_address"`
}

func init() {
	// 设置配置文件名称（不含扩展名）
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 添加配置文件搜索路径
	viper.AddConfigPath("./config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[CONFIG]读取配置文件失败: %v", err)
	} else {
		log.Printf("[CONFIG]读取配置文件成功: %s", viper.ConfigFileUsed())
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// 返回解析成功的配置对象
	return &config, nil
}
