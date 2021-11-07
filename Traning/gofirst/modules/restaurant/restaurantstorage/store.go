package restaurantstorage

type Create struct {
	d    int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Create) NewSQLStore() {
	return
}
