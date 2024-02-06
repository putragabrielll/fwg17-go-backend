package services

import (
	"github.com/LukaGiorgadze/gonull"
	"time"
)







// ------------ UNIVERSAL ------------
// Page Info
type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	NextPage    int `json:"nextPage"`
	PrevPage    int `json:"prevPage"`
	TotalData   int `json:"totalData"`
}

// Response ALL DATA
type ResponseAll struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type ResponseList struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

// RESPONSE ERROR
type ResponseBack struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}







// ------------ USERS ------------
// Untuk users model
type Person struct {
	Id          int                        `db:"id" json:"id"`
	FullName    string                     `db:"fullName" json:"fullName" form:"fullName"`
	Email       string                     `db:"email" json:"email" form:"email" binding:"email" binding:"required"`
	PhoneNumber string                     `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Address     string                     `db:"address" json:"address" form:"address"`
	Picture     string                     `db:"picture" json:"picture" form:"picture"`
	Role        string                     `db:"role" json:"role"`
	Password    string                     `db:"password" json:"password" form:"password" binding:"required"`
	CreatedAt   time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt   gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}
type PersonNet struct { // untuk struck respon saja dimana data yg tidak di isi diperbolehkan nil/nul
	Id          int                        `db:"id" json:"id"`
	FullName    gonull.Nullable[string]    `db:"fullName" json:"fullName" form:"fullName"`
	Email       string                     `db:"email" json:"email" form:"email"`
	PhoneNumber gonull.Nullable[string]    `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Address     gonull.Nullable[string]    `db:"address" json:"address" form:"address"`
	Picture     gonull.Nullable[string]    `db:"picture" json:"picture" form:"picture"`
	Role        string                     `db:"role" json:"role"`
	Password    string                     `db:"password" json:"password" form:"password"`
	CreatedAt   time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt   gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}

// Untuk users Register & Login
type RLUsers struct {
	Email    string `db:"email" json:"email" form:"email" binding:"email" binding:"required"`
	Role     string `db:"role" json:"role"`
	Password string `db:"password" json:"password" form:"password" binding:"required"`
}

// Untuk forgot password
type FormReset struct {
	Id              int                        `db:"id"`
	Email           string                     `db:"email" form:"email" binding:"email"`
	Otp             string                     `db:"otp" form:"otp"`
	Password        string                     `form:"password"`
	ConfirmPassword string                     `form:"confirmPassword"`
	CreatedAt       time.Time                  `db:"createdAt"`
	UpdatedAt       gonull.Nullable[time.Time] `db:"updatedAt"`
}







// ------------ PRODUCTS ------------
// Untuk Products model
type Products struct {
	Id            int                        `db:"id" json:"id"`
	Name          string                     `db:"name" json:"name" form:"name" binding:"required"`
	Price         int64                      `db:"price" json:"price" form:"price" binding:"required"`
	Image         *string                    `db:"image" json:"image" form:"image" binding:"required"`
	Description   *string                    `db:"description" json:"description" form:"description" binding:"required"`
	Discount      *int32                     `db:"discount" json:"discount" form:"discount" binding:"required"`
	IsRecommended *bool                      `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	Qty           int64                      `db:"qty" json:"qty" form:"qty" binding:"required"`
	IsActive      bool                       `db:"isActive" json:"isActive" form:"isActive"`
	CreatedAt     time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt     gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}
type ProductsNet struct {
	Id            int                        `db:"id" json:"id"`
	Name          string                     `db:"name" json:"name" form:"name"`
	Price         int64                      `db:"price" json:"price" form:"price"`
	Image         gonull.Nullable[string]    `db:"image" json:"image" form:"image"`
	Description   gonull.Nullable[string]    `db:"description" json:"description" form:"description"`
	Discount      int64                      `db:"discount" json:"discount" form:"discount"`
	IsRecommended gonull.Nullable[bool]      `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	Qty           int64                      `db:"qty" json:"qty" form:"qty"`
	IsActive      gonull.Nullable[bool]      `db:"isActive" json:"isActive" form:"isActive"`
	CreatedAt     time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt     gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}







// ------------ PROMO ------------
type Promo struct {
	Id            int                        `db:"id" json:"id"`
	Name          *string                    `db:"name" json:"name" form:"name" binding:"required"`
	Code          *string                    `db:"code" json:"code" form:"code" binding:"required"`
	Description   string                     `db:"description" form:"description" json:"description"`
	Percentage    *float32                   `db:"percentage" form:"percentage" json:"percentage" binding:"required"`
	IsExpired     bool                       `db:"isExpired" form:"isExpired" json:"isExpired"`
	MaximumPromo  *int32                     `db:"maximumPromo" form:"maximumPromo" json:"maximumPromo" binding:"required"`
	MinimumAmount *int32                     `db:"minimumAmount" form:"minimumAmount" json:"minimumAmount" binding:"required"`
	CreatedAt     time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt     gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}
type PromoNet struct {
	Id            int                        `db:"id" json:"id"`
	Name          string                     `db:"name" json:"name"`
	Code          string                     `db:"code" json:"code"`
	Description   gonull.Nullable[string]    `db:"description" json:"description"`
	Percentage    float32                    `db:"percentage" json:"percentage"`
	IsExpired     gonull.Nullable[bool]      `db:"isExpired" json:"isExpired"`
	MaximumPromo  int32                      `db:"maximumPromo" json:"maximumPromo"`
	MinimumAmount int32                      `db:"minimumAmount" json:"minimumAmount"`
	CreatedAt     time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt     gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}







// ------------ PRODUCT SIZE ------------
type Ps struct {
	Id              int                        `db:"id" json:"id"`
	Size            string                     `db:"size" json:"size"`
	AdditionalPrice int64                      `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice" binding:"required"`
	CreatedAt       time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt       gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}
type PsNet struct {
	Id              int                        `db:"id" json:"id"`
	Size            string                     `db:"size" json:"size"`
	AdditionalPrice int64                      `db:"additionalPrice" json:"additionalPrice"`
	CreatedAt       time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt       gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}







// ------------ PRODUCT VARIANT ------------
type Pv struct {
	Id              int                        `db:"id" json:"id"`
	Name            string                     `db:"name" json:"name" form:"name" binding:"required"`
	AdditionalPrice int64                      `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice" binding:"required"`
	CreatedAt       time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt       gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}
type PvNet struct {
	Id              int                        `db:"id" json:"id"`
	Name            string                     `db:"name" json:"name"`
	AdditionalPrice int64                      `db:"additionalPrice" json:"additionalPrice"`
	CreatedAt       time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt       gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}






// ------------ CATEGORIES ------------
type Categories struct{
	Id            int                        `db:"id" json:"id"`
	Name          string                     `db:"name" json:"name" form:"name" binding:"required"`
	CreatedAt     time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt     gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}







// ------------ TAGS ------------
type Tags struct {
	Id        int                        `db:"id" json:"id"`
	Name      string                     `db:"name" json:"name" form:"name"`
	CreatedAt time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}
