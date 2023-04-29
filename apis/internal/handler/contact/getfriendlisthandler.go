package contact

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"miniTikTok/apis/internal/logic/contact"
	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"
)

func GetFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFriendListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := contact.NewGetFriendListLogic(r.Context(), svcCtx)
		resp, err := l.GetFriendList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
