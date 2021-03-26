/*
--------------------------------------------------
 File Name: uuid.go
 Author: hanxu

 Created Time: 2019-8-31-上午11:03
---------------------说明--------------------------

---------------------------------------------------
*/

package model

import (
	"github.com/google/uuid"
)

type ModelUuid struct {
	//Uuid string `gorm:"type:varchar(32);unique;" json:"ID"`
	Uuid string `gorm:"primary_key"`
}

func ParseUuid(id string) ModelUuid {
	guid, e := uuid.Parse(id)
	if e != nil {
		panic(e)
	}
	return ModelUuid{
		Uuid: guid.String(),
	}
}

func (this *ModelUuid) IsEmpty() bool {
	return this.Uuid == ""
}

func (this *ModelUuid) GetUUid() string {
	return this.Uuid
}
