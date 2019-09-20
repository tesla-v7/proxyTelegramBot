package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/Syfaro/telegram-bot-api"
)

var configName string = "config.json"

func getMesage(bot tgbotapi.BotAPI, chatId int64) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var msg string = r.URL.Query().Get("msg")
		if msg == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No message"))
			return
		}
		msgTelegram := tgbotapi.NewMessage(chatId, msg)
		bot.Send(msgTelegram)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	}
}

type botConfig struct {
	Token    string
	ChatId   int64
	HttpPort string
}

func getConfigBot() (botConfig, error) {
	var zero botConfig
	dataStr, err := ioutil.ReadFile(configName)
	if err != nil {
		return zero, err
	}
	var config botConfig
	errParse := json.Unmarshal(dataStr, &config)
	if errParse != nil {
		return zero, errParse
	}
	return config, nil
}

func main() {
	config, err := getConfigBot()
	if err != nil {
		panic(fmt.Sprintf("Error get config. %s", err.Error()))
	}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		panic(fmt.Sprintf("Not create bot. %s", err.Error()))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<body>Hi!<br> Goto <a href='http://localhost:%s/msg?msg=test'>http://localhost:%s/msg?msg=test</a></body>", config.HttpPort, config.HttpPort)
	})
	http.HandleFunc("/msg", getMesage(*bot, config.ChatId))

	fmt.Printf("Start server port %s\n", config.HttpPort)
	http.ListenAndServe(":"+config.HttpPort, nil)
}
