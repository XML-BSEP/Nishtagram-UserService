package saga

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"user-service/dto"
	"user-service/usecase"
)

type authSaga struct {
	ProfileInfoUsecase usecase.ProfileInfoUseCase
	redisClient *redis.Client
}


type AuthSaga interface {
	SagaAuth(context context.Context)
}

func NewAuthSaga(profileInfoUsecase usecase.ProfileInfoUseCase, redisClient *redis.Client) AuthSaga {
	return &authSaga{
		ProfileInfoUsecase: profileInfoUsecase,
		redisClient: redisClient,
	}
}

func (a *authSaga) SagaAuth(context context.Context) {
	pubsub := a.redisClient.Subscribe(context, UserChannel, ReplyChannel)
	if _, err := pubsub.Receive(context); err != nil {
		log.Fatalf("error subscribing %s", err)
	}

	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	for {
		select {
		case msg := <- ch:
			m := Message{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case UserChannel:
				if m.Action == ActionStart {



					user := m.Payload

					exists, err := a.ProfileInfoUsecase.Exists(user.Username, user.Email, context)

					if exists {
						sendToReplyChannel(context, a.redisClient, m, ActionRollback, AuthService, UserService)
						continue
					}
					newUser := dto.NewUserDTOtoEntity(user)
					newUser.Profile.Type = 1
					mediaToAttach, err := a.ProfileInfoUsecase.EncodeBase64(newUser.ProfileImage, newUser.ID, context)
					if err != nil {
						sendToReplyChannel(context, a.redisClient, m, ActionError, AuthService, UserService)
						continue
					}
					newUser.ProfileImage = mediaToAttach
					if err := a.ProfileInfoUsecase.SaveNewUser(newUser, context); err != nil {
						sendToReplyChannel(context, a.redisClient, m, ActionError, AuthService, UserService)
						continue
					}
					sendToReplyChannel(context, a.redisClient, m, ActionDone, AuthService, UserService)
				}
			}
		}
	}
}

func sendToReplyChannel(context context.Context, client *redis.Client, m Message, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	binaryMessage, _ := MarshalBinary(&m)
	if err = client.Publish(context, ReplyChannel, binaryMessage).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", ReplyChannel)
	}
	log.Printf("%s message published to channel :%s", m.Action, ReplyChannel)
}



