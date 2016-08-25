package config

import (
	"strconv"
	"strings"
)

// 导包时自动注册解析器
var err = registerParser(ConfigTypeIni, configParserIni)

// configParserIni ini类型配置文件解析器
func configParserIni(data []byte) (Config, error) {
	var config = newIniConfig()
	var currentSection = config.GlobalSection().(*IniSection) //GlobalSection().(*IniSection)
	var length = len(data)
	// 去除BOM头
	if length >= 3 && data[0] == 239 && data[1] == 187 && data[2] == 191 {
		length -= 3
		data = data[3:]
	}
	for i := 0; i < length; i++ {
		var char = string(data[i])
		//跳过空白字符
		if char == " " || char == "\n" || char == "\r" {
			continue
		}
		switch char {
		case "#", ";":
			{
				// 处理 # 和 ； 开头的注释文本
				for j := i + 1; j < length; j++ {
					if string(data[j]) == "\n" {
						i = j
						break
					}
				}
			}
		case "[":
			{
				var startPos = i + 1
				var isFound = false
				for j := i + 1; j < length; j++ {
					if string(data[j]) == "]" {
						var sectionName = string(data[startPos:j])
						var section = newIniSection(sectionName)
						config.sections[sectionName] = section
						currentSection = section
						isFound = true
					}
					if string(data[j]) == "\r" || j == length-1 {
						if !isFound {
							return nil, ConfigErrorReadError.Format("[").Error()
						}
						i = j
						break
					}
				}
			}
		default:
			{
				var key, value string
				var keyPos int
				var isFound = false
				for j := i + 1; j < length; j++ {
					if string(data[j]) == "=" && !isFound {
						key = string(data[i:j])
						keyPos = j
						isFound = true
					}
					if string(data[j]) == "\r" || j == length-1 {
						if !isFound {
							return nil, ConfigErrorNotMatch.Format("=").Error()
						}
						value = string(data[keyPos+1 : j])
						key = strings.TrimSpace(key)
						value = strings.TrimSpace(value)
						currentSection.add(key, value)
						i = j
						break
					}
				}
			}
		}
	}
	return config, nil
}

// Ini配置
type IniConfig struct {
	globalSection *IniSection            //全局配置段
	sections      map[string]*IniSection //命名配置段集合
}

// 创建IniConfig
func newIniConfig() *IniConfig {
	return &IniConfig{newIniSection(""),
		make(map[string]*IniSection)}
}

// 获取全局配置段
func (this *IniConfig) GlobalSection() Section {
	return this.globalSection
}

// 获取name对应命名字段
func (this *IniConfig) Section(name string) (Section, bool) {
	var section, ok = this.sections[name]
	if !ok {
		return nil, false
	}
	return section, true
}

// Ini配置段
type IniSection struct {
	name string            //段名称
	kvs  map[string]string //键值map
}

// 创建IniSection
func newIniSection(name string) *IniSection {
	return &IniSection{name, make(map[string]string)}
}

// 获取配置段名称
func (this *IniSection) Name() string {
	return this.name
}

// 添加键值对
func (this *IniSection) add(key, value string) {
	this.kvs[key] = value
}

// 获取string值
func (this *IniSection) String(key string) (string, error) {
	var value, ok = this.kvs[key]
	if !ok {
		return "", ConfigErrorInvalidKey.Format(key).Error()
	}
	return value, nil
}

// 获取int值
func (this *IniSection) Int(key string) (int64, error) {
	var value, ok = this.kvs[key]
	if !ok {
		return 0, ConfigErrorInvalidKey.Format(key).Error()
	}
	var i, err = strconv.ParseInt(value, 0, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// 获取float64值
func (this *IniSection) Float(key string) (float64, error) {
	var value, ok = this.kvs[key]
	if ok {
		var f, err = strconv.ParseFloat(value, 64)
		if err == nil {
			return f, nil
		}
		return 0.0, err
	}
	return 0.0, ConfigErrorInvalidKey.Format(key).Error()
}

// 获取布尔值
func (this *IniSection) Bool(key string) (bool, error) {
	var value, ok = this.kvs[key]
	if ok {
		var b, err = strconv.ParseBool(value)
		if err == nil {
			return b, nil
		}
		return false, err
	}
	return false, ConfigErrorInvalidKey.Format(key).Error()
}
