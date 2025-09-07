package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Constructor func() interface{}

type Container struct {
	container map[string]Constructor
}

func NewContainer() *Container {
	return &Container{
		container: make(map[string]Constructor),
	}
}

func (c *Container) RegisterType(name string, constructor Constructor) {
	if _, ok := c.container[name]; ok {
		panic(fmt.Sprintf("container with name %s already exists", name))
	}
	c.container[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	res, ok := c.container[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s not found in conraitner", name))
	}
	return res(), nil
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})

	userService1, err := container.Resolve("UserService")
	assert.NoError(t, err)
	userService2, err := container.Resolve("UserService")
	assert.NoError(t, err)

	u1 := userService1.(*UserService)
	u2 := userService2.(*UserService)
	assert.False(t, u1 == u2)

	messageService, err := container.Resolve("MessageService")
	assert.NoError(t, err)
	assert.NotNil(t, messageService)

	paymentService, err := container.Resolve("PaymentService")
	assert.Error(t, err)
	assert.Nil(t, paymentService)
}
