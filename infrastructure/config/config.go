package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Init read the base file
func init() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("init config error ", err)
		panic(err)
	}
}

// Config 导出的配置类
type Config struct {
	HttpPort        string
	Mode         	string
	DBMode       	string
	MysqlUrl 		string
	MysqlPort       int
	MysqlUserName   string
	MysqlPassWord   string
	DBName       	string
	LogFile      	string
	LogLevel     	string
	LogCount 		uint
}

// GetConfig get config
func GetConfig() *Config {
	return &Config{
		HttpPort:     	viper.GetString("httpport"),
		Mode:         	viper.GetString("mode"),
		DBMode:  	  	viper.GetString("dbmode"),
		MysqlUrl:  	  	viper.GetString("mysql.url"),
		MysqlPort:    	viper.GetInt("mysql.port"),
		MysqlUserName:  viper.GetString("mysql.username"),
		MysqlPassWord:  viper.GetString("mysql.password"),
		DBName:  		viper.GetString("mysql.db"),
		LogFile:      	viper.GetString("logfile"),
		LogLevel:     	viper.GetString("loglevel"),
		LogCount:     	viper.GetUint("logcount"),
	}
}
