package larkListen

import (
	"context"
	"fmt"

	"powerbot/core"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	dispatcher "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type LarkListenService struct {
	client       *lark.Client
	core         *core.CoreService
	EventHandler func(*gin.Context)
	CardHandler  func(*gin.Context)
}

func NewLarkListenService(appId string, appSecret string, verificationToken string, eventKey string, core *core.CoreService) *LarkListenService {
	service := &LarkListenService{}

	service.client = lark.NewClient(appId, appSecret)
	service.core = core
	service.EventHandler = sdkginext.NewEventHandlerFunc(
		dispatcher.NewEventDispatcher(verificationToken, eventKey).
			OnP2MessageReceiveV1(service.OnMessageReceive).         // 接收消息
			OnP2ChatMemberBotAddedV1(service.OnMemberBotAdded).     // 加入群
			OnP2ChatMemberBotDeletedV1(service.OnMemberBotDeleted). // 退出群
			OnP2ChatDisbandedV1(service.OnChatDisbanded),           // 群解散
	)
	service.CardHandler = sdkginext.NewCardActionHandlerFunc(
		larkcard.NewCardActionHandler(
			verificationToken, eventKey, func(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
				fmt.Println(larkcore.Prettify(cardAction))
				fmt.Println(cardAction.RequestId())
				return nil, nil
			},
		),
	)
	return service
}

func (l *LarkListenService) OnMessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	return l.core.SendMenu(&core.SendMenuRequest{ChatId: *event.Event.Message.ChatId})
}

func (l *LarkListenService) OnMemberBotAdded(ctx context.Context, event *larkim.P2ChatMemberBotAddedV1) error {
	err := l.core.SendWelcome(&core.SendWelcomeRequest{ChatId: *event.Event.ChatId})
	return err
}

func (l *LarkListenService) OnMemberBotDeleted(ctx context.Context, event *larkim.P2ChatMemberBotDeletedV1) error {
	err := l.core.Goodbye(&core.GoodbyeRequest{ChatId: *event.Event.ChatId})
	return err
}

func (l *LarkListenService) OnChatDisbanded(ctx context.Context, event *larkim.P2ChatDisbandedV1) error {
	err := l.core.Goodbye(&core.GoodbyeRequest{ChatId: *event.Event.ChatId})
	return err
}
