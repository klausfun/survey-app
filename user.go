package survey

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	Email    string `json:"email" binding:"required" db:"email"`
	Role     string `json:"role" binding:"required" db:"role"`
	//AdminCode string `json:"adminCode" binding:"required"`
}
