package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-repository/user/api/internal/logic"
	"go-zero-repository/user/api/internal/svc"
	"go-zero-repository/user/api/internal/types"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		_, err := l.Api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, svcCtx.Config.TestConfig.Name)
		}
	}
}
