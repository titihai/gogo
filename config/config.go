// package config 实现一个ini解析器
package config

import (
	"sync"
)

// Config 配置信息存储接口
type Config interface {
	// GlobalSection 获取全局配置段
	GlobalSection() Section
	// Section 获取特定配置段
	Section(name string) (Section, error)
}

// Section 配置信息段接口
type Section interface {
	// Name 获取配置端名
	Name() string
	// Int 根据key获取整数
	Int(key string) (int64, error)
	// Bool 获取布尔值
	Bool(key string) (bool, error)
	// Float 获取浮点值
	Float(key string) (float64, error)
	// String 获取字符串
	String(key string) (string, error)
}

// ConfigParser 配置解析器
type ConfigParser func([]bytes) (Config,error)

var (
	parserMu sync.Mutex	//生成器互斥锁
	parsers = make(map(ConfigType)ConfigParser) //配置文件解析器
)


