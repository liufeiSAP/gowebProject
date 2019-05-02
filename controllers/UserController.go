package controllers

import (
	"encoding/json"
	"fmt"
	md "gowebProject/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

}

func (this *UserController) AddUser() {
	var ob md.User                     //这是一个model，struct类型
	body := this.Ctx.Input.RequestBody //这是获取到的json二进制数据
	json.Unmarshal(body, &ob)          //解析二进制json，把结果放进ob中
	// user := &md.User{Id: ob.UserName, Mobile: ob.Mobile}

	md.CreateUser(&ob) //这是添加用户函数
	// if nil != err {
	//     this.Data["json"] = map[string]interface{}{"result": false, "msg": err}
	// } else {
	//     this.Data["json"] = map[string]interface{}{"result": true, "msg": "新增成功"}
	// }
	// this.Data["json"] = map[string]interface{}{"result": false, "msg": err}
	this.Data["json"] = ob
	this.ServeJSON()

}
