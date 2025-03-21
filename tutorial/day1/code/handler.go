package main

import (
	"context"
	echo "hello/kitex_gen/echo"
	"hello/biz/service"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *echo.Req) (resp *echo.Resp, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
