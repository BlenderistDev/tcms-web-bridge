package tcms

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	tcms2 "tcms-web-bridge/pkg/tcms"
)

type Tcms interface {
	GetActions(ctx context.Context) (*tcms2.ActionList, error)
	GetConditions(ctx context.Context) (*tcms2.ConditionList, error)
	GetTriggers(ctx context.Context) (*tcms2.TriggerList, error)
	GetAutomations(ctx context.Context) (*tcms2.AutomationList, error)
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

func NewTcms(client tcms2.TcmsClient) Tcms {
	return tcms{
		client: client,
	}
}
