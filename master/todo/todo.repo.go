package todo

import (
	"errors"

	"todoapi.wisnu.net/database"
	"todoapi.wisnu.net/modules"
);

type Todo modules.Todo

func (todo *Todo) Create(p *Todo) (*Todo,error) {
	result := database.DB.Create(&p)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	return p,nil
}

func (todo Todo) Get(id uint) (*Todo,error){
	result := database.DB.Where(&Todo{ID: id}).Find(&todo)
	if result.RowsAffected == 0 {
		
		return nil,errors.New("Error")
	}
	

	return &todo,nil
}

func (todo Todo) GetAll(activityId uint) ([]Todo,error){
	data := []Todo{}
	whereQuery := database.DB

	if activityId > 0 {
		whereQuery = whereQuery.Where(&Todo{ActivityId: activityId})
	}

	result := whereQuery.Find(&data)
	if result.RowsAffected == 0 {
		return nil,errors.New("Error")
	}

	return data,nil
}

func (todo Todo) Delete(id uint,title string) error {
	result := database.DB.Where(&Todo{ID: id,Title: title}).Delete(todo)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (todo Todo) Update(id uint, c *Todo) (*Todo,error) {
	result := database.DB.Model(todo).Where(&Todo{ID: id}).Updates(&c)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	data := database.DB.Where(&Todo{ID: id}).Find(&todo)
	if data.RowsAffected == 0 {
		return nil,data.Error
	}

	return &todo,nil
}