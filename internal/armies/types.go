package armies

import "time"

type Army struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	UserId       string    `json:"userId"`
	CreateTime   time.Time `json:"createTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
	IsDeleted    bool      `json:"isDeleted"`
}
