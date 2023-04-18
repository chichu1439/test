package controller

import "github.com/gin-gonic/gin"

// FP000050 godoc
// @Summary     scenario list
// @Description scenario list
// @Tags        paymentSwitch
// @Param       FP000050Request body     models.FP000050I true "scenario list"
// @Success     200                        {object} controller.Response{response=models.FP000050O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/list [post]
func (c *Controller) FP000050(ctx *gin.Context) {

}
