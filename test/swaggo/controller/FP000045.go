package controller

import "github.com/gin-gonic/gin"

// FP000045 godoc
// @Summary     transaction query
// @Description transaction query
// @Tags        paymentSwitch
// @Param       FP000045Request body     models.FP000045I true "transaction query"
// @Success     200                        {object} controller.Response{response=models.FP000045O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert/query [post]
func (c *Controller) FP000045(ctx *gin.Context) {

}
