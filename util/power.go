package util

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	Timestamp int64
	HttpCode  int64
	Msg       string
	Data      ResponseData
}

type ResponseData struct {
	Retcode  string
	AreaId   string
	Build    string
	RoomId   string
	RoomName string
	Sydl     string // 电量 千瓦时
	Syje     string // 金额 元
	Msg      string
}

const SCHOOL_API = "https://wx.uestc.edu.cn/power/oneCartoon/list"

func PowerLookup(roomCode string) (balance int64, power int64, err error) {
	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").SetBody("roomCode=" + url.QueryEscape(roomCode)).Post(SCHOOL_API)
	if err != nil {
		return 0, 0, err
	}

	r := Response{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return 0, 0, err
	}

	balanceFloat, err := strconv.ParseFloat(r.Data.Syje, 64)
	if err != nil {
		return 0, 0, err
	}
	powerFloat, err := strconv.ParseFloat(r.Data.Sydl, 64)
	if err != nil {
		return 0, 0, err
	}

	balance = int64(balanceFloat * 100) // 元 -> 分
	power = int64(powerFloat * 1000)    // 千瓦时 -> 瓦时

	return balance, power, nil
}
