package config

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
    "os"
    "time"
)

type appConfig struct {
    Env string `yaml:"env"`
}
type wxappConfig struct {
    AppId     string `yaml:"app_id"`
    AppSecret string `yaml:"app_secret"`
}
type domainConfig struct {
    User string `yaml:"user"`
}
type kafkaConfig struct {
    Addrs []string `yaml:"addrs"`
}

type httpConfig struct {
    HttpPort     int           `yaml:"http_port"`
    HttpHost     string        `yaml:"http_host"`
    ReadTimeout  time.Duration `yaml:"read_timeout"`
    WriteTimeout time.Duration `yaml:"write_timeout"`
    RunMode      string        `yaml:"run_mode"`
}

type mysqlConfig struct {
    Database string `yaml:"database"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    Timeout  string `yaml:"timeout"`
    Charset  string `yaml:"charset"`
    MaxIdle  int    `yaml:"max_idle"`
    MaxOpen  int    `yaml:"max_open"`
    RunMode  string `yaml:"run_mode"`
}
type uploadConfig struct {
    ImagePrefixUrl string `yaml:"image_prefix_url"`
    ImageSavePath  string `yaml:"image_save_path"`

    ImageMaxSize   int      `yaml:"image_max_size"`
    ImageAllowExts []string `yaml:"image_allow_exts"`

    RuntimeRootPath string `yaml:"runtime_root_path"`
}
type redisConfig struct {
    Network            string        `yaml:"network"`              // 网络类型，tcp or unix，默认tcp
    Addr               string        `yaml:"addr"`                 // 主机名+冒号+端口，默认localhost:6379
    Password           string        `yaml:"password"`             // 密码
    DB                 int           `yaml:"db"`                   // redis数据库index
    PoolSize           int           `yaml:"pool_size"`            // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
    MinIdleConns       int           `yaml:"min_idle_conns"`       // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
    DialTimeout        time.Duration `yaml:"dial_timeout"`         // 连接建立超时时间，默认5秒。
    ReadTimeout        time.Duration `yaml:"read_timeout"`         // 读超时，默认3秒， -1表示取消读超时
    WriteTimeout       time.Duration `yaml:"write_timeout"`        // 写超时，默认等于读超时
    PoolTimeout        time.Duration `yaml:"pool_timeout"`         // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
    IdleCheckFrequency time.Duration `yaml:"idle_check_frequency"` // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
    IdleTimeout        time.Duration `yaml:"idle_timeout"`         // 闲置超时，默认5分钟，-1表示取消闲置超时检查
    MaxConnAge         time.Duration `yaml:"max_conn_age"`         // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
    MaxRetries         int           `yaml:"max_retries"`          // 命令执行失败时，最多重试多少次，默认为0即不重试
    MinRetryBackoff    time.Duration `yaml:"min_retry_backoff"`    // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
    MaxRetryBackoff    time.Duration `yaml:"max_retry_backoff"`    // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

}
type loggerConfig struct {
    MaxSize     int    `yaml:"max_size"`
    MaxAge      int    `yaml:"max_age"`
    AtomicLevel string `yaml:"atomic_level"`
    FilePath    string `yaml:"file_path"`
}
type config struct {
    Http   httpConfig   `yaml:"http"`
    Mysql  mysqlConfig  `yaml:"mysql"`
    Redis  redisConfig  `yaml:"redis"`
    Kafka  kafkaConfig  `yaml:"kafka"`
    Upload uploadConfig `yaml:"upload"`
    Domain domainConfig `yaml:"domain"`
    Wxapp  wxappConfig  `yaml:"wxapp"`
    App    appConfig    `yaml:"app"`
    Logger loggerConfig `yaml:"logger"`
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

    yamlFile, err := ioutil.ReadFile("config/" +configEnv+".yaml")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'config/"+configEnv+".yaml': %v", err)
        os.Exit(1)
    }
    c := &config{}
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        fmt.Println(err.Error())
    }
    Http = c.Http
    Mysql = c.Mysql
    Redis = c.Redis
    Kafka = c.Kafka
    Upload = c.Upload
    Domain = c.Domain
    Wxapp = c.Wxapp
    App = c.App
    Logger = c.Logger
    fmt.Printf("%+v",c);
    return

}
