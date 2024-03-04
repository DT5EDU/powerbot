package webController

import (
	"net/http"
	"strconv"

	"powerbot/core"

	"github.com/gin-gonic/gin"
)

type WebController struct {
	core *core.CoreService
}

func NewWebController(core *core.CoreService) *WebController {
	return &WebController{core: core}
}

func (w *WebController) SettingsHandler(ctx *gin.Context) {
	chatId := ctx.Param("chatId")
	room := ctx.Query("room")
	hour := ctx.Query("hour")

	hourInt, err := strconv.ParseInt(hour, 10, 64)
	if err != nil {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = w.core.Settings(&core.SettingsRequest{ChatId: chatId, RoomNumber: room, Hour: hourInt})
	if err != nil {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func (w *WebController) QueryHandler(ctx *gin.Context) {
	chatId := ctx.Param("chatId")

	err := w.core.QueryPower(&core.QueryPowerRequest{ChatId: chatId})
	if err != nil {
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}
