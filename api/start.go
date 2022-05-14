package api

import (
	"github.com/gin-gonic/gin"
	"github.com/se2022-qiaqia/course-system/dao"
	"github.com/se2022-qiaqia/course-system/model/req"
	"github.com/se2022-qiaqia/course-system/model/resp"
)

type Start struct{}

func isInitialized() bool {
	var count int64
	err := dao.DB.Model(&dao.User{}).Count(&count).Error
	return err == nil && count >= 1
}

// IsInitialized
// @Summary					检测是否已初始化系统。
// @Description
// @Tags					初始化
// @Accept					json
// @Produce					json
// @Security				ApiKeyAuth
// @Success 				200 			{object}	boolean
// @Failure 				400 			{object} 	resp.ErrorResponse
// @Router					/init			[get]
func (api Start) IsInitialized(c *gin.Context) {
	resp.Ok(isInitialized(), c)
}

// InitRequest 初始化系统参数
type InitRequest struct {
	Id       uint   `json:"id" binding:"required" description:"管理员ID"`                      // 管理员ID
	Username string `json:"username" binding:"required,username" description:"管理员用户名"`      // 管理员用户名
	Password string `json:"password" binding:"required,password" description:"管理员密码"`       // 管理员密码
	RealName string `json:"realName" binding:"required,min=2,max=10" description:"管理员真实姓名"` // 管理员真实姓名
}

// InitSystem
// @Summary					初始化系统。
// @Description
// @Tags					初始化
// @Accept					json
// @Produce					json
// @Param					params			body		InitRequest		true	"初始化信息"
// @Security				ApiKeyAuth
// @Success 				200 			{object}	boolean
// @Failure 				400 			{object} 	resp.ErrorResponse
// @Router					/init			[post]
func (api Start) InitSystem(c *gin.Context) {
	if isInitialized() {
		resp.FailJust("系统已初始化", c)
		return
	} else {
		var b InitRequest
		if !req.BindAndValidate(c, &b) {
			return
		}

		if b.Id == 0 || b.Username == "" || b.Password == "" {
			// 这里应该交给 req.BindAndValidate 处理
			resp.FailJust("请完整指定有效的ID、用户名、密码", c)
			return
		}

		var adminCollegeName = "_____ADMIN"

		user := &dao.User{
			Model:    dao.Model{ID: b.Id},
			Username: b.Username,
			RealName: b.RealName,
			Role:     dao.RoleAdmin,
			College: dao.College{
				Name: adminCollegeName,
			},
		}
		user.SetPassword(b.Password)
		result := dao.DB.Create(user)
		if err := result.Error; err != nil {
			resp.Fail(resp.ErrCodeFail, err.Error(), c)
			return
		}
		resp.Ok(true, c)
		return
	}
}
