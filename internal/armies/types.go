package armies

import "time"

type Army struct {
	Id           int       `json:"id"`
	Name         string    `json:"title"`
	Description  string    `json:"description"`
	UserId       int       `json:"userId"`
	CreateTime   time.Time `json:"createdAt"`
	ModifiedTime time.Time `json:"modifiedAt"`
	IsDeleted    bool      `json:"isDeleted"`
}
