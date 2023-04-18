package controller

import "github.com/gin-gonic/gin"

// FP000021 godoc
// @Summary     transaction limit add
// @Description transaction limit add
// @Tags        paymentSwitch
// @Param       FP000021Request body     models.FP000021I true "transaction limit add"
// @Success     200                        {object} controller.Response{response=models.FP000021O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/transaction/limit/add [post]
func (c *Controller) FP000021(ctx *gin.Context) {

}
