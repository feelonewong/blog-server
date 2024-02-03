package database

import (
	"blog-server/utils"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type StringContextKey string

var (
	blog_mysql      *gorm.DB // 连接池
	blog_mysql_once sync.Once
	dblog           ormlog.Interface

	blog_redis      *redis.Client
	blog_redis_once sync.Once
)

func init() {
	dblog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		ormlog.Config{
			SlowThreshold: 100 * time.Microsecond, // 超出100ms为slow query
			LogLevel:      ormlog.Info,            // 打印输出的等级
			Colorful:      true,                   // 彩色打印
		},
	)
}

func createMySQLDB(dbname, host, user, pass string, port int) *gorm.DB {
	// 数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	var err error
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true}) // 预编译
	if err != nil {
		fmt.Println("error...")
		utils.LogRus.Panic("connect to mysql use dsn %s failed: %s", dsn, err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) // 数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  // 数据库连接池最大空闲连接数
	//utils.LogRus.Infof("connect to mysql db %s", dbname)
	return db
}

func GetBlogDBConnection() *gorm.DB {
	blog_mysql_once.Do(func() {
		dbName := "blog"
		if blog_mysql == nil {
			viper := utils.CreateConfig("mysql")
			host := viper.GetString(dbName + ".host")
			port := viper.GetInt(dbName + ".port")
			user := viper.GetString(dbName + ".user")
			pass := viper.GetString(dbName + ".pass")
			blog_mysql = createMySQLDB(dbName, host, user, pass, port)
		}
	})

	return blog_mysql
}
