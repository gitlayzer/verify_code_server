package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gitlayzer/callback_platform/utils"
	"github.com/google/uuid"
	"net/http"
)

var (
	codeMap = make(map[string]string)
)

func SendMail(c *gin.Context) {
	// 生成随机验证码
	randCode := utils.GenerateCode()

	//生成唯一的ID
	id := uuid.New().String()

	// 将验证码存储到映射关系中
	codeMap[id] = randCode

	//发送验证码的逻辑，这里只是简单地打印验证码
	email := c.Param("email")
	err := utils.SendVerificationCode(email, randCode, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "验证码发送失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "验证码已发送",
	})
}

func Verify(c *gin.Context) {
	// 从URL参数中获取ID和验证码
	id := c.Param("id")
	// 从请求体中获取验证码
	code := c.Param("code")

	// 比较验证码
	if savedCode, ok := codeMap[id]; ok {
		if code == savedCode {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"message": "验证成功",
			})
		} else {
			c.HTML(http.StatusUnauthorized, "index.tmpl", gin.H{
				"message": "验证失败",
			})
		}
	} else {
		c.HTML(http.StatusNotFound, "index.tmpl", gin.H{
			"message": "验证失败",
		})
	}
}
