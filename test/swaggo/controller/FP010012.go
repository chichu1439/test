package controller

import "github.com/gin-gonic/gin"

// FP010012 godoc
// @Summary     direct debit
// @Description direct debit
// @Tags        paymentSwitch
// @Param       FP010012Request body     models.FP010012I true "direct debit"
// @Success     200                        {object} controller.Response{response=models.FP010012O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/direct-debit [post]
func (c *Controller) FP010012(ctx *gin.Context) {

}
