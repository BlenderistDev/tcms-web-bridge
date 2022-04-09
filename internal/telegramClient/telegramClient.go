package telegramClient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"tcms-web-bridge/internal/config"
	"tcms-web-bridge/pkg/telegram"
)

type TelegramClient interface {
	Authorization(phone string) error
	AuthSignIn(code string) error
	GetCurrentUser() (*telegram.User, error)
	Dialogs() (*telegram.DialogsResponse, error)
	SendMessage(peer, message string) error
	MuteUser(id, accessHash string, unMute bool) error
	MuteChat(id string, unMute bool) error
}

type telegramClient struct {
	telegram telegram.TelegramClient
}

// NewTelegram create new telegram client
func NewTelegram(config config.Config) (TelegramClient, error) {
	conn, err := grpc.Dial(config.TelegramBridgeHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	tg := telegram.NewTelegramClient(conn)

	return &telegramClient{telegram: tg}, nil
}

// Authorization request for authorization in telegram client
func (t *telegramClient) Authorization(phone string) error {
	_, err := t.telegram.Login(context.Background(), &telegram.LoginMessage{Phone: phone})
	return err
}

// AuthSignIn request for sign in telegram client with auth code
func (t *telegramClient) AuthSignIn(code string) error {
	_, err := t.telegram.Sign(context.Background(), &telegram.SignMessage{Code: code})

	return err
}

// GetCurrentUser return current telegram user
func (t *telegramClient) GetCurrentUser() (*telegram.User, error) {
	request := telegram.GetUserRequest{Peer: "me"}
	user, err := t.telegram.GetUser(context.Background(), &request)
	return user.GetUser(), err
}

// Dialogs return telegram dialogs
func (t *telegramClient) Dialogs() (*telegram.DialogsResponse, error) {
	dialogs, err := t.telegram.GetDialogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return dialogs, nil
}

// SendMessage send message throw telegram
func (t *telegramClient) SendMessage(peer, message string) error {
	request := telegram.SendMessageRequest{
		Peer:    peer,
		Message: message,
	}

	_, err := t.telegram.Send(context.Background(), &request)

	return err
}

// MuteUser mute or unmute telegram user
func (t telegramClient) MuteUser(id, accessHash string, unMute bool) error {
	request := telegram.MuteUserRequest{
		Id:         id,
		AccessHash: accessHash,
		Unmute:     unMute,
	}
	res, err := t.telegram.MuteUser(context.Background(), &request)

	if err != nil {
		return err
	}

	if !res.GetSuccess() {
		return fmt.Errorf("error while setting user notify settings")
	}

	return nil
}

// MuteChat mute or unmute telegram chat
func (t telegramClient) MuteChat(id string, unMute bool) error {
	request := telegram.MuteChatRequest{
		Id:     id,
		Unmute: unMute,
	}
	res, err := t.telegram.MuteChat(context.Background(), &request)

	if err != nil {
		return err
	}

	if !res.GetSuccess() {
		return fmt.Errorf("error while setting chat notify settings")
	}

	return nil
}
