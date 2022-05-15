package database

import (
	"context"
	"fmt"
	"gin-template/server/config"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database is the list of connection of database
type Database struct {
	MySql *gorm.DB
	Redis *redis.Client
}

var Instance *Database

// InitDB is the function to initialize all database
func InitDB() *Database {
	Instance = &Database{
		MySql: initMySql(),
		Redis: initRedis(),
	}
	return Instance
}

// initMySql is the function to initialize mysql connection
func initMySql() *gorm.DB {
	logrus.WithField("server", "MySql").Info("Connect to MySql...")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: MakeDSN(struct {
			Addr string
			User string
			Pass string
			DB   string
		}{
			Addr: config.Cfg.Database.MySql.Host + ":" + config.Cfg.Database.MySql.Port,
			User: config.Cfg.Database.MySql.User,
			Pass: config.Cfg.Database.MySql.Password,
			DB:   config.Cfg.Database.MySql.Database,
		}), // DSN data source name
	}), &gorm.Config{})
	if err != nil {
		logrus.WithField("server", "MySql").Panic("Connect to MySql failed!")
		return nil
	}

	logrus.WithField("server", "MySql").Info("MySql auto migrate...")
	err = db.AutoMigrate(MigrateList()...)
	if err != nil {
		return nil
	}
	return db
}

func initRedis() *redis.Client {
	logrus.WithField("server", "Redis").Info("Connect to Redis...")
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Database.Redis.Host + ":" + config.Cfg.Database.Redis.Port,
		Password: config.Cfg.Database.Redis.Password,
		DB:       config.Cfg.Database.Redis.Database,
	})

	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		logrus.WithField("server", "Redis").Panic("Connect to Redis failed!")
		return nil
	}
	return rdb
}

// MakeDSN is the function to make dsn
func MakeDSN(dbBaseConfig struct {
	Addr string
	User string
	Pass string
	DB   string
}) string {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbBaseConfig.User, dbBaseConfig.Pass, dbBaseConfig.Addr, dbBaseConfig.DB)
}

func MigrateList() []interface{} {
	return []interface{}{
		&AccessToken{},
		&UserInfo{},
	}
}
