package serializers

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
