// 配置文件错误处理
package config

// 配置错误信息
type ConfigError string

const (
	ConfigErrorInvalidKey            ConfigError = "C10010:ConfigErrorInvalidKey,无效的key(%s)"
	ConfigErrorInvalidTypeConvertion ConfigError = "C10020:ConfigErrorInvalidTypeConvertion,无效的类型转换,string(%s)转换为%s"
	ConfigErrorInvalidConfigKind     ConfigError = "C10030:ConfigErrorInvalidConfigKind,无效的配置类型(%s)"
	ConfigErrorReadError             ConfigError = "C10040:ConfigErrorReadError,无法读取配置文件(%s)"
	ConfigErrorNotMatch              ConfigError = "C10050:ConfigErrorNotMatch,未找到匹配的字符(%s)"
)
