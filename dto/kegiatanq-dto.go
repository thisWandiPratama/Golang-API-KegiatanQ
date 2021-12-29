package dto

type KegiatanQUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Name      string `json:"title" form:"title" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	UserID    uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type KegiatanQCreateDTO struct {
	Name      string `json:"title" form:"title" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	UserID    uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type IsFinishUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	IsFinish bool   `json:"is_finish" form:"is_finish"`
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
