package controller

import "github.com/gin-gonic/gin"

// FP000046 godoc
// @Summary     participant add
// @Description participant add
// @Tags        paymentSwitch
// @Param       FP000046Request body     models.FP000046I true "participant add"
// @Success     200                        {object} controller.Response{response=models.FP000046O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert/participant/add [post]
func (c *Controller) FP000046(ctx *gin.Context) {

}
