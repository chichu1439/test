package controller

import "github.com/gin-gonic/gin"

// FP000026 godoc
// @Summary     query history
// @Description query history
// @Tags        paymentSwitch
// @Param       FP000026Request body     models.FP000026I true "query history"
// @Success     200                        {object} controller.Response{response=models.FP000026O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/transaction/limit/query/history [post]
func (c *Controller) FP000026(ctx *gin.Context) {

}
