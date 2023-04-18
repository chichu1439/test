package controller

import "github.com/gin-gonic/gin"

// FP000013 godoc
// @Summary     random code creation
// @Description random code creation
// @Tags        paymentSwitch
// @Param       FP000013Request body     models.FP000013I true "random code creation"
// @Success     200                        {object} controller.Response{response=models.FP000013O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/random-code/creation [post]
func (c *Controller) FP000013(ctx *gin.Context) {

}
