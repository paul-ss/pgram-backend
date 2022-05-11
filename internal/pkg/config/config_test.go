package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestInit(t *testing.T) {
	dir, err := os.Getwd()
	assert.Nil(t, err)

	c := newConfig(dir + "/../../../configs")
	assert.NotNil(t, c)
	fmt.Println(*c)
}

func TestCheckUnsetFields(t *testing.T) {
	dir, err := os.Getwd()
	assert.Nil(t, err)

	c := newConfig(dir + "/../../../configs")
	assert.NotNil(t, c)
	fmt.Println(*c)

	checkUnsetFields(c)
}

func TestGetAllTags(t *testing.T) {
	exp := []string{"f1", "Field2", "f3.n", "f3.NestField2"}
	sort.Strings(exp)

	s := struct {
		Field1 int `mapstructure:"f1"`
		Field2 string
		Field3 struct {
			NestField  int `mapstructure:"n"`
			NestField2 int
		} `mapstructure:"f3"`
	}{}

	got := getAllStructTags("", "", reflect.TypeOf(s))
	sort.Strings(got)

	assert.Equal(t, exp, got)
}
