package dao

import (
	"time"

	"powerbot/util"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ChatId      string    // 飞书 chat_id
	RoomNumber  string    // 房间号
	RoomCode    string    // 加密的房间号（为什么要加密？）
	Hour        int64     // 选择推送的小时
	LastPower   int64     // 单位：瓦时
	LastBalance int64     // 单位：分
	LastQuery   time.Time // 上次查询时间
}

func (d *DB) NewGroup(ChatId string) error {
	d.db.Create(&Group{ChatId: ChatId, Hour: -1})
	return nil
}

func (d *DB) FindGroupByChatId(chatId string) (*Group, error) {
	group := Group{}
	d.db.Where("chat_id = ?", chatId).First(&group)
	return &group, nil
}

func (d *DB) FindGroups(hour int64) ([]Group, error) {
	groups := []Group{}
	d.db.Find(&groups)
	return groups, nil
}

func (d *DB) FindGroupsByHour(hour int64) ([]Group, error) {
	groups := []Group{}
	d.db.Where("hour = ?", hour).Find(&groups)
	return groups, nil
}

func (d *DB) SetGroupSettingsByChatId(chatId string, roomNumber string, hour int64) error {
	group := Group{}
	d.db.Where("chat_id = ?", chatId).First(&group)
	group.RoomNumber = roomNumber
	group.RoomCode = util.RoomCode(roomNumber)
	group.Hour = hour
	d.db.Save(&group)
	return nil
}

func (d *DB) SetGroupQueryResultByChatId(chatId string, power int64, balance int64) error {
	group := Group{}
	d.db.Where("chat_id = ?", chatId).First(&group)
	group.LastBalance = balance
	group.LastPower = power
	group.LastQuery = time.Now()
	d.db.Save(&group)
	return nil
}

func (d *DB) DeleteGroupById(chatId string) error {
	d.db.Where("chat_id = ?", chatId).Delete(&Group{})
	return nil
}
