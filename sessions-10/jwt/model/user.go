package model

type UserRegister struct {
	FullName string `json:"fullName" gorm:"not null;" valid:"required-Your Full Name is required"`
	Email    string `json:"email" gorm:"not null; uniqueIndex" valid:"required-Your Email is required, email-invalid email format"`
	Password string `json:"password" gorm:"not null;" valid:"required-Your is required, minstringlength(8) - Password has to have a minimun length 6 characters"`
}

type UserModel struct {
	Id       uint   `json:"id"`
	FullName string `json:"fullName" gorm:"not null;" valid:"required-Your Full Name is required"`
	Email    string `json:"email" gorm:"not null; uniqueIndex" valid:"required-Your Email is required, email-invalid email format"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"not null; uniqueIndex" valid:"required-Your Email is required, email-invalid email format"`
	Password string `json:"password" gorm:"not null;" valid:"required-Your is required, minstringlength(8) - Password has to have a minimun length 6 characters"`
}
