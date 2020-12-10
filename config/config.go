package config

import (
    "strconv"
    "time"
)

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
)

func init() {
    Http.HttpHost = "127.0.0.1"
    Http.HttpPort = 8083
    Http.ReadTimeout = 60 * time.Second
    Http.WriteTimeout = 60 * time.Second

    Mysql.Database = "my_test"
    Mysql.Username = "root"
    Mysql.Password = "123456"
    Mysql.Host = "127.0.0.1"
    Mysql.Port = 3306
    Mysql.Timeout = "5000ms"
    Mysql.Charset = "utf8mb4"
    Mysql.MaxIdle = 10
    Mysql.MaxOpen = 20

    Redis.Network = "tcp"                          // 网络类型，tcp or unix，默认tcp
    Redis.Addr = "127.0.0.1:6379"                  // 主机名+冒号+端口，默认localhost:6379
    Redis.Password = ""                            // 密码
    Redis.DB = 0                                   // redis数据库index
    Redis.PoolSize = 4                             // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
    Redis.MinIdleConns = 2                         // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
    Redis.DialTimeout = 5 * time.Second            // 连接建立超时时间，默认5秒。
    Redis.ReadTimeout = 3 * time.Second            // 读超时，默认3秒， -1表示取消读超时
    Redis.WriteTimeout = 3 * time.Second           // 写超时，默认等于读超时
    Redis.PoolTimeout = 4 * time.Second            // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
    Redis.IdleCheckFrequency = 60 * time.Second    // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
    Redis.IdleTimeout = 5 * time.Minute            // 闲置超时，默认5分钟，-1表示取消闲置超时检查
    Redis.MaxConnAge = 0 * time.Second             // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
    Redis.MaxRetries = 1                           // 命令执行失败时，最多重试多少次，默认为0即不重试
    Redis.MinRetryBackoff = 8 * time.Millisecond   // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
    Redis.MaxRetryBackoff = 512 * time.Millisecond // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

    Kafka.Addrs = []string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"}

    Upload.ImagePrefixUrl = "http:" + Http.HttpHost + ":" + strconv.Itoa(Http.HttpPort)
    Upload.ImageSavePath = "upload/images/"
    Upload.ImageMaxSize = 5 * 1024 * 1024
    Upload.ImageAllowExts = []string{".jpg", ".jpeg", ".png"}
    Upload.RuntimeRootPath = "runtime/"

}
