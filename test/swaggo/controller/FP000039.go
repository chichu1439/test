package controller

import "github.com/gin-gonic/gin"

// FP000039 godoc
// @Summary     history query list
// @Description history query list
// @Tags        paymentSwitch
// @Param       FP000039Request body     models.FP000039I true "history query list"
// @Success     200                        {object} controller.Response{response=models.FP000039O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/prod/fee/history/query/list [post]
func (c *Controller) FP000039(ctx *gin.Context) {

}
