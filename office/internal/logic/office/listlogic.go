package office

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
	"github.com/kr/pretty"
	"zero-scaffold/common/errorx"

	"zero-scaffold/office/internal/svc"
	"zero-scaffold/office/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.OfficeSearch) (resp *types.OfficeListResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("%# v\n", pretty.Formatter(req))
	fmt.Printf("%# v\n", pretty.Formatter(req.Page))
	fmt.Printf("%# v\n", pretty.Formatter(req.PageSize))

	validateErr := validation.Errors{
		"page":      validation.Validate(req.Page, validation.Required.Error("请填写page")),
		"page_size": validation.Validate(req.PageSize, validation.Required.Error("请填写page_size")),
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
	officeResp.Page = req.Page
	officeResp.PageSize = req.PageSize
	officeResp.Total = count

	for _, item := range all {
		var office types.Office
		copier.Copy(&office, item)
		officeResp.List = append(officeResp.List, office)
	}
	return officeResp, nil
}
