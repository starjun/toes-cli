package toes

import (
	"errors"
	"github.com/starjun/gotools/v2"
	"log"
	"os"
	"path"
	"strings"
)

var confTpl = `
# toes-server 全配置

# REST-ful Server
server:
  mode: debug # server mode: release, debug, test，默认 release
  addr: :8080

# MySQL
mysql:
  host: 127.0.0.1 # MySQL 机器 ip 和端口，默认 127.0.0.1:3306
  username: root # MySQL 用户名(建议授权最小权限集)
  password:  # MySQL 用户密码
  database: demo # 系统所用的数据库名
  maxIdleConnections: 100 # MySQL 最大空闲连接数，默认 100
  maxOpenConnections: 100 # MySQL 最大打开的连接数，默认 100
  maxConnectionLifeTime: 10s # 空闲连接最大存活时间，默认 10s
  logLevel: 4 # GORM log level, 1: silent, 2:error, 3:warn, 4:info

# Redis
redis:
  host: 127.0.0.1:6379 # redis 地址，默认 127.0.0.1:6379
  username: # 用户名
  password:  # redis 密码


log:
  level: debug # 日志级别，优先级从低到高依次为：debug, info, warn, error, dpanic, panic, fatal。
  days: 7 # 日志保留天数
  format: json # 支持的日志输出格式，目前支持 console 和json两种。raw 其实就是text格式。
  console: true # 是否同步输出到命令行
  path: ./log.log


# 加密之后
seckey:
  basekey: ""
  jwtKey: ""
  jwtttl: 1024 # token 过期时间(分钟)
  pproftoken : on # 配置访问 pprof 是否启用 token 检查


# 防重放
checkHeader:
  all: false
  nonce: true
  nonceCacheSeconds: 30
  time: true
  seconds: 120
  sign: true
  key: ""

# 请求头相关配置
header:
  realip: x-realip-from
  requestid: x-request-id
`

func (b *ToesGen) ConfFileGen() error {
	filepath := path.Join(b.confPath, "toes.config.yaml")
	re, err := gotools.PathExists(filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	if re {
		log.Println(filepath, "已存在")
		return errors.New(filepath + " 已存在")
	}

	f, e := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 777)
	if e != nil {
		log.Println(e)
		return e
	}
	confTpl = strings.Replace(confTpl, "{{at}}", "`", -1)
	f.WriteString(confTpl)
	f.Close()
	return nil
}
