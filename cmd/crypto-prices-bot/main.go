package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/deadbutnotcry/crypto-prices-bot/internal/telegram"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("BOT_TOKEN")
	ids := getAdminIds()
	fmt.Printf("Bot starting token: %s admins: %v\n", token, ids)
	bot, _ := telegram.NewTgBot(token, ids)
	err = bot.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func getAdminIds() []int {
	stringIds := strings.Split(os.Getenv("ADMIN_IDS"), ",")
	ids := make([]int, len(stringIds))

	for i, v := range stringIds {
		id, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Incorrect Id list format")
		}
		ids[i] = id
	}
	return ids
}
