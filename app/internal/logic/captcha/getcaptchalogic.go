package captcha

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"zero-scaffold/app/internal/svc"
	"zero-scaffold/app/internal/types"
	"zero-scaffold/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *types.IDPathReq) (resp *types.OfficeListResp, err error) {
	// todo: add your logic here and delete this line
	validateErr := validation.Errors{
		"id": validation.Validate(req.Id, validation.Required.Error("请填写id")),
		//"name1": validation.Validate(req.Id, validation.Required.Error("name1"), validation.Length(5, 20)),
	}.Filter()
	if errs, ok := validateErr.(validation.Errors); ok {
		fmt.Printf("%# v\n", pretty.Formatter(errs))
		for _, msg := range errs {
			return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, msg.Error())
		}
	}
	queryBuilder := l.svcCtx.WxOfficeAccount.RowBuilder()
	countBuilder := l.svcCtx.WxOfficeAccount.CountBuilder("id")
	count, err := l.svcCtx.WxOfficeAccount.FindCount(l.ctx, countBuilder)

	all, err := l.svcCtx.WxOfficeAccount.FindPageListByPage(l.ctx, queryBuilder, 1, 3, "")

	var officeResp = &types.OfficeListResp{}
	officeResp.Page = 1
	officeResp.PageSize = 10
	officeResp.Total = count

	for _, item := range all {
		var office types.OfficeResp
		copier.Copy(&office, item)
		officeResp.List = append(officeResp.List, office)
	}
	return officeResp, errors.Wrapf(errorx.NewCodeError(errorx.ERR_PARAMS, "请尽快处理"), "test")
}
