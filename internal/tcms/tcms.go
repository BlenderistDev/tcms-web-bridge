package tcms

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	tcms2 "tcms-web-bridge/pkg/tcms"
)

type Tcms interface {
	getActions(ctx context.Context) (*tcms2.ActionList, error)
	getConditions(ctx context.Context) (*tcms2.ConditionList, error)
	getTriggers(ctx context.Context) (*tcms2.TriggerList, error)
	getAutomations(ctx context.Context) (*tcms2.AutomationList, error)
}

type tcms struct {
	client tcms2.TcmsClient
}

func (t tcms) getActions(ctx context.Context) (*tcms2.ActionList, error) {
	return t.client.GetActionList(ctx, &emptypb.Empty{})
}

func (t tcms) getConditions(ctx context.Context) (*tcms2.ConditionList, error) {
	return t.client.GetConditionList(ctx, &emptypb.Empty{})
}

func (t tcms) getTriggers(ctx context.Context) (*tcms2.TriggerList, error) {
	return t.client.GetTriggerList(ctx, &emptypb.Empty{})
}

func (t tcms) getAutomations(ctx context.Context) (*tcms2.AutomationList, error) {
	return t.client.GetList(ctx, &emptypb.Empty{})
}

func NewTcms(client tcms2.TcmsClient) Tcms {
	return tcms{
		client: client,
	}
}
