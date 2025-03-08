package pack

import (
	"TikTok/biz/model/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Base Base `json:"base"`
}

// 数据类型多样-用interface
func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}
func BuildBaseResp(err error) *model.BaseResp {
	if err == nil {
		return &model.BaseResp{
			Code: 10000,
			Msg:  "success",
		}
	}
	return &model.BaseResp{
		Code: 10001,
		Msg:  err.Error(),
	}
}

func SendFailResponse(c *app.RequestContext, err error) {
	SendResponse(c, BuildBaseResp(err))
}

// SendFailResponse 发送错误响应
func SendFailResponse_UsedInJWT(c *app.RequestContext, code int, message string) {
	c.JSON(code, map[string]interface{}{
		"code": code,
		"msg":  message,
	})
}
