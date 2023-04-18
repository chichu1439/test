package controller

import "github.com/gin-gonic/gin"

// FP000025 godoc
// @Summary     limit query list
// @Description limit query list
// @Tags        paymentSwitch
// @Param       FP000025Request body     models.FP000025I true "limit query list"
// @Success     200                        {object} controller.Response{response=models.FP000025O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/transaction/limit/query/list [post]
func (c *Controller) FP000025(ctx *gin.Context) {

}
