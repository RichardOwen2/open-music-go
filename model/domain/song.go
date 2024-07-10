package domain

type Song struct {
	ID        string  `gorm:"type:varchar(50);primaryKey"`
	Title     string  `gorm:"type:text;not null"`
	Year      int     `gorm:"type:int;not null"`
	Genre     string  `gorm:"type:text;not null"`
	Performer string  `gorm:"type:text;not null"`
	Duration  int     `gorm:"type:int"`
	AlbumID   *string `gorm:"type:varchar(50)"`
}
