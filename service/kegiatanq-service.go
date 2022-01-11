package service

import (
	"fmt"
	"golang_api_kegiatanQ/dto"
	"golang_api_kegiatanQ/entity"
	"golang_api_kegiatanQ/repository"
	"log"

	"github.com/mashingan/smapping"
)

//BookService is a ....
type KegiatanQService interface {
	Insert(b dto.KegiatanQCreateDTO) entity.KegiatanQ
	Update(b dto.KegiatanQUpdateDTO) entity.KegiatanQ
	Delete(b entity.KegiatanQ)
	All(userID string) []entity.KegiatanQ
	FindByID(kegiatanqID uint64) entity.KegiatanQ
	IsAllowedToEdit(userID string, kegiatanqID uint64) bool
	AllUser() []entity.User
}

type kegiatanqService struct {
	kegiatanqRepository repository.KegiatanRepository
}

//NewBookService .....
func NewKegiatanQService(kegiatanqRepo repository.KegiatanRepository) KegiatanQService {
	return &kegiatanqService{
		kegiatanqRepository: kegiatanqRepo,
	}
}

func (service *kegiatanqService) Insert(b dto.KegiatanQCreateDTO) entity.KegiatanQ {
	hutang := entity.KegiatanQ{}
	err := smapping.FillStruct(&hutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.kegiatanqRepository.InsertKegiatanq(hutang)
	return res
}

func (service *kegiatanqService) Update(b dto.KegiatanQUpdateDTO) entity.KegiatanQ {
	hutang := entity.KegiatanQ{}
	err := smapping.FillStruct(&hutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.kegiatanqRepository.UpdateKegiatanq(hutang)
	return res
}

func (service *kegiatanqService) Delete(b entity.KegiatanQ) {
	service.kegiatanqRepository.DeleteKegiatanq(b)
}

func (service *kegiatanqService) All(userID string) []entity.KegiatanQ {
	return service.kegiatanqRepository.AllKegiatanq(userID)
}

func (service *kegiatanqService) AllUser() []entity.User {
	return service.kegiatanqRepository.AllUser()
}

func (service *kegiatanqService) FindByID(kegiatanqID uint64) entity.KegiatanQ {
	return service.kegiatanqRepository.FindKegiatanqByID(kegiatanqID)
}

func (service *kegiatanqService) IsAllowedToEdit(userID string, kegiatanqID uint64) bool {
	b := service.kegiatanqRepository.FindKegiatanqByID(kegiatanqID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
