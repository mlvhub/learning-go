package go_micro_srv_user

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Couldn't create UUID: %s", err)
		return err
	}
	return scope.SetColumn("Id", uuid.String())
}
