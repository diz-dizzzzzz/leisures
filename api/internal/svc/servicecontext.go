package svc

import (
	"strings"

	"acupofcoffee/api/internal/config"
	"acupofcoffee/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initDB(c.MySQL)

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

func initDB(cfg config.MySQLConfig) *gorm.DB {
	var db *gorm.DB
	var err error

	// 如果 DataSource 以 sqlite: 开头，使用 SQLite
	if strings.HasPrefix(cfg.DataSource, "sqlite:") {
		dbPath := strings.TrimPrefix(cfg.DataSource, "sqlite:")
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		db, err = gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if !strings.HasPrefix(cfg.DataSource, "sqlite:") {
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to get database instance: " + err.Error())
		}
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	// Auto migrate models
	if err := db.AutoMigrate(
		&model.User{},
		&model.Article{},
		&model.ArticleVersion{},
		&model.ArticleDraft{},
	); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}
