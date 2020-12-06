/*
 * @Author: Jame Yu
 * @Date: 2020-12-06 11:33:33
 * @LastEditors: Jame Yu
 * @LastEditTime: 2020-12-06 22:15:00
 * @FilePath: /app/routers/router.go
 */
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/post",
			beego.NSInclude(
				&controllers.PostController{},
			),
		),
		beego.NSNamespace("/device",
			beego.NSInclude(
				&controllers.DeviceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
