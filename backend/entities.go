package backend

import "gorm.io/gorm"

type ListItem struct {
	gorm.Model
	Title string
}

type ListItemDto struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type Tag struct {
	gorm.Model
	Name string
}

type TagDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type ListItemFact struct {
	gorm.Model
	ListItemID uint
	FactType   uint
	Fact       string
}

type FactType uint

const (
	FactTypeTag FactType = 1
)
