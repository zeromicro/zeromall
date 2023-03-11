package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/authn/api/internal/logic/auth"
	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
