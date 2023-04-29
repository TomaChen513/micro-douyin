package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"miniTikTok/apis/internal/logic/user"
	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"
)

func FollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFollowLogic(r.Context(), svcCtx)
		resp, err := l.Follow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
