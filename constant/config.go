package constant

import (
	"fmt"
	"github.com/spf13/viper"
	"gin-template/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	appConfig = initConfig()
	Db        = initMysqlConfig()
)

// Config 配置
type Config struct {
	MySQL MysqlConfig
}

// MysqlConfig MySQL 配置
type MysqlConfig struct {
	Host     string
	Port     uint16
	DbName   string
	Arg      string
	UserName string
	Password string
}

// initConfig 初始化配置
func initConfig() Config {
	// 把配置文件读取到结构体上
	var config Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置失败 " + err.Error())
	}
	// 将配置文件绑定到 config 上
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("自动建表失败 " + err.Error())
	}
	fmt.Println("config", config)
	return config
}

// initMysqlTable 初始化表
func initMysqlTable(db *gorm.DB) (err error) {
	err = db.AutoMigrate(model.User{})
	if err == nil {
		return nil
	}
	return err
}

// initConfig 初始化 MySQL 配置
func initMysqlConfig() *gorm.DB {
	m := appConfig.MySQL
	link := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		m.UserName,
		m.Password,
		m.Host,
		m.Port,
		m.DbName,
		m.Arg,
	)
	config := &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			// 使用单数表名，结构体 User 对应的表名为 user
			SingularTable: true,
		},
	}
	db, err := gorm.Open(mysql.Open(link), config)
	if err != nil {
		panic("连接数据库失败 " + err.Error())
	}
	err = initMysqlTable(db)
	if err != nil {
		panic("初始化表失败 " + err.Error())
	}
	return db
}
