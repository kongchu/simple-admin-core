// Code generated by goctl. DO NOT EDIT!
// Source: core.proto

package server

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/logic"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type CoreServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedCoreServer
}

func NewCoreServer(svcCtx *svc.ServiceContext) *CoreServer {
	return &CoreServer{
		svcCtx: svcCtx,
	}
}

// init
func (s *CoreServer) InitDatabase(ctx context.Context, in *core.Empty) (*core.BaseResp, error) {
	l := logic.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// user service
func (s *CoreServer) Login(ctx context.Context, in *core.LoginReq) (*core.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *CoreServer) ChangePassword(ctx context.Context, in *core.ChangePasswordReq) (*core.BaseResp, error) {
	l := logic.NewChangePasswordLogic(ctx, s.svcCtx)
	return l.ChangePassword(in)
}

func (s *CoreServer) CreateOrUpdateUser(ctx context.Context, in *core.CreateOrUpdateUserReq) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateUserLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateUser(in)
}

func (s *CoreServer) GetUserById(ctx context.Context, in *core.UUIDReq) (*core.UserInfoResp, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *CoreServer) GetUserList(ctx context.Context, in *core.GetUserListReq) (*core.UserListResp, error) {
	l := logic.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}

func (s *CoreServer) DeleteUser(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *CoreServer) UpdateProfile(ctx context.Context, in *core.UpdateProfileReq) (*core.BaseResp, error) {
	l := logic.NewUpdateProfileLogic(ctx, s.svcCtx)
	return l.UpdateProfile(in)
}

// menu service
func (s *CoreServer) CreateOrUpdateMenu(ctx context.Context, in *core.CreateOrUpdateMenuReq) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateMenuLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenu(in)
}

func (s *CoreServer) DeleteMenu(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *CoreServer) GetMenuListByRole(ctx context.Context, in *core.IDReq) (*core.MenuInfoList, error) {
	l := logic.NewGetMenuListByRoleLogic(ctx, s.svcCtx)
	return l.GetMenuListByRole(in)
}

func (s *CoreServer) GetMenuByPage(ctx context.Context, in *core.PageInfoReq) (*core.MenuInfoList, error) {
	l := logic.NewGetMenuByPageLogic(ctx, s.svcCtx)
	return l.GetMenuByPage(in)
}

func (s *CoreServer) CreateOrUpdateMenuParam(ctx context.Context, in *core.CreateOrUpdateMenuParamReq) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateMenuParamLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenuParam(in)
}

func (s *CoreServer) DeleteMenuParam(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteMenuParamLogic(ctx, s.svcCtx)
	return l.DeleteMenuParam(in)
}

func (s *CoreServer) GeMenuParamListByMenuId(ctx context.Context, in *core.IDReq) (*core.MenuParamListResp, error) {
	l := logic.NewGeMenuParamListByMenuIdLogic(ctx, s.svcCtx)
	return l.GeMenuParamListByMenuId(in)
}

// role service
func (s *CoreServer) CreateOrUpdateRole(ctx context.Context, in *core.RoleInfo) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateRoleLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateRole(in)
}

func (s *CoreServer) DeleteRole(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

func (s *CoreServer) GetRoleById(ctx context.Context, in *core.IDReq) (*core.RoleInfo, error) {
	l := logic.NewGetRoleByIdLogic(ctx, s.svcCtx)
	return l.GetRoleById(in)
}

func (s *CoreServer) GetRoleList(ctx context.Context, in *core.PageInfoReq) (*core.RoleListResp, error) {
	l := logic.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}

func (s *CoreServer) SetRoleStatus(ctx context.Context, in *core.SetStatusReq) (*core.BaseResp, error) {
	l := logic.NewSetRoleStatusLogic(ctx, s.svcCtx)
	return l.SetRoleStatus(in)
}

// api management service
func (s *CoreServer) CreateOrUpdateApi(ctx context.Context, in *core.ApiInfo) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateApiLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateApi(in)
}

func (s *CoreServer) DeleteApi(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteApiLogic(ctx, s.svcCtx)
	return l.DeleteApi(in)
}

