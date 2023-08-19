package test

import (
	"testing"

	"api.aifuxi.cool/models"
)

func Test(t *testing.T) {
	tag := new(models.Tag)

	println("tag: %v\n", tag)

	println(tag == nil)
}
