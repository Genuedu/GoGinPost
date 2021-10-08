package entity

// album represents data about a record album.
type Album struct {
	ID     string  `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title  string  `gorm:"size:100;not null;unique" json:"title"`
	Artist string  `gorm:"size:100;not null" json:"artist"`
	Price  float64 `gorm:"size:10;not null" json:"price"`
	//	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UpdateAlbumInput struct {
	Title  string  `gorm:"size:100;not null;unique" json:"title"`
	Artist string  `gorm:"size:100;not null" json:"artist"`
	Price  float64 `gorm:"size:10;not null" json:"price"`  
  }


