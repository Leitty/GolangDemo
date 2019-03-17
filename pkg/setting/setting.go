package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize int
	JwtSecret string
	RuntimeRootPath string
	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string
}

var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
	MaxIdleConn int
	MaxOpenConn int
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Eureka struct {
	AppName string
	EurekaServerUrl string
	StatusUrl string
	HealthUrl string
	DataCenterInfo string
	SecurePort string
}

var EurekaSetting = &Eureka{}

var (
	EurekaHomeUrl string
)

// 注意，使用MapTo时，结构体中的名称和对象中的名称要一致，比如结构体中叫RunMode，app.ini中也要叫RunMode，不能写出RUNMODE
func Setup(){
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err  = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting  err: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting  err: %v", err)
	}
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	err = Cfg.Section("eureka").MapTo(EurekaSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo EurekaSetting  err: %v", err)
	}
	EurekaHomeUrl = EurekaSetting.EurekaServerUrl + "apps/" + EurekaSetting.AppName
}
