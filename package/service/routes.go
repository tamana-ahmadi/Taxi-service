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
	err := repository.EditRoutes(route.From, route.Into, route.Distance, route.Pricekm, route.IsResponse, did, id)
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

func Printreport(isresp, isdeletedr, isdeletedt, isblocked, isdeletedu bool, price int) (route []models.GetRoutes, err error) {
	route, err = repository.Report(isresp, isdeletedr, isdeletedt, isblocked, isdeletedu, price)
	if err != nil {

		return route, err
	}
	return route, nil
}
func PrintAllRoutes(isdeleted, isresp bool, price int) (route []models.Route, err error) {
	route, err = repository.GetAllRoutes(isdeleted, isresp, price)
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
