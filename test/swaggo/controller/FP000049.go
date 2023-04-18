package controller

import "github.com/gin-gonic/gin"

// FP000049 godoc
// @Summary     transaction trigger
// @Description transaction trigger
// @Tags        paymentSwitch
// @Param       FP000049Request body     models.FP000049I true "transaction trigger"
// @Success     200                        {object} controller.Response{response=models.FP000049O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert/trigger [post]
func (c *Controller) FP000049(ctx *gin.Context) {

}
