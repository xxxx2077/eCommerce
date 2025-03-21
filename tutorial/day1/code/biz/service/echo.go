package service

import (
	"context"
	"fmt"
	echo "hello/kitex_gen/echo"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *echo.Req) (resp *echo.Resp, err error) {
	// Finish your business logic.

	value, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	if ok {
		fmt.Println("client_name = ", value)
	}
	return &echo.Resp{
		Resp: req.Mes,
	}, nil
}
