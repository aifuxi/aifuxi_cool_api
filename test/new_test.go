package test

import (
	"testing"

	"github.com/aifuxi/aifuxi_cool_api/models"
)

func Test(t *testing.T) {
	tag := new(models.Tag)

	println("tag: %v\n", tag)

	println(tag == nil)
}
