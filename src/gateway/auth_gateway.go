package gateway

import (
	"context"
	"github.com/go-resty/resty/v2"
	"os"
	"user-service/dto"
)

func DeleteProfileInfo(ctx context.Context, username string) error {
	client := resty.New()
	userDto := dto.DeleteProfileInfoDTO{Username: username}

	domain := os.Getenv("AUTH_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}

	if os.Getenv("DOCKER_ENV") == "" {
		_, _ = client.R().
			SetBody(userDto).
			EnableTrace().
			Post("https://" + domain + ":8091/deleteProfileInfo")

		return nil
	} else {
		_, _ = client.R().
			SetBody(userDto).
			EnableTrace().
			Post("http://" + domain + ":8091/deleteProfileInfo")

		return nil
	}



}


