package status

import (
	"log"
	"testing"

	"github.com/pkg/errors"
)

const (
	GetUserListError = Code(28001)
)

func TestError(t *testing.T) {
	err := Error(GetUserListError)
	log.Println(err)
	status := FromError(err)
	log.Println(status.code, status.message)
}

func TestUnknownError(t *testing.T) {
	err := Error(28000)
	log.Println(err)

	status := FromError(err)
	log.Println(status.code, status.message)
}

func TestWrap(t *testing.T) {
	err := Wrap(GetUserListError, errors.New("connect to db error"))
	status := FromError(err)
	log.Println(status.Error())
}

func init() {
	Register(GetUserListError)
}
