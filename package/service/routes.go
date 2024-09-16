package service

import (
	"Taxi_service/models"
	"Taxi_service/package/repository"
	"fmt"
)

func AddRoute(routes models.Route) error {
	err := repository.InsertRoutes(routes)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateRoute(route models.Route, did, id int) error {
	err := repository.EditRoutes(route.From, route.Into, route.Distance, route.Price, route.IsResponse, did, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func DeleteRoute(isdeleted bool, id int) error {
	err := repository.SoftDeleteRoutes(isdeleted, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func CheckRouteasResponse(isresp bool, cid, id int) error {
	err := repository.CheckRoutesAsResponse(isresp, cid, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func PrintAllRoutes(isresp, isdeleted bool, price int) (route []models.Route, err error) {
	route, err = repository.GetAllRoutes(isresp, isdeleted, price)
	if err != nil {

		return route, err
	}
	return route, nil
}
func PrintAllRouteByID(isdeleted bool, id uint) (route []models.Route, err error) {
	route, err = repository.GetAllRoutesByID(isdeleted, id)
	if err != nil {

		return route, err
	}
	return route, nil
}

func OrdersReport(isresp, isdeletedr, isblocku, isdeletedu, isdeletedt bool) (rep []models.OrdersReport, err error) {
	rep, err = repository.OrdersReport(isresp, isdeletedr, isblocku, isdeletedu, isdeletedt)
	if err != nil {
		return rep, err
	}
	return rep, nil
}
