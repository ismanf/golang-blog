package session

import (
	"reflect"
	"time"
	"github.com/ismayilmalik/golang-blog/infrastructure/cache"
	"github.com/nu7hatch/gouuid"
)

func newToken() string {
	token, _ := uuid.NewV4()
	return token.String()
}

func New(userInfo string) (string, error) {
	token := newToken()
	err := cache.Set(token, userInfo, 10 * time.Minute)
	if err != nil {
		return "", err
	}

	return token, nil
}

func IsValid(token string) (bool, error) {
	_, err := cache.Get(token)

	if reflect.TypeOf(err) == reflect.TypeOf(cache.KeyNotExistsErorr{}) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}


