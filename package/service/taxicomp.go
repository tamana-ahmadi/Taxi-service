package service

import (
	"Taxi_service/models"
	"Taxi_service/package/repository"
	"fmt"
)

func AddTaxicomp(txcm models.TaxiComp) error {
	err := repository.InsertTaxicomps(txcm)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateTaxiComp(txcm models.TaxiComp, id int) error {
	err := repository.EditTaxicomps(txcm.CompTitle, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func DeleteTaxiComp(isdeleted bool, id int) error {
	err := repository.SoftDeleteTaxicomps(isdeleted, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func PrintAllTaxiComps(isdeletedt, isblocked, isdeletedu bool) (txcms []models.TaxiComp, err error) {
	txcms, err = repository.GetAllTaxicomps(isdeletedt, isblocked, isdeletedu)
	if err != nil {

		return txcms, err
	}
	return txcms, nil
}
func PrintAllTaxiCompByID(isdeleted, isblocked, isdeletedu bool, id int) (txcm []models.TaxiComp, err error) {
	txcm, err = repository.GetAllTaxicompsByID(isdeleted, isblocked, isdeletedu, id)
	if err != nil {

		return txcm, err
	}
	return txcm, nil
}
