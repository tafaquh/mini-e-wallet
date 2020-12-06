package dto

type Transfer struct {
	Amount	int  `json:"amount"`
	UserTargetId int `json:"user_target_id"`
}