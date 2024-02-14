package hello

import (
	mock_midtrans "belajarut/mock/midtrans"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("menyiapkan test dep")
	m.Run()
	fmt.Println("test selesai")
}

func TestHelloWorld(t *testing.T) {
	resp := HelloWorld("faruq")
	expect := "hello faruq"

	if resp != expect {
		t.Errorf("diff resp, got %s want %s", resp, expect)
	}
}

func TestHayComplex(t *testing.T) {
	type arg struct {
		name   string
		amount float64
	}
	type expectation struct {
		want HayComplexResp
	}
	ctrl := gomock.NewController(t)
	mockMidtrans := mock_midtrans.NewMockMidtrans(ctrl)
	hw := New(mockMidtrans)
	tests := []struct {
		name        string
		arg         arg
		expectation expectation
		mockFunc    func()
	}{
		{
			name: "name contains faruq",
			arg: arg{
				name:   "faruq",
				amount: 10,
			},
			expectation: expectation{
				want: HayComplexResp{
					Name:    "hello faruq",
					Address: "sby",
					Phone:   "1234",
				},
			},
			mockFunc: func() {},
		},
		{
			name: "call midtrans but got an error",
			arg: arg{
				name:   "agus",
				amount: 99,
			},
			mockFunc: func() {
				mockMidtrans.EXPECT().Send(context.Background(), float64(99)).Return(fmt.Errorf("err"))
			},
			expectation: expectation{
				want: HayComplexResp{
					Name:    "agus",
					Address: "sby error",
					Phone:   "1234",
				},
			},
		},
		{
			name: "success call midtrans",
			arg: arg{
				name:   "agus",
				amount: 100,
			},
			mockFunc: func() {
				mockMidtrans.EXPECT().Send(context.Background(), float64(100)).Return(nil)
			},
			expectation: expectation{
				want: HayComplexResp{
					Name:    "agus",
					Address: "sby",
					Phone:   "1234",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			resp := hw.HayComplex(test.arg.name, test.arg.amount)
			assert.Equal(t, test.expectation.want, resp)
		})
	}

}
func TestHay(t *testing.T) {
	resp := Hay("faruq")
	expect := "faruq hai"

	if resp != expect {
		t.Errorf("diff resp, got %s want %s", resp, expect)
	}
}
