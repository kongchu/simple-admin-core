package user

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /user/status user UpdateUserStatus
//
// Set user's status | 更新用户状态
//
// Set user's status | 更新用户状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StatusCodeReq
//
// Responses:
//  200: BaseMsgResp

func UpdateUserStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUpdateUserStatusLogic(r, svcCtx)
		resp, err := l.UpdateUserStatus(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
