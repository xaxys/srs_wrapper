package database

import (
	"fmt"
	"srs_wrapper/config"
	"srs_wrapper/model"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Cache *cache.Cache
)

func init() {
	dbType := config.AppConfig.GetString("database.driver")
	switch dbType {
	case "mysql":
		DB = initMysql()
	case "sqlite":
		DB = initSqlite()
	default:
		panic(fmt.Errorf("support mysql and sqlite only"))
	}
	initTable()

	Cache = initCache()
}

func initSqlite() *gorm.DB {
	dbPath := config.AppConfig.GetString("database.sqlite.path")
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		panic(fmt.Errorf("No error should happen when connecting to database, but got: %+v", err))
	}
	return db
}

func initMysql() *gorm.DB {
	dbHost := config.AppConfig.GetString("database.mysql.host")
	dbPort := config.AppConfig.GetString("database.mysql.port")
	dbName := config.AppConfig.GetString("database.mysql.name")
	dbParams := config.AppConfig.GetString("database.mysql.params")
	dbUser := config.AppConfig.GetString("database.mysql.user")
	dbPasswd := config.AppConfig.GetString("database.mysql.password")
	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbUser, dbPasswd, dbHost, dbPort, dbName, dbParams)

	db, err := gorm.Open(mysql.Open(dbURL))
	if err != nil {
		panic(fmt.Errorf("No error should happen when connecting to database, but got: %+v", err))
	}
	return db
}

func initTable() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.Permission{},
	)
	if err != nil {
		panic(fmt.Errorf("Unable to sync the struct to database: %+v", err))
	}
}

func initCache() *cache.Cache {
	cacheExpire := config.AppConfig.GetInt64("cache.expire")
	cachePurge := config.AppConfig.GetInt64("cache.purge")

	cache := cache.New(time.Duration(cacheExpire)*time.Second, time.Duration(cachePurge)*time.Second)
	return cache
}
