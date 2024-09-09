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
func UpdateRoute(route models.Route, id int) error {
	err := repository.EditRoutes(route.From, route.Into, route.DriverID, route.Distance, route.Price, id)
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
func CheckRouteasResponse(isresp bool, id int) error {
	err := repository.CheckRoutesAsResponse(isresp, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func PrintAllRoutes(isdeleted, isresp, isblocked bool, price int, uid uint) (route []models.Route, err error) {
	route, err = repository.GetAllRoutes(isdeleted, isresp, isblocked, price, uid)
	if err != nil {

		return route, err
	}
	return route, nil
}
func PrintAllRouteByID(isdeleted, isresp, isblocked bool, uid, rid uint) (route []models.Route, err error) {
	route, err = repository.GetAllRoutesByID(isdeleted, isresp, isblocked, uid, rid)
	if err != nil {

		return route, err
	}
	return route, nil
}
