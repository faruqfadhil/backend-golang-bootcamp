package midtrans

import (
	"context"
	"fmt"
)

type Midtrans interface {
	Send(ctx context.Context, cr float64) error
}

type mt struct{}

func New() Midtrans {
	return &mt{}
}

func (u *mt) Send(ctx context.Context, cr float64) error {
	// Call to midtrans API.
	return fmt.Errorf("errr")
}
