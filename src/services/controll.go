package services

import (
	"time"
	"github.com/LukaGiorgadze/gonull"
)




//------------ UNIVERSAL ------------
// Response ALL DATA
type ResponseList struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
	Results interface{}	`json:"results"`
}

// RESPONSE ERROR
type ResponseBack struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
}







//------------ USERS ------------
// Untuk users model
type Person struct{
    Id 				int 							`db:"id" json:"id"`
	FullName 		string 							`db:"fullName" json:"fullName" form:"fullName"`
	Email 			string 							`db:"email" json:"email" form:"email"`
	PhoneNumber 	string 							`db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Address 		string 							`db:"address" json:"address" form:"address"`
	Picture 		string 							`db:"picture" json:"picture" form:"picture"`
	Role 			string 							`db:"role" json:"role"`
	Password 		string 							`db:"password" json:"password" form:"password"`
	CreatedAt 		time.Time 						`db:"createdAt" json:"createdAt"`
	UpdatedAt 		gonull.Nullable[time.Time] 		`db:"updatedAt" json:"updatedAt"`
}
type PersonNet struct{ // untuk struck respon saja dimana data yg tidak di isi diperbolehkan nil/nul 
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
type RLUsers struct{
	Email 			string 			`db:"email" json:"email" form:"email" binding:"email"`
	Role 			string 			`db:"role" json:"role"`
	Password 		string 			`db:"password" json:"password" form:"password"`
}

// Untuk forgot password
type FormReset struct{
	Id 						int 							`db:"id"`
	Email 					string 							`db:"email" form:"email" binding:"email"`
	Otp 					string 							`db:"otp" form:"otp"`
	Password 				string 							`form:"password"`
	ConfirmPassword 		string 							`form:"confirmPassword"`
	CreatedAt 				time.Time 						`db:"createdAt"`
	UpdatedAt 				gonull.Nullable[time.Time] 		`db:"updatedAt"`
}






//------------ PRODUCTS ------------
// Untuk Products model
type Products struct{
	Id 				int 							`db:"id" json:"id"`
	Name 			string 							`db:"name" json:"name" form:"name"`
	Price 			int64 							`db:"price" json:"price" form:"price"`
	Image 			gonull.Nullable[string] 		`db:"image" json:"image" form:"image"`
	Description 	string 							`db:"description" json:"description" form:"description"`
	Discount 		int32 							`db:"discount" json:"discount" form:"discount"`
	IsRecommended 	gonull.Nullable[bool] 			`db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	Qty 			int64 							`db:"qty" json:"qty" form:"qty"`
	IsActive 		gonull.Nullable[bool] 			`db:"isActive" json:"isActive" form:"isActive"`
	CreatedAt 		time.Time 						`db:"createdAt" json:"createdAt"`
	UpdatedAt 		gonull.Nullable[time.Time] 		`db:"updatedAt" json:"updatedAt"`
}