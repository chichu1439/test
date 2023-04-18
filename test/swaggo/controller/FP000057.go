package controller

import "github.com/gin-gonic/gin"

// FP000057 godoc
// @Summary     step creation
// @Description step creation
// @Tags        paymentSwitch
// @Param       FP000057Request body     models.FP000057I true "step creation"
// @Success     200                        {object} controller.Response{response=models.FP000057O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/step/creation [post]
func (c *Controller) FP000057(ctx *gin.Context) {

}
