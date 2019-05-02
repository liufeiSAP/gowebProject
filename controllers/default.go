package controllers

import (
	"gowebProject/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	models.CreateUsers()
	models.ListUsers()
	models.CountUser()
	models.GetUser()
	models.GetUsers()
	models.LimitoffsetUser()
	models.DelUser()
	models.UpdateUser()
}
