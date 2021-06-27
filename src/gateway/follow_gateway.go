package gateway

import (
	"context"
	"github.com/go-resty/resty/v2"
	"os"
	"user-service/dto"
)

func BanFollow(ctx context.Context, userId string) error {
	client := resty.New()
	userDto := dto.BanProfileDTO{ProfileId: userId}

	domain := os.Getenv("FOLLOW_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		_, _ = client.R().
			SetBody(userDto).
			EnableTrace().
			Post("https://" + domain + ":8089/banUser")

		return nil
	} else {
		_, _ = client.R().
			SetBody(userDto).
			EnableTrace().
			Post("http://" + domain + ":8089/banUser")

		return nil
	}



}

