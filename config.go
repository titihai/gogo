package tinygo

import (
	"path/filepath"
	"strings"
)

// 规范化文件路径
// \test\index.html  ->  /test/index.html
func NormalizePath(path string) string {
	return strings.Trim(filepath.ToSlash(path, "/"))
}

// 基本配置
var tinyConfig = struct {
	appName       string   // 应用名称
	path          string   // 应用启动路径
	mode          string   //启动模式,可以为debug或release
	https         bool     //是否启用https,可选,默认为false
	port          uint16   //监听端口,可选,默认为80，https为true则默认为443
	cert          string   //证书(PEM)路径,如果启用了https则必填
	pkey          string   //私钥(PEM)路径,如果启用了https则必填
	home          string   //首页地址
	session       bool     //是否启用session
	sessiontype   string   //session类型,参考tinygo/session,默认为memory
	sessionexpire int64    //session过期时间,单位为秒
	csrf          bool     //是否启用csrf
	csrfexpire    int64    //csrf token过期时间
	static        []string //静态文件目录,默认为"content"
	view          string   //视图文件目录,默认为"views"
	pageerr       string   //默认错误页面路径,默认为空
	precompile    bool     //是否预编译页面路径,默认为false
	api           string   //使用Api返回的数据的解析格式,默认为auto
	//自动设置项
	sessionName string //session对应的Cookie名称 app+DefaultSessionCookieName
	csrfName    string //csrf对应的Cookie名称 app+DefaultCSRFCookieName
}{}

// 判断是否为发布模式
func IsRealse() bool {
	return tinyConfig.mode == "realse"
}
