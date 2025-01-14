package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(in *core.IDReq) (*core.BaseResp, error) {

	err := utils.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		_, err := tx.DictionaryDetail.Delete().Where(dictionarydetail.HasDictionaryWith(dictionary.IDEQ(in.Id))).Exec(l.ctx)
		if err != nil {
			return err
		}

		err = tx.Dictionary.DeleteOneID(in.Id).Exec(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logx.Errorf("delete dictionary failed, error : %s", err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
