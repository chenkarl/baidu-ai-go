package controllers

import (
	"fmt"
	"log"
	"time"

	"../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/httplib"
)

// CommenController operations for Commen
type CommenController struct {
	beego.Controller
}

type BaiduAccessTokenResult struct {
	Refresh_token  string `json:"refresh_token"`
	Expires_in     int32  `json:"expires_in"`
	Session_key    string `json:"session_key"`
	Access_token   string `json:"access_token"`
	Scope          string `json:"scope"`
	Session_secret string `json:"session_secret"`
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

// Get ...
// @Title Get
// @Description get BaiduAccessToken save to redis
func (c *CommenController) Get() {
	url := "http://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + beego.AppConfig.String("client_id") + "&client_secret=" + beego.AppConfig.String("client_secret")
	fmt.Println(url)
	req := httplib.Get(url)
	var result BaiduAccessTokenResult
	err := req.ToJSON(&result)
	if err != nil {
		log.Fatal(err)
	}

	bm, err := cache.NewCache("redis", `{"conn": "127.0.0.1:6379","dbNum":"0"}`)
	if err != nil {
		log.Fatal(err)
	}
	err = bm.Put("BaiduAccessToken", result.Access_token, 30*24*3600*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Access_token)
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
