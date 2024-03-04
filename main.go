package main

import (
	"net/http"

	"powerbot/config"
	"powerbot/core"
	"powerbot/cron"
	"powerbot/dao"
	"powerbot/larkListen"
	"powerbot/larkSend"
	"powerbot/webController"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()

	// 数据层
	db := dao.NewDatabase(
		conf.Database,
	)
	sendService := larkSend.NewLarkSendService(
		conf.FeishuAppId,
		conf.FeishuAppSecret,
	)

	// 业务层
	coreService := core.NewCoreService(
		db,
		sendService,
	)

	// 控制层
	webController := webController.NewWebController(
		coreService,
	)
	listenService := larkListen.NewLarkListenService(
		conf.FeishuAppId,
		conf.FeishuAppSecret,
		conf.VerificationToken,
		conf.EventKey,
		coreService,
	)
	cronService := cron.NewCron(
		coreService,
	)

	cronService.Scheduler.Start()

	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/api/query/:chatId", webController.QueryHandler)
	r.GET("/api/settings/:chatId", webController.SettingsHandler)

	r.POST("/api/lark/event", listenService.EventHandler)
	r.POST("/api/lark/card", listenService.CardHandler)

	r.GET("/go/topup", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "alipays://platformapi/startapp?appId=10000007&qrcode=http%3A%2F%2Fwx.weiweixiao.net%2Findex.php%2FWap%2FModZhjf%2FitemList%2Ftoken%2FusgMkdVR6BGAAAAWPwAVGQ%2Fid%2FskIWNmy96BGAAAAWPwAVGQ.html")
	})

	r.Run("127.0.0.1:6666")
}
