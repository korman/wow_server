package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wow_server/wowapi/internal/logic"
	"wow_server/wowapi/internal/svc"
	"wow_server/wowapi/internal/types"
)

func WowapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWowapiLogic(r.Context(), svcCtx)
		resp, err := l.Wowapi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
