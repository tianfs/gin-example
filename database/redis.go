package database

import (
    "fmt"
    "github.com/go-redis/redis"
    "gin-example/config"
)

var Redis *redis.Client

func init() {
    Redis = redis.NewClient(&redis.Options{
        Network:            config.Redis.Network,            // 网络类型，tcp or unix，默认tcp
        Addr:               config.Redis.Addr,               // 主机名+冒号+端口，默认localhost:6379
        Password:           config.Redis.Password,           // 密码
        DB:                 config.Redis.DB,                 // redis数据库index
        PoolSize:           config.Redis.PoolSize,           // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
        MinIdleConns:       config.Redis.MinIdleConns,       // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
        DialTimeout:        config.Redis.DialTimeout,        // 连接建立超时时间，默认5秒。
        ReadTimeout:        config.Redis.ReadTimeout,        // 读超时，默认3秒， -1表示取消读超时
        WriteTimeout:       config.Redis.WriteTimeout,       // 写超时，默认等于读超时
        PoolTimeout:        config.Redis.PoolTimeout,        // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
        IdleCheckFrequency: config.Redis.IdleCheckFrequency, // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
        IdleTimeout:        config.Redis.IdleTimeout,        // 闲置超时，默认5分钟，-1表示取消闲置超时检查
        MaxConnAge:         config.Redis.MaxConnAge,         // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
        MaxRetries:         config.Redis.MaxRetries,         // 命令执行失败时，最多重试多少次，默认为0即不重试
        MinRetryBackoff:    config.Redis.MinRetryBackoff,    // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
        MaxRetryBackoff:    config.Redis.MinRetryBackoff,    // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
        // 钩子函数
        OnConnect: func(conn *redis.Conn) error { // 仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
            fmt.Printf("redis OnConnect =%v\n", conn)
            return nil
        },
    })

    pong, err := Redis.Ping().Result()
    if err == redis.Nil {
        fmt.Printf("redis异常")
    } else if err != nil {
        fmt.Printf("redis失败:", err)
    } else {
        fmt.Printf("redis成功:", pong)
    }

}
