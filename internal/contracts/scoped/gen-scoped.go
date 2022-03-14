// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package scoped

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeIScoped used when your service claims to implement IScoped
var ReflectTypeIScoped = di.GetInterfaceReflectType((*IScoped)(nil))

// AddSingletonIScoped adds a type that implements IScoped
func AddSingletonIScoped(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddSingletonWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddSingletonIScopedByObj adds a prebuilt obj
func AddSingletonIScopedByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIScopedByFunc adds a type by a custom func
func AddSingletonIScopedByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIScoped adds a type that implements IScoped
func AddTransientIScoped(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIScopedByFunc adds a type by a custom func
func AddTransientIScopedByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIScoped adds a type that implements IScoped
func AddScopedIScoped(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIScopedByFunc adds a type by a custom func
func AddScopedIScopedByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIScoped)
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// RemoveAllIScoped removes all IScoped from the DI
func RemoveAllIScoped(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIScoped)
}

// GetIScopedFromContainer alternative to SafeGetIScopedFromContainer but panics of object is not present
func GetIScopedFromContainer(ctn di.Container) IScoped {
	return ctn.GetByType(ReflectTypeIScoped).(IScoped)
}

// GetManyIScopedFromContainer alternative to SafeGetManyIScopedFromContainer but panics of object is not present
func GetManyIScopedFromContainer(ctn di.Container) []IScoped {
	objs := ctn.GetManyByType(ReflectTypeIScoped)
	var results []IScoped
	for _, obj := range objs {
		results = append(results, obj.(IScoped))
	}
	return results
}

// SafeGetIScopedFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIScopedFromContainer(ctn di.Container) (IScoped, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIScoped)
	if err != nil {
		return nil, err
	}
	return obj.(IScoped), nil
}

// SafeGetManyIScopedFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIScopedFromContainer(ctn di.Container) ([]IScoped, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIScoped)
	if err != nil {
		return nil, err
	}
	var results []IScoped
	for _, obj := range objs {
		results = append(results, obj.(IScoped))
	}
	return results, nil
}