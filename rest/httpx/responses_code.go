package httpx

import (
	"context"
	"net/http"

	"github.com/cloudernative/go-zero/core/errs"
)

// CommonResp Api通用返回结构体
type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CommonJson Api通用返回
func CommonJson(err error, w http.ResponseWriter, v any) {
	rsp := &CommonResp{
		Code:    errs.Code(err),
		Message: errs.Msg(err),
		Data:    v,
	}
	OkJson(w, rsp)
}

// CommonJsonCtx Api通用返回
func CommonJsonCtx(ctx context.Context, err error, w http.ResponseWriter, v any) {
	rsp := &CommonResp{
		Code:    errs.Code(err),
		Message: errs.Msg(err),
		Data:    v,
	}
	OkJsonCtx(ctx, w, rsp)
}
