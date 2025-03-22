package home

import (
	"context"

	"eCommerce/app/frontend/biz/service"
	"eCommerce/app/frontend/biz/utils"
	home "eCommerce/app/frontend/hertz_gen/frontend/home"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 返回html
	// resp为响应参数
	c.HTML(consts.StatusOK, "home.tmpl", resp)
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
