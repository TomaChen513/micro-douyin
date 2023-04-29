package video

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"miniTikTok/apis/internal/logic/video"
	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"
)

func CommentVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentVideoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewCommentVideoLogic(r.Context(), svcCtx)
		resp, err := l.CommentVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
