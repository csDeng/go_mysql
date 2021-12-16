package main

// 定时任务
import (
	"log"
	"time"

	"github.com/csDeng/go-gin-example/models"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	// 根据本地时间创建一个新（空白）的 Cron job runner
	c := cron.New()

	// AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
	// 定时硬删除
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	// 在当前执行的程序中启动 Cron 调度程序。
	c.Start()

	// 创建新的定时器
	t1 := time.NewTimer(time.Second * 10)
	// go 中没有while
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
