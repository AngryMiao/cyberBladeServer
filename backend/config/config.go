package config

/*
	配置模板, 用来做统一配置的入口,
	将代码和配置分离，
*/

const (
	DevMode     = "dev"
	ReleaseMode = "release"
	TestMode    = "test"
)

var (
	Conf     = &Config{}
	confPath string
)

type Config struct {
	Mode string

	// db
	DB *DBConfig

	// redis
	RedisCache *RedisConfig

	App *AppConfig

	JWT *JWTConfig

	OSS *AliOSSConfig

	AllowHosts []string

	RPC *RPCConfig

	OpenAi *OpenAiConfig

	Discord *DiscordConfig

	Doubao *DoubaoConfig
}

type DiscordConfig struct {
	Token string
}

type OpenAiConfig struct {
	Key     string
	BaseUrl string
}

type AppConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	DSN string
	// 数据库类型
	Dialect string
}

type RedisConfig struct {
	Addr       string
	Password   string
	DB         int
	Service    string
	LuaScripts struct {
		CouponActivityNumberExecution    string
		CouponActivityNumberCompensation string
	}
}

type JWTConfig struct {
	PublicKeyPath string
	Key           string
	Issuer        string
}

type AliOSSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	Host            string
	Domain          string
	ExpireTime      int64
}

type RPCConfig struct {
	User struct {
		Address string
	}
}

type DoubaoConfig struct {
	ApiKey string
}
