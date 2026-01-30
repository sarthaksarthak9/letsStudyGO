package internal

type Notes struct {
	Id     int    `gorm:"primaryKey" json:"id"` 
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// to register this in db we have to go in services
