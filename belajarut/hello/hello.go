package hello

import (
	"belajarut/midtrans"
	"context"
	"fmt"
	"strings"
)

type Hw struct {
	Midtrans midtrans.Midtrans
}

func New(mt midtrans.Midtrans) *Hw {
	return &Hw{
		Midtrans: mt,
	}
}

func HelloWorld(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func Hay(name string) string {
	return fmt.Sprintf("%s hai", name)
}

type HayComplexResp struct {
	Name    string
	Address string
	Phone   string
}

func (u *Hw) HayComplex(name string, amount float64) HayComplexResp {
	if strings.Contains(name, "faruq") {
		return HayComplexResp{
			Name:    fmt.Sprintf("hello %s", name),
			Address: "sby",
			Phone:   "1234",
		}
	}

	err := u.Midtrans.Send(context.Background(), amount)
	if err != nil {
		return HayComplexResp{
			Name:    name,
			Address: "sby error",
			Phone:   "1234",
		}
	}

	return HayComplexResp{
		Name:    name,
		Address: "sby",
		Phone:   "1234",
	}
}
