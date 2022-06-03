package lib

import (
	"fmt"
	"github.com/Sakurasan/MiraiGo-Template/bot"
	"github.com/Sakurasan/MiraiGo-Template/config"
	_ "golang.org/x/mobile/bind"
)

func init() {
	fmt.Println("bot as lib")
}

func InitBot(configJSONContent string, deviceJSONContent string) {
	config.InitWithContent([]byte(configJSONContent))
	bot.InitWithDeviceJSONContent([]byte(deviceJSONContent))
	bot.StartService()
	bot.UseProtocol(bot.IPad)
	bot.Login()
	bot.RefreshList()
}
