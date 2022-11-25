package di

import (
	"fmt"
	"log"
	"reflect"
)

var container = map[string]Injectable{}

type Injectable interface{}

func Inject(module Injectable) {
	moduleName := getModuleName(module)
	inject(moduleName, module)
}

func InjectNamed(moduleName string, module Injectable) {
	inject(moduleName, module)
}

func inject(name string, module Injectable) {
	log.Println("Inject ModuleName: " + name)
	if container[name] != nil {
		panic(fmt.Sprintf("%s Already Exists", name))
	}

	container[name] = module
}

func Resolve[T any]() T {
	var empty T
	moduleName := getModuleName(empty)
	return resolve(moduleName).(T)
}

func ResolveNamed[T any](moduleName string) T {
	return resolve(moduleName).(T)
}

func resolve(name string) Injectable {
	log.Println("Resolve ModuleName: " + name)
	if m := container[name]; m == nil {
		panic(fmt.Sprintf("%s Not Exists", name))
	} else {
		return m
	}
}

func getModuleName(module Injectable) string {
	return reflect.TypeOf(module).String()
}
