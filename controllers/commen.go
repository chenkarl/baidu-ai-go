package controllers

import (
	"fmt"

	"../models"
	"github.com/astaxie/beego"
)

// CommenController operations for Commen
type CommenController struct {
	beego.Controller
}

// URLMapping ...
func (c *CommenController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Commen
// @Param	body		body 	models.Commen	true		"body for Commen content"
// @Success 201 {object} models.Commen
// @Failure 403 body is empty
// @router / [post]
func (c *CommenController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Commen by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Commen
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommenController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Commen
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Commen
// @Failure 403
// @router / [get]
func (c *CommenController) GetAll() {
	users := models.GetAllUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

func (c *CommenController) Get() {
	c.Ctx.WriteString("hello")
	fmt.Println("test")
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}

// Put ...
// @Title Put
// @Description update the Commen
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Commen	true		"body for Commen content"
// @Success 200 {object} models.Commen
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommenController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Commen
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommenController) Delete() {

}
