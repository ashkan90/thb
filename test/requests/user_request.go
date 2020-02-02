package requests

import (
	"thb/system"
)

type UserRequest struct {
	Name            string `validate:"required"`
	Surname         string `validate:"required"`
	*system.Request `validate:"-"`
}

// reflect.ValueOf(T).MethodByName("TMethod").Type().In(0).Elem().Name()
func (ur *UserRequest) Validate() error {
	err := ur.Request.Validator().Struct(ur)

	return err
}
