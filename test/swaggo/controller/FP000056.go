package controller

import "github.com/gin-gonic/gin"

// FP000056 godoc
// @Summary     scenario delete
// @Description scenario delete
// @Tags        paymentSwitch
// @Param       FP000056Request body     models.FP000056I true "scenario delete"
// @Success     200                        {object} controller.Response{response=models.FP000056O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/delete [post]
func (c *Controller) FP000056(ctx *gin.Context) {

}
