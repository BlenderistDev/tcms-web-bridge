package telegramClient

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	mock_telegram "tcms-web-bridge/internal/testing/telegram"
	"tcms-web-bridge/pkg/telegram"
)

func TestTelegramClient_Authorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const phone = "999999"
	request := &telegram.LoginMessage{Phone: phone}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Login(gomock.Eq(ctx), gomock.Eq(request))

	tg := newTelegram(client)
	err := tg.Authorization(ctx, phone)
	assert.Nil(t, err)
}

func TestTelegramClient_AuthSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const code = "999999"
	request := &telegram.SignMessage{Code: code}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Sign(gomock.Eq(ctx), gomock.Eq(request))

	tg := newTelegram(client)
	err := tg.AuthSignIn(ctx, code)
	assert.Nil(t, err)
}

func TestTelegramClient_GetCurrentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	request := &telegram.GetUserRequest{Peer: "me"}
	user := &telegram.User{}
	response := &telegram.UserResponse{User: user}

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().GetUser(gomock.Eq(ctx), gomock.Eq(request)).Return(response, nil)

	tg := newTelegram(client)
	res, err := tg.GetCurrentUser(ctx)
	assert.Nil(t, err)
	assert.Equal(t, user, res)
}

func TestTelegramClient_Dialogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	request := &emptypb.Empty{}
	response := &telegram.DialogsResponse{}

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().GetDialogs(gomock.Eq(ctx), gomock.Eq(request)).Return(response, nil)

	tg := newTelegram(client)
	res, err := tg.Dialogs(ctx)
	assert.Nil(t, err)
	assert.Equal(t, response, res)
}

func TestTelegramClient_SendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		peer    = "peer"
		message = "message"
	)
	request := &telegram.SendMessageRequest{
		Peer:    peer,
		Message: message,
	}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Send(gomock.Eq(ctx), gomock.Eq(request))

	tg := newTelegram(client)
	err := tg.SendMessage(ctx, peer, message)
	assert.Nil(t, err)
}

func TestGetTelegram(t *testing.T) {
	conn, err := grpc.Dial("host", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err)
	_ = GetTelegram(conn)
}
