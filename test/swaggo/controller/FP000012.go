package controller

import "github.com/gin-gonic/gin"

// FP000012 godoc
// @Summary     settlement account creation
// @Description settlement account creation
// @Tags        paymentSwitch
// @Param       FP000012Request body     models.FP000012I true "settlement account creation"
// @Success     200                        {object} controller.Response{response=models.FP000012O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/settlement-account/creation [post]
func (c *Controller) FP000012(ctx *gin.Context) {

}
