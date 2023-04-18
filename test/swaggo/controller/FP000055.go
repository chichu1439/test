package controller

import "github.com/gin-gonic/gin"

// FP000055 godoc
// @Summary     scenario update
// @Description scenario update
// @Tags        paymentSwitch
// @Param       FP000055Request body     models.FP000055I true "scenario update"
// @Success     200                        {object} controller.Response{response=models.FP000055O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/update [post]
func (c *Controller) FP000055(ctx *gin.Context) {

}
