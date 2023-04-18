package controller

import "github.com/gin-gonic/gin"

// FP000035 godoc
// @Summary     download file
// @Description download file
// @Tags        paymentSwitch
// @Param       FP000035Request body     models.FP000035I true "download file"
// @Success     200                        {object} controller.Response{response=models.FP000035O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/download/file [post]
func (c *Controller) FP000035(ctx *gin.Context) {

}
