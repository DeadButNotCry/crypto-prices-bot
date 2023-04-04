package prices

import (
	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
	"github.com/redis/go-redis/v9"
)

type CryptoBotPriceService struct {
	cryptoBotClient *cryptobot.Client
	redisClient     *redis.Client
}

func NewCryptoBotPricrService(apiKey string, redisClient *redis.Client) *CryptoBotPriceService {
	return &CryptoBotPriceService{
		cryptoBotClient: cryptobot.NewClient(cryptobot.Options{
			APIToken: apiKey,
		}),
		redisClient: redisClient,
	}
}


