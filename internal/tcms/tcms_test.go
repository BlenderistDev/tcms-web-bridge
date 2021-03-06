package tcms

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"tcms-web-bridge/internal/dry"
	mock_tcms "tcms-web-bridge/internal/testing/tcms"
	tcms2 "tcms-web-bridge/pkg/tcms"
)

func TestTcms_GetActions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.ActionList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetActionList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := newTcms(c)
	list, err := tcms.GetActions(ctx)
	dry.TestCheckEqual(t, expected, list)
	assert.Nil(t, err)
}

func TestTcms_GetConditions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.ConditionList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetConditionList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := newTcms(c)
	list, err := tcms.GetConditions(ctx)
	dry.TestCheckEqual(t, expected, list)
	assert.Nil(t, err)
}

func TestTcms_GetTriggers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.TriggerList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetTriggerList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := newTcms(c)
	list, err := tcms.GetTriggers(ctx)
	dry.TestCheckEqual(t, expected, list)
	assert.Nil(t, err)
}

func TestTcms_GetAutomations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.AutomationList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := newTcms(c)
	list, err := tcms.GetAutomations(ctx)
	dry.TestCheckEqual(t, expected, list)
	assert.Nil(t, err)
}

func TestTcms_AddAutomation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	automation := &tcms2.Automation{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().AddAutomation(gomock.Eq(ctx), gomock.Eq(automation))
	tcms := newTcms(c)
	err := tcms.AddAutomation(ctx, automation)
	assert.Nil(t, err)
}

func TestTcms_UpdateAutomation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	request := &tcms2.UpdateAutomationRequest{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().UpdateAutomation(gomock.Eq(ctx), gomock.Eq(request))
	tcms := newTcms(c)
	err := tcms.UpdateAutomation(ctx, request)
	assert.Nil(t, err)
}

func TestTcms_RemoveAutomation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	request := &tcms2.RemoveAutomationRequest{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().RemoveAutomation(gomock.Eq(ctx), gomock.Eq(request))
	tcms := newTcms(c)
	err := tcms.RemoveAutomation(ctx, request)
	assert.Nil(t, err)
}

func TestTcms_GetAutomation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	request := &tcms2.AutomationRequest{}

	automation := &tcms2.Automation{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetOne(gomock.Eq(ctx), gomock.Eq(request)).Return(automation, nil)
	tcms := newTcms(c)
	res, err := tcms.GetAutomation(ctx, request)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, automation, res)
}

func TestGetTcms(t *testing.T) {
	conn := &grpc.ClientConn{}
	GetTcms(conn)
}
