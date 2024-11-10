package main

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/sirupsen/logrus"
)

type User struct {
	isAdmin  bool
	Username string
	Password string
}

func (u *User) ChangeUsername(username string) bool {
	u.Username = username

	return true
}

func (u *User) ChangePassword(password string) bool {
	logrus.Infof("Changing password for user %s", u.Username)

	h := sha256.New()
	h.Write([]byte(password))

	u.Password = hex.EncodeToString(h.Sum(nil))

	return true
}

func (u *User) CheckPassword(password string) bool {
	h := sha256.New()
	h.Write([]byte(password))

	return u.Password == hex.EncodeToString(h.Sum(nil))
}
