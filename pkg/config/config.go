package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	OSS    OSSConfig    `mapstructure:"oss"`
	MinIO  MinIOConfig  `mapstructure:"minio"`
}
type ServerConfig struct {
	Port       string `mapstructure:"PORT"`
	UploadType string `mapstructure:"UPLOAD_TYPE"`
	RunMode    string `mapstructure:"RUN_MODE"`
}
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}
type MinIOConfig struct {
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
}
type OSSConfig struct {
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	Domain    string `mapstructure:"domain"`
}

var Conf = new(Config)

func Init() {
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("未找到配置文件或读取失败， 尝试使用环境变量，错误信息：%v", err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalf("解析配置文件失败：%v", err)
	}
	log.Println("配置加载成功")
	//热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("检测到配置文件已经修改 ", e.Name)
		//重新解析
		if err := viper.Unmarshal(Conf); err != nil {
			log.Println("配置文件重载失败，使用原配置", err)
		} else {
			log.Println("配置文件重载成功，新配置已经生效")
		}
	})

}
