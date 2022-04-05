package tcms

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
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
	tcms := NewTcms(c)
	list, err := tcms.GetActions(ctx)
	dry.TestCheckEqual(t, expected, list)
	dry.TestHandleError(t, err)
}

func TestTcms_GetConditions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.ConditionList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetConditionList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := NewTcms(c)
	list, err := tcms.GetConditions(ctx)
	dry.TestCheckEqual(t, expected, list)
	dry.TestHandleError(t, err)
}

func TestTcms_GetTriggers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.TriggerList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetTriggerList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := NewTcms(c)
	list, err := tcms.GetTriggers(ctx)
	dry.TestCheckEqual(t, expected, list)
	dry.TestHandleError(t, err)
}

func TestTcms_GetAutomations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expected := &tcms2.AutomationList{}

	c := mock_tcms.NewMockTcmsClient(ctrl)
	c.EXPECT().GetList(gomock.Eq(ctx), gomock.Eq(&emptypb.Empty{})).Return(expected, nil)
	tcms := NewTcms(c)
	list, err := tcms.GetAutomations(ctx)
	dry.TestCheckEqual(t, expected, list)
	dry.TestHandleError(t, err)
}