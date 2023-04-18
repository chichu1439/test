package controller

import "github.com/gin-gonic/gin"

// FP000067 godoc
// @Summary     host query
// @Description host query
// @Tags        paymentSwitch
// @Param       FP000067Request body     models.FP000067I true "host query"
// @Success     200                        {object} controller.Response{response=models.FP000067O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/host/query [post]
func (c *Controller) FP000067(ctx *gin.Context) {

}
