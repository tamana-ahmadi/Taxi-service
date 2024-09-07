package service

import (
	"Taxi_service/models"
	"Taxi_service/package/repository"
	"fmt"
)

func AddOrder(ord models.Order) error {
	err := repository.InsertOrders(ord)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateOrder(ord models.Order, id int) error {
	err := repository.EditOrders(ord.RouteID, ord.ClientID,ord.DriverID, ord.TaxicompID, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func DeleteOrder(isdeleted bool, id int) error {
	err := repository.SoftDeleteOrders(isdeleted, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func CheckOrderasDone(isdone bool, id int) error {
	err := repository.CheckOrdersasDone(isdone, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func PrintAllOrderByID(isdeleted, isdone, isblocked bool, uid, oid uint) (ord []models.Order, err error) {
	ord, err = repository.GetAllOrdersByID(isdeleted, isdone, isblocked, uid, oid)
	if err != nil {

		return ord, err
	}
	return ord, nil
}

func ReportofOrder(isdone, isdeletedo, isdeletedt, isdeletedr, isresp, isblocked,isdeletedu bool) (ord []models.Order, err error) {
	ord, err = repository.OrdersReport(isdone, isdeletedo, isdeletedt, isdeletedr, isresp, isblocked,isdeletedu)
	if err != nil {
		return ord, err
	}
	return ord, nil

}
