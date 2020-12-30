package config

import (
    "fmt"
    "github.com/go-ini/ini"
    "log"
    "os"
    "time"
)

type appConfig struct {
    Env string
}
type wxappConfig struct {
    AppId     string
    AppSecret string
}
type domainConfig struct {
    User string
}
type kafkaConfig struct {
    Addrs []string
}

type httpConfig struct {
    HttpPort     int
    HttpHost     string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

type mysqlConfig struct {
    Database string
    Username string
    Password string
    Host     string
    Port     int
    Timeout  string
    Charset  string
    MaxIdle  int
    MaxOpen  int
}
type uploadConfig struct {
    ImagePrefixUrl string
    ImageSavePath  string

    ImageMaxSize   int
    ImageAllowExts []string

    RuntimeRootPath string
}
type redisConfig struct {
    Network            string        // 网络类型，tcp or unix，默认tcp
    Addr               string        // 主机名+冒号+端口，默认localhost:6379
    Password           string        // 密码
    DB                 int           // redis数据库index
    PoolSize           int           // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
    MinIdleConns       int           // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
    DialTimeout        time.Duration // 连接建立超时时间，默认5秒。
    ReadTimeout        time.Duration // 读超时，默认3秒， -1表示取消读超时
    WriteTimeout       time.Duration // 写超时，默认等于读超时
    PoolTimeout        time.Duration // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
    IdleCheckFrequency time.Duration // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
    IdleTimeout        time.Duration // 闲置超时，默认5分钟，-1表示取消闲置超时检查
    MaxConnAge         time.Duration // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
    MaxRetries         int           // 命令执行失败时，最多重试多少次，默认为0即不重试
    MinRetryBackoff    time.Duration // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
    MaxRetryBackoff    time.Duration // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

}
type loggerConfig struct {
    MaxSize     int
    MaxAge      int
    AtomicLevel string
    FilePath    string
}

var (
    Http   httpConfig
    Mysql  mysqlConfig
    Redis  redisConfig
    Kafka  kafkaConfig
    Upload uploadConfig
    Domain domainConfig
    Wxapp  wxappConfig
    App    appConfig
    Logger loggerConfig
)

func Setup() {
    configEnv := os.Getenv("CONFIG_ENV")
    if configEnv == "" {
        configEnv = "dev"
    }
    fmt.Println("当前运行加载配置环境,CONFIG_ENV", configEnv)
    cfg, err := ini.Load("config/" + configEnv + ".ini")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'config/"+configEnv+".ini': %v", err)
        os.Exit(1)
    }
    // 基础配置
    err = cfg.Section("app").MapTo(&App)
    if err != nil {
        panic(err)
    }

    // http服务配置
    err = cfg.Section("http").MapTo(&Http)
    if err == nil {
        Http.ReadTimeout = Http.ReadTimeout * time.Second
        Http.WriteTimeout = Http.WriteTimeout * time.Second
    } else {
        os.Exit(2)
    }

    // mysql配置
    err = cfg.Section("mysql").MapTo(&Mysql)

    // redis配置
    err = cfg.Section("redis").MapTo(&Redis)
    if err == nil {
        Redis.DialTimeout = Redis.DialTimeout * time.Second               // 连接建立超时时间，默认5秒。
        Redis.ReadTimeout = Redis.ReadTimeout * time.Second               // 读超时，默认3秒， -1表示取消读超时
        Redis.WriteTimeout = Redis.WriteTimeout * time.Second             // 写超时，默认等于读超时
        Redis.PoolTimeout = Redis.PoolTimeout * time.Second               // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
        Redis.IdleCheckFrequency = Redis.IdleCheckFrequency * time.Second // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
        Redis.IdleTimeout = Redis.IdleTimeout * time.Minute               // 闲置超时，默认5分钟，-1表示取消闲置超时检查
        Redis.MaxConnAge = Redis.MaxConnAge * time.Second                 // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接         // 命令执行失败时，最多重试多少次，默认为0即不重试
        Redis.MinRetryBackoff = Redis.MinRetryBackoff * time.Millisecond  // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
        Redis.MaxRetryBackoff = Redis.MaxRetryBackoff * time.Millisecond  // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
    }

    // kafka配置
    cfg.Section("kafka").MapTo(&Kafka)


    // 文件上传配置
    err = cfg.Section("upload").MapTo(&Upload)
    if err == nil {
        Upload.ImageMaxSize = Upload.ImageMaxSize * 1024 * 1024
    }

    // 相关请求域名
    cfg.Section("domain").MapTo(&Domain)

    // 微信小程序配置
    cfg.Section("wxapp").MapTo(&Wxapp)

    //日志配置
    err = cfg.Section("logger").MapTo(&Logger)
    if err != nil {
        os.Exit(3)
    }

}
