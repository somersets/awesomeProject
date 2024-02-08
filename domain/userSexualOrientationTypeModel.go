package domain

type UserSexualOrientationType struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Orientation string `json:"orientation"`
	User        []User `gorm:"foreignKey:SexOrientationID" json:"-"`
}

func (UserSexualOrientationType) TableName() string { return "sex-orientation-types" }
