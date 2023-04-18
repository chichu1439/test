package controller

import "github.com/gin-gonic/gin"

// FP000103 godoc
// @Summary     param enquiry
// @Description param enquiry
// @Tags        paymentSwitch
// @Param       FP000103Request body     models.FP000103I true "param enquiry"
// @Success     200                        {object} controller.Response{response=models.FP000103O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/netting/param/enquiry [post]
func (c *Controller) FP000103(ctx *gin.Context) {

}
