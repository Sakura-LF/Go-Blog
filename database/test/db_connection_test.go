package test

import (
	"blog/database"
	"log"
	"sync"
	"testing"
)

// 测试高并发下gorm.DB是否还是单例
func TestGetBlogDBConnection(t *testing.T) {
	const C = 100
	wg := sync.WaitGroup{}
	wg.Add(C)
	for i := 0; i < C; i++ {
		go func() {
			defer wg.Done()
			database.GetBlogDBConnection()
		}()
	}
	wg.Wait()
}

// go test -v .\database\test\ -run=^TestGetBlogDBConnection$ -count=1
// 自动迁移数据表
func TestDB(t *testing.T) {
	connection := database.GetBlogDBConnection()

	if err := connection.AutoMigrate(&database.User{}); err != nil {
		log.Fatalln(err)
	}

	if err := connection.AutoMigrate(&database.Blog{}); err != nil {
		log.Fatalln(err)
	}
}

func TestRedis(t *testing.T) {
	database.GetRedisClient()

}
