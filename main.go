package main

import (
	_ "nohassls_material2/routers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb") + "?charset=utf8&parseTime=true&loc=Australia%2FSydney"
	orm.RegisterDataBase("default", "mysql", mysqlReg)
	//orm.RegisterModel(new(models.LifeInsurance))
}

func main() {
	//name := "default"
	//force := false
	//verbose := true

	//err := orm.RunSyncdb(name, force, verbose)

	//	if err != nil {
	//	beego.Debug(err)
	//	}
	beego.AddFuncMap("setchecked", SetChecked)
	beego.Run()
	// add simple helper function for setting checkbox chceked values

}

// simple function to return checked if value =1 for checkboxes and blank if null or 0
func SetChecked(in int) (out string) {
	switch in {
	case 1:
		out = "checked"
	case 0:
		out = ""
	default:
		out = ""
	}
	return
}
