package service

type AuthService interface {
	IsDuplicateEmail(email string) bool
}
