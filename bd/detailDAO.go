package bd

import (
	"github.com/brayanzv/FamiliarBudget2/models"
	"log"
)

func InsertDetail(d models.Details)(bool,error){

	err := ConectionBD().Create(&d).Error
	ConectionBD().Close()
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetDetailsUserDB(id uint, codFamily string )([]*models.GetDetailsID, bool){
	var searchs []*models.GetDetailsID

	if err := ConectionBD().Table("details").
		Joins("left join users on users.id = details.id_user").
		Select("users.*, details.*").Where("id_user = ?",id).Where("cod_familiar =?",codFamily).Scan(&searchs).Error; err != nil{

			log.Fatal(err)
	}
	err :=ConectionBD().Close()

	if err != nil{
		return nil, false
	}

	return searchs, true
}

func GetDetailsFamilyDB(codFamily string)([]*models.GetDetailsID, bool){
	var searchs []*models.GetDetailsID

	if err := ConectionBD().Table("details").
		Joins("left join users on users.id = details.id_user").
		Select("users.*, details.*").Where("cod_familiar =?",codFamily).Scan(&searchs).Error; err != nil{

		log.Fatal(err)
	}
	err :=ConectionBD().Close()

	if err != nil{
		return nil, false
	}

	return searchs, true
}

func DeleteDetailUserDB(id int) bool{

	if err := ConectionBD().Delete(&models.Details{},id).Error; err != nil{
		log.Fatal(err)
		ConectionBD().Close()
	}
	return true
}