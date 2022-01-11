package repository

import (
	"fmt"
	"golang_api_kegiatanQ/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type KegiatanRepository interface {
	InsertKegiatanq(h entity.KegiatanQ) entity.KegiatanQ
	UpdateKegiatanq(h entity.KegiatanQ) entity.KegiatanQ
	DeleteKegiatanq(h entity.KegiatanQ)
	AllKegiatanq(userID string) []entity.KegiatanQ
	FindKegiatanqByID(hutangID uint64) entity.KegiatanQ
	AllUser() []entity.User
}

type kegiatanqConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewKegiatanRepository(dbConn *gorm.DB) KegiatanRepository {
	return &kegiatanqConnection{
		connection: dbConn,
	}
}

func (db *kegiatanqConnection) InsertKegiatanq(h entity.KegiatanQ) entity.KegiatanQ {
	db.connection.Save(&h)
	db.connection.Preload("User").Find(&h)
	return h
}

func (db *kegiatanqConnection) UpdateKegiatanq(h entity.KegiatanQ) entity.KegiatanQ {
	db.connection.Save(&h)
	fmt.Println(h.Name)
	db.connection.Preload("User").Find(&h)
	return h
}

func (db *kegiatanqConnection) DeleteKegiatanq(h entity.KegiatanQ) {
	db.connection.Delete(&h)
}

func (db *kegiatanqConnection) FindKegiatanqByID(hutangID uint64) entity.KegiatanQ {
	var hutang entity.KegiatanQ
	db.connection.Preload("User").Find(&hutang, hutangID)
	return hutang
}

func (db *kegiatanqConnection) AllKegiatanq(userID string) []entity.KegiatanQ {
	var hutangs []entity.KegiatanQ
	db.connection.Preload("User").Where("user_id", userID).Find(&hutangs)
	return hutangs

}

func (db *kegiatanqConnection) AllUser() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users

}
