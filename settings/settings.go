package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 保存整个程序的配置信息
var Conf = new(AppConfig)

// 作用 : 1. 可以查看配置项的详细信息 很清晰
// 注意 不管是什么格式的文件都是用mapstructure这个来配置
type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Dbname      string `mapstructure:"dbname"`
	Password    string `mapstructure:"password"`
	Port        int    `mapstructure:"port"`
	IdConns     int    `mapstructure:"id_conns"`
	OpenConnect int    `mapstructure:"open_connect"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Dbname   int    `mapstructure:"dbname"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//viper.SetConfigName("config") //加载指定配置文件名 不需要后缀
	//viper.SetConfigType("yaml") //指定配置文件类型 这里也可以使用json 专用于从etcd远程获取配置信息时指定配置文件类型 只会查看config文件 如果有相同的json叫config那就不行
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigFile("./config.yaml")
	viper.AddConfigPath(".")   //指定查找文件的路径 这里使用相对路径
	err = viper.ReadInConfig() //读取配置文件信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() filed %v\n", err)
		return
	}
	//把读取的配置信息反序列化到Conf变量中  程序中就可以读取配置信息了
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal(Conf) is failed%v\n", err)
	}
	viper.WatchConfig()                            // 启动监听配置文件
	viper.OnConfigChange(func(in fsnotify.Event) { // fsnotify进行监听
		// 这里监听后需要继续更新
		fmt.Println("配置文件修改了")
		// 序列化赋值给全局的文件
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal(Conf) is failed%v\n", err)
		}
	})
	return
}
