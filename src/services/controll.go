package services

import (
	"time"
	"github.com/LukaGiorgadze/gonull"
)



// Untuk users model
type Person struct {
    Id 				int 							`db:"id" json:"id"`
	FullName 		gonull.Nullable[string] 		`db:"fullName" json:"fullName" form:"fullName"`
	Email 			string 							`db:"email" json:"email" form:"email"`
	PhoneNumber 	gonull.Nullable[string] 		`db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Address 		gonull.Nullable[string] 		`db:"address" json:"address" form:"address"`
	Picture 		gonull.Nullable[string] 		`db:"picture" json:"picture" form:"picture"`
	Role 			string 							`db:"role" json:"role"`
	Password 		string 							`db:"password" json:"password" form:"password"`
	CreatedAt 		time.Time 						`db:"createdAt" json:"createdAt"`
	UpdatedAt 		gonull.Nullable[time.Time] 		`db:"updatedAt" json:"updatedAt"`
}

// Untuk users Register & Login
type RLUsers struct {
	Email 			string 			`db:"email" json:"email" form:"email" binding:"email"`
	Role 			string 			`db:"role" json:"role"`
	Password 		string 			`db:"password" json:"password" form:"password"`
}
// Untuk users Controller
type ResponseList struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
	Results interface{}	`json:"results"`
}

// Untuk file utils
type ResponseBack struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
}