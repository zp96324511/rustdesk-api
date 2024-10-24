package admin

import (
	"Gwen/model"
)

type UserForm struct {
	Id       uint   `json:"id"`
	Username string `json:"username" validate:"required,gte=4,lte=12"`
	//Password string           `json:"password" validate:"required,gte=4,lte=20"`
	Nickname string           `json:"nickname"`
	Avatar   string           `json:"avatar"`
	GroupId  uint             `json:"group_id" validate:"required"`
	IsAdmin  *bool            `json:"is_admin" `
	Status   model.StatusCode `json:"status" validate:"required,gte=0"`
}

func (uf *UserForm) FromUser(user *model.User) *UserForm {
	uf.Id = user.Id
	uf.Username = user.Username
	uf.Nickname = user.Nickname
	uf.Avatar = user.Avatar
	uf.GroupId = user.GroupId
	uf.IsAdmin = user.IsAdmin
	uf.Status = user.Status
	return uf
}
func (uf *UserForm) ToUser() *model.User {
	user := &model.User{}
	user.Id = uf.Id
	user.Username = uf.Username
	user.Nickname = uf.Nickname
	user.Avatar = uf.Avatar
	user.GroupId = uf.GroupId
	user.IsAdmin = uf.IsAdmin
	user.Status = uf.Status
	return user
}

type PageQuery struct {
	Page     uint `form:"page"`
	PageSize uint `form:"page_size"`
}

type UserQuery struct {
	PageQuery
	Username string `form:"username"`
}
type UserPasswordForm struct {
	Id       uint   `json:"id" validate:"required"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
}

type ChangeCurPasswordForm struct {
	OldPassword string `json:"old_password" validate:"required,gte=4,lte=20"`
	NewPassword string `json:"new_password" validate:"required,gte=4,lte=20"`
}
