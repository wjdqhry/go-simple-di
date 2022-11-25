package di_test

import (
	"github.com/stretchr/testify/assert"
	"go-simple-di/di"
	"testing"
)

func TestDiWithReflection(t *testing.T) {
	type AStruct struct {
		name string
	}

	di.Inject(&AStruct{name: "Hi"})
	aStruct := di.Resolve[*AStruct]()
	assert.Equal(t, "Hi", aStruct.name)
}

func TestDiWithNamed(t *testing.T) {
	type AStruct struct {
		name string
	}

	di.InjectNamed("A", &AStruct{name: "Hi"})
	aStruct := di.ResolveNamed[*AStruct]("A")
	assert.Equal(t, "Hi", aStruct.name)
}
