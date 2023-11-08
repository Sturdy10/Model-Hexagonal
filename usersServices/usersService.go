package usersServices

import (
	"log"
	"users/usersModels"
	"users/usersRepository"
)

type IServices interface {
	PotsRegisterSer(register usersModels.RequestRegister) error
}

type services struct {
	r usersRepository.IRepository
}

func NewServices(r usersRepository.IRepository) IServices {
	return &services{r: r}
}

func (s *services) PotsRegisterSer(register usersModels.RequestRegister) error {

	err := s.r.PotsRegister(register)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
