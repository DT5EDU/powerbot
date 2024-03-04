package core

import (
	"fmt"

	"powerbot/dao"
	"powerbot/larkSend"
	"powerbot/util"
)

type CoreService struct {
	db          *dao.DB
	sendService *larkSend.LarkSendService
}

func NewCoreService(db *dao.DB, sendService *larkSend.LarkSendService) *CoreService {
	return &CoreService{
		db:          db,
		sendService: sendService,
	}
}

type SendWelcomeRequest struct {
	ChatId string
}

type QueryPowerRequest struct {
	ChatId string
}

type QueryPowerFirstTimeRequest struct {
	ChatId string
}

type SendMenuRequest struct {
	ChatId string
}

type GoodbyeRequest struct {
	ChatId string
}

type TimedQueryRequest struct {
	Hour int64
}

type SettingsRequest struct {
	ChatId     string
	RoomNumber string
	Hour       int64
}

func (c *CoreService) SendWelcome(r *SendWelcomeRequest) error {
	err := c.db.NewGroup(r.ChatId)
	if err != nil {
		return err
	}
	return c.sendService.SendWelcome(r.ChatId)
}

func (c *CoreService) QueryPower(r *QueryPowerRequest) error {
	group, err := c.db.FindGroupByChatId(r.ChatId)
	if err != nil {
		return err
	}

	newBalance, newPower, err := util.PowerLookup(group.RoomCode)
	if err != nil {
		return err
	}

	c.sendService.SendPower(group.ChatId, group.RoomNumber, newBalance, newPower, group.LastQuery, group.LastBalance, group.LastPower)
	if err != nil {
		return err
	}

	err = c.db.SetGroupQueryResultByChatId(group.ChatId, newPower, newBalance)
	return err
}

func (c *CoreService) QueryPowerFirstTime(r *QueryPowerFirstTimeRequest) error {
	group, err := c.db.FindGroupByChatId(r.ChatId)
	if err != nil {
		return err
	}

	newBalance, newPower, err := util.PowerLookup(group.RoomCode)
	if err != nil {
		return err
	}

	c.sendService.SendPowerFirstTime(group.ChatId, group.RoomNumber, newBalance, newPower)
	if err != nil {
		return err
	}

	err = c.db.SetGroupQueryResultByChatId(group.ChatId, newPower, newBalance)
	return err
}

func (c *CoreService) SendMenu(r *SendMenuRequest) error {
	return c.sendService.SendMenu(r.ChatId)
}

func (c *CoreService) Goodbye(r *GoodbyeRequest) error {
	return c.db.DeleteGroupById(r.ChatId)
}

func (c *CoreService) TimedQuery(r *TimedQueryRequest) error {
	groups, err := c.db.FindGroupsByHour(r.Hour)
	if err != nil {
		return err
	}

	for _, group := range groups {
		newBalance, newPower, err := util.PowerLookup(group.RoomCode)
		if err != nil {
			fmt.Println(err)
			continue
		}

		c.sendService.SendPower(group.ChatId, group.RoomNumber, newBalance, newPower, group.LastQuery, group.LastBalance, group.LastPower)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = c.db.SetGroupQueryResultByChatId(group.ChatId, newPower, newBalance)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func (c *CoreService) Settings(r *SettingsRequest) error {
	err := c.db.SetGroupSettingsByChatId(r.ChatId, r.RoomNumber, r.Hour)
	if err != nil {
		return err
	}
	return c.QueryPowerFirstTime(&QueryPowerFirstTimeRequest{ChatId: r.ChatId})
}
