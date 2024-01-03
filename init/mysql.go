package init

import (
	"Go-Blog/service"
	"Go-Blog/util"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// DB 数据库连接池
var DB *gorm.DB

func InitMysql() *gorm.DB {
	// 获取配置
	// 配置 dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", MysqlConfig.Get("mysql.user"), MysqlConfig.Get("mysql.pass"), MysqlConfig.Get("mysql.host"), MysqlConfig.Get("mysql.port"), MysqlConfig.Get("mysql.dbname")) //mb4兼容emoji表情符号
	fmt.Println(dsn)
	// 配置 gorm 日志
	rootPath := util.GetRootPath()
	mysql_log, _ := os.OpenFile(rootPath+"/logs/mysql.log", os.O_CREATE|os.O_APPEND, 0644)
	gormLoggerConfig := logger.New(
		log.New(mysql_log, "\r\n", log.LstdFlags),
		logger.Config{
			// 慢查询阈值
			SlowThreshold: MysqlConfig.GetDuration("gorm.log.SlowThreshold"),
			// 不彩色化
			Colorful: MysqlConfig.GetBool("gorm.log.Colorful"),
			// 日志级别
			LogLevel: logger.LogLevel(MysqlConfig.GetInt("gorm.log.LogLevel")),
		})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		Logger:                 gormLoggerConfig,
	})
	if err != nil {
		Logger.Panic("fail to op db", zap.Error(err))
	}
	s, err := db.DB()
	if err != nil {
		Logger.Panic("err", zap.Error(err))
	}
	s.SetMaxOpenConns(MysqlConfig.GetInt("gorm.db.MaxOpenConn")) // 数据库连接池最大连接数
	s.SetMaxIdleConns(MysqlConfig.GetInt("gorm.db.MaxIdleConn")) // 数据库最大允许的空闲丽娜姐
	DB = db
	Logger.Info("Mysql Init success")
	return DB
}

func Migrate() {
	user := service.User{}
	if err := DB.AutoMigrate(&user); err != nil {
		Logger.Panic("automigrate table Blog fail", zap.Error(err))
	} else {
		Logger.Info("Table migrate success", zap.Any("table", user))
	}
	blog := service.Blog{}
	if err := DB.AutoMigrate(&blog); err != nil {
		Logger.Panic("automigrate table User fail", zap.Error(err))
	} else {
		Logger.Info("Table migrate success", zap.Any("table", blog))
	}
	Logger.Info("AutoMigrate success")
}
