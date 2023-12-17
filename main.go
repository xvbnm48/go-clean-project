package main

import (
	"fmt"
	"github.com/xvbnm48/go-clean-project/src/infrastructure"
	"github.com/xvbnm48/go-clean-project/src/interfaces"
	"github.com/xvbnm48/go-clean-project/src/usecases"
	"log"
	"net/http"
)

func main() {
	//dbHandler := infrastructure.NewSqliteHandler("/var/tmp/production.sqlite")
	dbHandlerMysql, err := infrastructure.NewMysqlHandler("root", "", "store")
	if err != nil {
		log.Fatal("Gagal menginisialisasi handler MySQL:", err)
		panic(err)
	}

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandlerMysql
	handlers["DbCustomerRepo"] = dbHandlerMysql
	handlers["DbItemRepo"] = dbHandlerMysql
	handlers["DbOrderRepo"] = dbHandlerMysql

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)
	orderInteractor.Logger = new(infrastructure.Logger)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	//http.ListenAndServe(":8080", nil)
	fmt.Println("Server running at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed running server")
	}
}
