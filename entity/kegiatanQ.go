package entity

//Book struct represents books table in database
type KegiatanQ struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"title"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	IsFinish  bool   `gorm:"type:bool;default:0;not null" json:"is_finish"`
	UserID    uint64 `gorm:"not null" json:"-"`
	User      User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
