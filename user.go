package coffee

type User struct {
	Id         int    `json:"-" db:"id"`
	PhoneCode  string `json:"phone_code" binding:"required" db:"phone_code"`
	Phone      string `json:"phone" binding:"required" db:"phone"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Email      string `json:"email" db:"email"`
	Birthday   string `json:"birthday" db:"birthday"`
	Value      int    `json:"value" db:"value"`
	MessageKey string `json:"message_key" db:"message_key"`
}