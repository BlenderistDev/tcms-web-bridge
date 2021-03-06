package telegramClient

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"tcms-web-bridge/pkg/telegram"
)

type TelegramClient interface {
	Authorization(ctx context.Context, phone string) error
	AuthSignIn(ctx context.Context, code string) error
	GetCurrentUser(ctx context.Context) (*telegram.User, error)
	Dialogs(ctx context.Context) (*telegram.DialogsResponse, error)
	SendMessage(ctx context.Context, peer, message string) error
}

type telegramClient struct {
	telegram telegram.TelegramClient
}

func newTelegram(tg telegram.TelegramClient) TelegramClient {
	return &telegramClient{telegram: tg}
}

// GetTelegram create new telegram client
func GetTelegram(conn *grpc.ClientConn) TelegramClient {
	tg := telegram.NewTelegramClient(conn)
	return newTelegram(tg)
}

// Authorization request for authorization in telegram client
func (t *telegramClient) Authorization(ctx context.Context, phone string) error {
	_, err := t.telegram.Login(ctx, &telegram.LoginMessage{Phone: phone})
	return err
}

// AuthSignIn request for sign in telegram client with auth code
func (t *telegramClient) AuthSignIn(ctx context.Context, code string) error {
	_, err := t.telegram.Sign(ctx, &telegram.SignMessage{Code: code})
	return err
}

// GetCurrentUser return current telegram user
func (t *telegramClient) GetCurrentUser(ctx context.Context) (*telegram.User, error) {
	request := telegram.GetUserRequest{Peer: "me"}
	user, err := t.telegram.GetUser(ctx, &request)
	return user.GetUser(), err
}

// Dialogs return telegram dialogs
func (t *telegramClient) Dialogs(ctx context.Context) (*telegram.DialogsResponse, error) {
	return t.telegram.GetDialogs(ctx, &emptypb.Empty{})
}

// SendMessage send message throw telegram
func (t *telegramClient) SendMessage(ctx context.Context, peer, message string) error {
	request := &telegram.SendMessageRequest{
		Peer:    peer,
		Message: message,
	}

	_, err := t.telegram.Send(ctx, request)

	return err
}
