package main

import (
	"github.com/PonyWilliam/go-category/domain/repository"
	services2 "github.com/PonyWilliam/go-category/domain/service"
	"github.com/PonyWilliam/go-category/handler"
	category2 "github.com/PonyWilliam/go-category/proto"
	"github.com/PonyWilliam/go-common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"strconv"
	"time"
)

func main() {
	//配置中心
	consulConfig,err := common.GetConsualConfig("127.0.0.1",8500,"/micro/config")
	if err != nil{
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options){
		options.Addrs = []string{"127.0.0.1"}
		options.Timeout = 10 * time.Second
	})
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.services.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
		)
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		log.Error(err)

	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	rp := repository.NewCategoryRepository(db)
	rp.InitTable()

	categoryService := services2.NewCategoryService(repository.NewCategoryRepository(db))
	err = category2.RegisterCategoryHandler(srv.Server(),&handler.Category{
		CategoryService: categoryService,
	})

	if err!=nil{
		log.Fatal(err)
	}
	if err:=srv.Run();err!=nil{
		log.Fatal(err)
	}
}
