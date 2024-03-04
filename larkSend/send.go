package larkSend

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type LarkSendService struct {
	client *lark.Client
}

func NewLarkSendService(appId string, appSecret string) *LarkSendService {
	return &LarkSendService{client: lark.NewClient(appId, appSecret)}
}

func (l *LarkSendService) sendMsg(msgType string, msgReceiver string, msgContent string) error {
	// 创建请求对象
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeChatId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(msgReceiver).
			MsgType(msgType).
			Content(msgContent).
			Uuid(uuid.NewString()).
			Build()).
		Build()

	// 发起请求
	resp, err := l.client.Im.Message.Create(context.Background(), req)
	// 处理错误
	if err != nil {
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return err
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
	return nil
}

func (l *LarkSendService) SendPower(chatId string, roomNumber string, balance int64, power int64, lastQuery time.Time, lastBalance int64, lastPower int64) error {
	msg, err := json.Marshal(map[string]any{
		"type": "template",
		"data": map[string]any{
			"template_id": "ctp_AA1O0dNH0beS",
			"template_variable": map[string]string{
				"roomNumber":    roomNumber,
				"time":          time.Now().Format("2006-01-02 15:04"),
				"power":         fmt.Sprintf("%.2f", float64(power)/1000.0),
				"balance":       fmt.Sprintf("%.2f", float64(balance)/100.0),
				"delta_time":    fmt.Sprintf("%.1f", time.Since(lastQuery).Hours()),
				"delta_power":   fmt.Sprintf("%+.2f", float64(power-lastPower)/1000.0),
				"delta_balance": fmt.Sprintf("%+.2f", float64(balance-lastBalance)/100.0),
			},
		},
	})
	if err != nil {
		return err
	}
	l.sendMsg(larkim.MsgTypeInteractive, chatId, string(msg))
	return nil
}

func (l *LarkSendService) SendPowerFirstTime(chatId string, roomNumber string, balance int64, power int64) error {
	msg, err := json.Marshal(map[string]any{
		"type": "template",
		"data": map[string]any{
			"template_id": "ctp_AA1OV7S7CwfE",
			"template_variable": map[string]string{
				"roomNumber": roomNumber,
				"time":       time.Now().Format("2006-01-02 15:04"),
				"power":      fmt.Sprintf("%.2f", float64(power)/1000.0),
				"balance":    fmt.Sprintf("%.2f", float64(balance)/100.0),
			},
		},
	})
	if err != nil {
		return err
	}
	l.sendMsg(larkim.MsgTypeInteractive, chatId, string(msg))
	return nil
}

func (l *LarkSendService) SendMenu(chatId string) error {
	msg, err := json.Marshal(map[string]any{
		"type": "template",
		"data": map[string]any{
			"template_id": "ctp_AA1BZPqZhbwE",
			"template_variable": map[string]string{
				"settings":           "https://bot.gla-moodle.site/static/settings.html?id=" + chatId,
				"settings_string":    "https://bot.gla-moodle.site/static/settings.html?id=" + chatId,
				"settings_en":        "https://bot.gla-moodle.site/static/settings-en.html?id=" + chatId,
				"settings_en_string": "https://bot.gla-moodle.site/static/settings-en.html?id=" + chatId,
				"query":              "https://bot.gla-moodle.site/static/query.html?id=" + chatId,
			},
		},
	})
	if err != nil {
		return err
	}
	l.sendMsg(larkim.MsgTypeInteractive, chatId, string(msg))
	return nil
}

func (l *LarkSendService) SendWelcome(chatId string) error {
	msg, err := json.Marshal(map[string]any{
		"type": "template",
		"data": map[string]any{
			"template_id": "ctp_AA1Rsr2n0Qtv",
			"template_variable": map[string]string{
				"settings":    "https://bot.gla-moodle.site/static/settings.html?id=" + chatId,
				"settings_en": "https://bot.gla-moodle.site/static/settings-en.html?id=" + chatId,
			},
		},
	})
	if err != nil {
		return err
	}
	l.sendMsg(larkim.MsgTypeInteractive, chatId, string(msg))
	return nil
}