func (s *CoreServer) GetApiList(ctx context.Context, in *core.ApiPageReq) (*core.ApiListResp, error) {
	l := logic.NewGetApiListLogic(ctx, s.svcCtx)
	return l.GetApiList(in)
}

// authorization management service
func (s *CoreServer) GetMenuAuthority(ctx context.Context, in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	l := logic.NewGetMenuAuthorityLogic(ctx, s.svcCtx)
	return l.GetMenuAuthority(in)
}

func (s *CoreServer) CreateOrUpdateMenuAuthority(ctx context.Context, in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateMenuAuthorityLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenuAuthority(in)
}

// dictionary management service
func (s *CoreServer) CreateOrUpdateDictionary(ctx context.Context, in *core.DictionaryInfo) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateDictionaryLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateDictionary(in)
}

func (s *CoreServer) DeleteDictionary(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteDictionaryLogic(ctx, s.svcCtx)
	return l.DeleteDictionary(in)
}

func (s *CoreServer) GetDictionaryList(ctx context.Context, in *core.DictionaryPageReq) (*core.DictionaryList, error) {
	l := logic.NewGetDictionaryListLogic(ctx, s.svcCtx)
	return l.GetDictionaryList(in)
}

func (s *CoreServer) GetDetailByDictionaryName(ctx context.Context, in *core.DictionaryDetailReq) (*core.DictionaryDetailList, error) {
	l := logic.NewGetDetailByDictionaryNameLogic(ctx, s.svcCtx)
	return l.GetDetailByDictionaryName(in)
}

func (s *CoreServer) CreateOrUpdateDictionaryDetail(ctx context.Context, in *core.DictionaryDetail) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateDictionaryDetailLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateDictionaryDetail(in)
}

func (s *CoreServer) DeleteDictionaryDetail(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteDictionaryDetailLogic(ctx, s.svcCtx)
	return l.DeleteDictionaryDetail(in)
}

// oauth management
func (s *CoreServer) CreateOrUpdateProvider(ctx context.Context, in *core.ProviderInfo) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateProviderLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateProvider(in)
}

func (s *CoreServer) DeleteProvider(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteProviderLogic(ctx, s.svcCtx)
	return l.DeleteProvider(in)
}

func (s *CoreServer) GetProviderList(ctx context.Context, in *core.PageInfoReq) (*core.ProviderListResp, error) {
	l := logic.NewGetProviderListLogic(ctx, s.svcCtx)
	return l.GetProviderList(in)
}

func (s *CoreServer) OauthLogin(ctx context.Context, in *core.OauthLoginReq) (*core.OauthRedirectResp, error) {
	l := logic.NewOauthLoginLogic(ctx, s.svcCtx)
	return l.OauthLogin(in)
}

func (s *CoreServer) OauthCallback(ctx context.Context, in *core.CallbackReq) (*core.LoginResp, error) {
	l := logic.NewOauthCallbackLogic(ctx, s.svcCtx)
	return l.OauthCallback(in)
}

// Token management
func (s *CoreServer) CreateOrUpdateToken(ctx context.Context, in *core.TokenInfo) (*core.BaseResp, error) {
	l := logic.NewCreateOrUpdateTokenLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateToken(in)
}

func (s *CoreServer) DeleteToken(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := logic.NewDeleteTokenLogic(ctx, s.svcCtx)
	return l.DeleteToken(in)
}

func (s *CoreServer) GetTokenList(ctx context.Context, in *core.TokenListReq) (*core.TokenListResp, error) {
	l := logic.NewGetTokenListLogic(ctx, s.svcCtx)
	return l.GetTokenList(in)
}

func (s *CoreServer) SetTokenStatus(ctx context.Context, in *core.SetStatusReq) (*core.BaseResp, error) {
	l := logic.NewSetTokenStatusLogic(ctx, s.svcCtx)
	return l.SetTokenStatus(in)
}

func (s *CoreServer) BlockUserAllToken(ctx context.Context, in *core.UUIDReq) (*core.BaseResp, error) {
	l := logic.NewBlockUserAllTokenLogic(ctx, s.svcCtx)
	return l.BlockUserAllToken(in)
}
