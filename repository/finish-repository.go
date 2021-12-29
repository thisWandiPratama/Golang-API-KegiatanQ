package repository

import (
	"golang_api_kegiatanQ/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type KegiatanQFinishRepository interface {
	UpdateIsFinishKegiatanQ(h entity.KegiatanQ) entity.KegiatanQ
	FindFinishKegiatanQByID(kegiatanqID uint64) entity.KegiatanQ
}

type kegiatanqfinishConnection struct {
	connection *gorm.DB
}

func NewKegiatanQFinishRepository(dbConn *gorm.DB) KegiatanQFinishRepository {
	return &kegiatanqfinishConnection{
		connection: dbConn,
	}
}

func (db *kegiatanqfinishConnection) FindFinishKegiatanQByID(kegiatanqID uint64) entity.KegiatanQ {
	var kegiatanq entity.KegiatanQ
	db.connection.Preload("User").Find(&kegiatanq, kegiatanqID)
	return kegiatanq
}

func (db *kegiatanqfinishConnection) UpdateIsFinishKegiatanQ(h entity.KegiatanQ) entity.KegiatanQ {
	db.connection.Preload("User").First(&h)
	h.IsFinish = !h.IsFinish
	db.connection.Save(&h)
	return h
}
