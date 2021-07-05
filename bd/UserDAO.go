package bd

import (
	"github.com/brayanzv/FamiliarBudget2/models"
)

/*InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertRegister(u models.Users) (string, bool, error) {

	u.Password, _ = EncrypPass(u.Password)
	err := ConectionBD().Create(&u).Close()
	if err != nil {
		return "", false, err
	}
	return u.Name, true, nil
}

func ModifyUsers(u models.Users, ID uint) (bool, error) {
	register := make(map[string]interface{})
	if len(u.Name) > 0{
		register["name"]= u.Name
	}
	if len(u.Last_name) > 0{
		register["last_name"]=u.Last_name
	}
	if len(u.Password)> 0 {
		register["password"],_=EncrypPass(u.Password)
	}
	if u.Id_role > 0 && u.Id_role <4 {
		register["id_role"]=u.Id_role
	}
	if u.Status ==true || u.Status==false  {
		register["status"]=u.Status
	}
	if len(u.Cod_familiar) > 0  {
		register["cod_familiar"]=u.Cod_familiar
	}

	err:=ConectionBD().Where("id=?",ID).Model(&u).Updates(register).Error
	ConectionBD().Close()
	if err!= nil{
		return  false, err
	}
	return  true, err
}

func FamilyModifyUser(u models.Users, ID uint,codFamily string) (bool, error) {
	register := make(map[string]interface{})
	if u.Status ==true || u.Status==false  {
		register["status"]=u.Status
	}

	err:=ConectionBD().Where("id=?",ID).Where("cod_familiar=?",codFamily).Model(&u).Updates(register).Error
	ConectionBD().Close()
	if err!= nil{
		return  false, err
	}
	return  true, err
}

func GetUsers()([]*models.Users, bool, int64){

	var searchs []*models.Users
	result:=ConectionBD().Table("users").Find(&searchs)
	lent :=result.RowsAffected
	err:=result.Error
	fail := result.Close()

	if fail != nil{
		return searchs, false, lent
	}

	if err != nil{
		return searchs, false, lent
	}

	return searchs, true,  lent
}
func Delete(id int)(error, int){

	delete:=ConectionBD().Delete(&models.Users{},id)
	err:=delete.Error

	if err!=nil{
		return err, 1
	}

	err=delete.Close()

	if err!=nil{
		return err, 2
	}
	return nil, 0
}