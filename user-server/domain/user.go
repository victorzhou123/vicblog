package domain

import "github.com/victorzhou123/vicblog/common/domain/primitive"

type User struct {
	Username primitive.Username
	Email    primitive.Email
}
