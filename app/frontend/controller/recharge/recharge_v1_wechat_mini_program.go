package recharge

import (
	"context"

	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_amount "server/app/common/utils/validator/amount"
	v1 "server/app/frontend/api/recharge/v1"
	"server/app/frontend/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) WechatMiniProgram(ctx context.Context, req *v1.WechatMiniProgramReq) (res *v1.WechatMiniProgramRes, err error) {
	// 1. 参数验证
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, utils_error.ErrMessage(response.PARAM_INVALID, "参数验证失败")
	}

	// 2. 金额格式额外验证
	amountValidator := &utils_amount.AmountValidator{}
	if err := amountValidator.ValidateAmount(gconv.String(req.Money)); err != nil {
		return nil, utils_error.ErrMessage(response.PARAM_INVALID, err.Error())
	}

	// 3. 格式化金额为两位小数
	formattedAmount, err := amountValidator.ValidateAndFormat(gconv.String(req.Money))
	if err != nil {
		return nil, utils_error.ErrMessage(response.PARAM_INVALID, err.Error())
	}
	req.Money = formattedAmount

	detail, err := service.Recharge().WechatMiniProgram(ctx, req.WechatMiniProgram)
	if err != nil {
		return nil, err
	}
	res = &v1.WechatMiniProgramRes{
		WechatMiniProgram: detail,
	}
	return
}
