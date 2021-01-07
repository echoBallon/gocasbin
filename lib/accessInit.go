package lib

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	//改变顺序
	initDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal()
	}
	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal()
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal()
	}
	E = e
	//初始化p
	initPolicy()
}

//初始化p
func initPolicy() {
	//角色层级
	m:=make([]*RoleRel,0)
	GetRoles(0,&m,"")
	for _,r:=range m{
		_,err:=E.AddRoleForUser(r.PRole,r.Role)
		if err!=nil{
			log.Fatal(err)
		}
	}
	/////// 初始化用户角色
	userRoles:=GetUserRoles()
	for _,ur:=range userRoles{
		_,err:=E.AddRoleForUser(ur.UserName,ur.RoleName)
		if err!=nil{
			log.Fatal(err)
		}
	}
	//初始化路由角色
	routerRoles:=GetRouterRoles()
	fmt.Println(routerRoles)
	for _,rr:=range routerRoles{
		_,err:=E.AddPolicy(rr.RoleName,rr.RouterUri,rr.RouterMethod)
		if err!=nil{
			log.Fatal(err)
		}
	}
}
