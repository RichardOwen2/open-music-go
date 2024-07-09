package domain


type Album struct {
	ID     string `gorm:"type:varchar(50);primaryKey"`
	Name   string `gorm:"type:text;not null"`
	Year   int    `gorm:"type:int;not null"`
	Songs  []Song `gorm:"foreignKey:AlbumID"`
}
