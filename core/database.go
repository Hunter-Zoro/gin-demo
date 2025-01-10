package core

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/models"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

func InitializeDB() *gorm.DB {
	switch global.DEMO_CONFIG.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}

}
func initMySqlGorm() *gorm.DB {
	dbConfig := global.DEMO_CONFIG.Database

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
	}); err != nil {
		global.LOG.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMySqlTables(db)
		return db
	}

}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {

	err := db.Migrator().AutoMigrate(
		&models.User{},
	)
	if err != nil {
		global.LOG.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}

func getGormLogWriter() logger.Writer {
	var writer io.Writer

	if global.DEMO_CONFIG.Database.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   global.DEMO_CONFIG.Log.RootDir + "/" + global.DEMO_CONFIG.Database.LogFilename,
			MaxSize:    global.DEMO_CONFIG.Log.MaxSize,
			MaxBackups: global.DEMO_CONFIG.Log.MaxBackups,
			MaxAge:     global.DEMO_CONFIG.Log.MaxAge,
			Compress:   global.DEMO_CONFIG.Log.Compress,
		}
	} else {
		writer = os.Stdout // 默认 Writer
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.DEMO_CONFIG.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                           // 慢 SQL 阈值
		LogLevel:                  logMode,                                          // 日志级别
		IgnoreRecordNotFoundError: false,                                            // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.DEMO_CONFIG.Database.EnableFileLogWriter, // 禁用彩色打印
	})

}
