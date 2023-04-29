package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"miniTikTok/apis/internal/logic/user"
	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"
)

func FansListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FansListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFansListLogic(r.Context(), svcCtx)
		resp, err := l.FansList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
