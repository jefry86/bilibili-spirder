//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/16 9:05 下午

package utils

import (
	"github.com/spf13/viper"
)

var Config Cfg

type Cfg struct {
	Rate int
	Path
	Http
	Db
	Search
}

type Path struct {
	Logs  string `json:"logs"`
	Video string `json:"video"`
	Csv   string `json:"csv"`
}

type Http struct {
	Timeout int  `json:"timeout"`
	Debug   bool `json:"debug"`
}

type Db struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Search struct {
	Keyword string `json:"keyword"`
}

func LoadCfg() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
	//Config.Path.Logs = viper.GetString("path.logs")
	//Config.Path.Video = viper.GetString("path.video")
	//Config.Path.Csv = viper.GetString("path.csv")
	//Config.Http.Timeout = viper.GetInt("http.timeout")
	//Config.Rate = viper.GetInt("rate")
}
