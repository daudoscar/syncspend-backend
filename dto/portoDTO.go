package dto

type UpdatePortoDTO struct {
	ID_Owner uint64 `json:"id_owner" binding:"required"`
	ID       uint64 `json:"id"`
}

type GetPortoDTO struct {
	ID       uint64 `json:"id" binding:"required"`
	ID_Owner uint64 `json:"id_owner" binding:"required"`
}

type GetPortoResponse struct {
	ID_Owner uint64 `json:"id_owner" binding:"required"`
	ID       uint64 `json:"id"`
}
