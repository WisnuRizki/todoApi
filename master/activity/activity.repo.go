package activity

import (
	"errors"

	"todoapi.wisnu.net/database"
	"todoapi.wisnu.net/modules"
);

type Activity modules.Activity

func (activity *Activity) TableName() string {
	return "activites"
}

func (activity *Activity) Create(p *Activity) (*Activity,error) {

	result := database.DB.Create(&p)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	return p,nil
}

func (activity Activity) Get(id uint) (*Activity,error){
	result := database.DB.Where(&Activity{ID: id}).Find(&activity)
	if result.RowsAffected == 0 {
		
		return nil,errors.New("Error")
	}
	

	return &activity,nil
}

func (activity Activity) GetAll() ([]Activity,error){
	data := []Activity{}
	result := database.DB.Find(&data)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	return data,nil
}

func (activity Activity) Update(id uint, c *Activity) (*Activity,error) {
	result := database.DB.Model(activity).Where(&Activity{ID: id}).Updates(&c)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	data := database.DB.Where(&Activity{ID: id}).Find(&activity)
	if data.RowsAffected == 0 {
		return nil,data.Error
	}

	return &activity,nil
}

func (activity Activity) Delete(id uint) error {
	result := database.DB.Unscoped().Where(&Activity{ID: id}).Delete(activity)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}