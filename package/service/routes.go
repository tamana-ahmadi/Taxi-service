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
	err := repository.EditRoutes(route.From, route.Into, route.DriverID, route.Distance, route.Price, route.IsResponse, id)
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
func CheckRouteasResponse(route models.Route, cid, id int) error {
	err := repository.CheckRoutesAsResponse(route.IsResponse, cid, id)
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
