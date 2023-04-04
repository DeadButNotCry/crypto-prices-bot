package prices

import (
	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
	"github.com/redis/go-redis/v9"
)

type CryptoBotPriceService struct {
	client      cryptobot.Client
	redisClient redis.Client
}

func CryptoBot(apiKey string)
