package config

import (
    "github.com/go-ini/ini"
    "log"
    "os"
    "time"
)

type WxappConfig struct {
    AppId     string
    AppSecret string
}
type DomainConfig struct {
    User string
}
type KafkaConfig struct {
    Addrs []string
}

type HttpConfig struct {
    HttpPort     int
    HttpHost     string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

type MysqlConfig struct {
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
type UploadConfig struct {
    ImagePrefixUrl string
    ImageSavePath  string

    ImageMaxSize   int
    ImageAllowExts []string

    RuntimeRootPath string
}
type RedisConfig struct {
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

var (
    Http   HttpConfig
    Mysql  MysqlConfig
    Redis  RedisConfig
    Kafka  KafkaConfig
    Upload UploadConfig
    Domain DomainConfig
    Wxapp  WxappConfig
)

func Setup() {
    cfg, err := ini.Load("config/dev.ini")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'config/dev.ini': %v", err)
        os.Exit(1)
    }

    if httpCfg, err := cfg.GetSection("http"); err == nil {

        Http.HttpHost = httpCfg.Key("HttpHost").MustString("")
        Http.HttpPort = httpCfg.Key("HttpPort").MustInt(0)
        Http.ReadTimeout = time.Duration(httpCfg.Key("ReadTimeout").MustInt(0)) * time.Second
        Http.WriteTimeout = time.Duration(httpCfg.Key("WriteTimeout").MustInt(0)) * time.Second
    }


    if mysqlCfg, err := cfg.GetSection("mysql"); err == nil {
        Mysql.Database = mysqlCfg.Key("Database").MustString("")
        Mysql.Username = mysqlCfg.Key("Username").MustString("")
        Mysql.Password = mysqlCfg.Key("Password").MustString("")
        Mysql.Host = mysqlCfg.Key("Host").MustString("")
        Mysql.Port = mysqlCfg.Key("Port").MustInt(0)
        Mysql.Timeout = mysqlCfg.Key("Timeout").MustString("")
        Mysql.Charset = mysqlCfg.Key("Charset").MustString("")
        Mysql.MaxIdle = mysqlCfg.Key("MaxIdle").MustInt(0)
        Mysql.MaxOpen = mysqlCfg.Key("MaxOpen").MustInt(0)
    }

    if redisCfg, err := cfg.GetSection("redis"); err == nil {
        Redis.Network = redisCfg.Key("Network").MustString("")                                                // 网络类型，tcp or unix，默认tcp
        Redis.Addr = redisCfg.Key("Addr").MustString("")                                                      // 主机名+冒号+端口，默认localhost:6379
        Redis.Password = redisCfg.Key("Password").MustString("")                                              // 密码
        Redis.DB = redisCfg.Key("DB").MustInt(0)                                                              // redis数据库index
        Redis.PoolSize = redisCfg.Key("PoolSize").MustInt(0)                                                  // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
        Redis.MinIdleConns = redisCfg.Key("MinIdleConns").MustInt(0)                                          // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
        Redis.DialTimeout = time.Duration(redisCfg.Key("DialTimeout").MustInt(0)) * time.Second               // 连接建立超时时间，默认5秒。
        Redis.ReadTimeout = time.Duration(redisCfg.Key("ReadTimeout").MustInt(1)) * time.Second               // 读超时，默认3秒， -1表示取消读超时
        Redis.WriteTimeout = time.Duration(redisCfg.Key("WriteTimeout").MustInt(0)) * time.Second             // 写超时，默认等于读超时
        Redis.PoolTimeout = time.Duration(redisCfg.Key("PoolTimeout").MustInt(0)) * time.Second               // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
        Redis.IdleCheckFrequency = time.Duration(redisCfg.Key("IdleCheckFrequency").MustInt(0)) * time.Second // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
        Redis.IdleTimeout = time.Duration(redisCfg.Key("IdleTimeout").MustInt(0)) * time.Minute               // 闲置超时，默认5分钟，-1表示取消闲置超时检查
        Redis.MaxConnAge = time.Duration(redisCfg.Key("MaxConnAge").MustInt(0)) * time.Second                 // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
        Redis.MaxRetries = redisCfg.Key("MaxRetries").MustInt(0)                                              // 命令执行失败时，最多重试多少次，默认为0即不重试
        Redis.MinRetryBackoff = time.Duration(redisCfg.Key("MinRetryBackoff").MustInt(0)) * time.Millisecond  // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
        Redis.MaxRetryBackoff = time.Duration(redisCfg.Key("MaxRetryBackoff").MustInt(0)) * time.Millisecond  // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
    }

    if kafkaCfg, err := cfg.GetSection("kafka"); err == nil {
        Kafka.Addrs = kafkaCfg.Key("Addrs").Strings(",")
    }

    if uploadCfg, err := cfg.GetSection("upload"); err == nil {
        Upload.ImagePrefixUrl = uploadCfg.Key("ImagePrefixUrl").MustString("")
        Upload.ImageSavePath = uploadCfg.Key("ImageSavePath").MustString("")
        Upload.ImageMaxSize = uploadCfg.Key("ImageMaxSize").MustInt(1) * 1024 * 1024
        Upload.ImageAllowExts = uploadCfg.Key("ImageAllowExts").Strings(",")
        Upload.RuntimeRootPath = uploadCfg.Key("RuntimeRootPath").MustString("")
    }

    if domainCfg, err := cfg.GetSection("domain"); err == nil {
        Domain.User = domainCfg.Key("User").MustString("")
    }

    if wxappCfg, err := cfg.GetSection("wxapp"); err == nil {
        Wxapp.AppId = wxappCfg.Key("AppId").MustString("")
        Wxapp.AppSecret = wxappCfg.Key("AppSecret").MustString("")
    }

}
