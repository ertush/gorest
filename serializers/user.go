package serializers

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Names    string `json:"names"`
	Email    string `json:"string"`
}
