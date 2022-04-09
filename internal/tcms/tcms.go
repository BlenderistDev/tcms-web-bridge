package tcms

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"tcms-web-bridge/internal/config"
	tcms2 "tcms-web-bridge/pkg/tcms"
)

type Tcms interface {
	GetActions(ctx context.Context) (*tcms2.ActionList, error)
	GetConditions(ctx context.Context) (*tcms2.ConditionList, error)
	GetTriggers(ctx context.Context) (*tcms2.TriggerList, error)
	GetAutomations(ctx context.Context) (*tcms2.AutomationList, error)
	AddAutomation(ctx context.Context, automation *tcms2.Automation) error
	UpdateAutomation(ctx context.Context, request *tcms2.UpdateAutomationRequest) error
	RemoveAutomation(ctx context.Context, request *tcms2.RemoveAutomationRequest) error
}

type tcms struct {
	client tcms2.TcmsClient
}

func (t tcms) GetActions(ctx context.Context) (*tcms2.ActionList, error) {
	return t.client.GetActionList(ctx, &emptypb.Empty{})
}

func (t tcms) GetConditions(ctx context.Context) (*tcms2.ConditionList, error) {
	return t.client.GetConditionList(ctx, &emptypb.Empty{})
}

func (t tcms) GetTriggers(ctx context.Context) (*tcms2.TriggerList, error) {
	return t.client.GetTriggerList(ctx, &emptypb.Empty{})
}

func (t tcms) GetAutomations(ctx context.Context) (*tcms2.AutomationList, error) {
	return t.client.GetList(ctx, &emptypb.Empty{})
}

func (t tcms) AddAutomation(ctx context.Context, automation *tcms2.Automation) error {
	_, err := t.client.AddAutomation(ctx, automation)
	return err
}

func (t tcms) UpdateAutomation(ctx context.Context, request *tcms2.UpdateAutomationRequest) error {
	_, err := t.client.UpdateAutomation(ctx, request)
	return err
}

func (t tcms) RemoveAutomation(ctx context.Context, request *tcms2.RemoveAutomationRequest) error {
	_, err := t.client.RemoveAutomation(ctx, request)
	return err
}

func NewTcms(client tcms2.TcmsClient) Tcms {
	return tcms{
		client: client,
	}
}

func GetTcms(config config.Config) (Tcms, error) {
	conn, err := grpc.Dial(config.TcmsHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := tcms2.NewTcmsClient(conn)

	return NewTcms(client), nil
}
