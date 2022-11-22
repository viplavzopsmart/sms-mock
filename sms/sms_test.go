package sms

import (
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestHandler_SendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDoer := NewMockSMSSender(ctrl)
	t1 := New(mockDoer)
	testCases := []struct {
		mobile   string
		expOut   error
		desc     string
		mockCall *gomock.Call
	}{
		{
			mobile:   "987654a32",
			expOut:   errors.New("invalid phone"),
			desc:     "Not valid phone number",
			mockCall: nil,
		},
		{
			mobile:   "9876543210",
			expOut:   errors.New("invalid sms message"),
			desc:     "Not valid sms number dfj jdfhuf igbuhfbg dhfgyudbg dhbfyub dhfbuhdf",
			mockCall: nil,
		},
		{
			mobile:   "5959393944",
			expOut:   nil,
			desc:     "twilio cannot send sms",
			mockCall: mockDoer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil),
		},
	}

	for _, tc := range testCases {
		receivedError := t1.SendMessage(tc.mobile, tc.desc)
		if !reflect.DeepEqual(receivedError, tc.expOut) {
			t.Errorf("expecting %v getting %v", receivedError, tc.expOut)
		}

	}

}
