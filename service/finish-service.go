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
type KegiatanQFinishService interface {
	UpdateLunas(b dto.IsFinishUpdateDTO) entity.KegiatanQ
	LunasIsAllowedToEdit(userID string, kegiatanID uint64) bool
}

type isfinishService struct {
	isfinishRepository repository.KegiatanQFinishRepository
}

//NewBookService .....
func NewKegiatanQFinishService(isfinishRepo repository.KegiatanQFinishRepository) KegiatanQFinishService {
	return &isfinishService{
		isfinishRepository: isfinishRepo,
	}
}

func (service *isfinishService) UpdateLunas(b dto.IsFinishUpdateDTO) entity.KegiatanQ {
	finishhutang := entity.KegiatanQ{}
	err := smapping.FillStruct(&finishhutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.isfinishRepository.UpdateIsFinishKegiatanQ(finishhutang)
	return res
}

func (service *isfinishService) LunasIsAllowedToEdit(userID string, kegiatanID uint64) bool {
	b := service.isfinishRepository.FindFinishKegiatanQByID(kegiatanID)
	id := fmt.Sprintf("%v", b.UserID)
	fmt.Println(kegiatanID)
	return userID == id
}
