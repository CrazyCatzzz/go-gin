package util

import "golang.org/x/crypto/bcrypt"

// CheckPassword 密码校验
func CheckPassword(old, new string) bool {
	return bcrypt.CompareHashAndPassword([]byte(old), []byte(new)) == nil
}

// GeneratePassword 密码生成
func GeneratePassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}
