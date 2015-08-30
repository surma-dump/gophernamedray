package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"

	"github.com/surma-dump/gophernamedray/gnr"
	"github.com/surma-dump/gophernamedray/gnr/object"
)

type refMap map[string]interface{}

func (rm refMap) MustResolve(ref otto.Value) interface{} {
	key, err := ref.ToString()
	if err != nil {
		panic(err)
	}
	v, ok := rm[key]
	if !ok {
		panic(fmt.Sprintf("Invalid reference: %s", key))
	}
	return v
}

func (j *job) runSceneScript() (*gnr.Scene, error) {
	vm := otto.New()

	rm := refMap{}
	rm["vm"] = vm
	for name, fn := range sceneEnv {
		f := fn(rm)
		vm.Set(name, func(call otto.FunctionCall) otto.Value {
			id := NewUUID()
			rm[id] = f(call)
			ref, _ := vm.ToValue(id)
			return ref
		})
	}
	ref, err := vm.Run(j.Scene)
	if err != nil {
		return nil, err
	}
	return rm.MustResolve(ref).(*gnr.Scene), nil
}

var sceneEnv = map[string]func(refMap) func(otto.FunctionCall) interface{}{
	"Scene": func(rm refMap) func(otto.FunctionCall) interface{} {
		return func(call otto.FunctionCall) interface{} {
			return &gnr.Scene{
				Camera: rm.MustResolve(call.Argument(0)).(gnr.Camera),
				Object: rm.MustResolve(call.Argument(1)).(gnr.Object),
			}
		}
	},
	"PlanarCamera": func(rm refMap) func(otto.FunctionCall) interface{} {
		return func(call otto.FunctionCall) interface{} {
			angle, err := call.Argument(2).ToFloat()
			if err != nil {
				panic(err)
			}
			return &gnr.PlanarCamera{
				Position:      rm.MustResolve(call.Argument(0)).(*gnr.Vector3f),
				ViewDirection: rm.MustResolve(call.Argument(1)).(*gnr.Vector3f),
				UpDirection:   &gnr.Vector3f{0, 1, 0},
				PixelWidth:    Width,
				PixelHeight:   Height,
				VirtualWidth:  1,
				VirtualHeight: 1,
				Angle:         angle,
			}
		}
	},
	"Vector": func(rm refMap) func(otto.FunctionCall) interface{} {
		return func(call otto.FunctionCall) interface{} {
			x, err := call.Argument(0).ToFloat()
			if err != nil {
				panic(err)
			}
			y, err := call.Argument(1).ToFloat()
			if err != nil {
				panic(err)
			}
			z, err := call.Argument(2).ToFloat()
			if err != nil {
				panic(err)
			}

			return &gnr.Vector3f{x, y, z}
		}
	},
	"AxisAlignedBox": func(rm refMap) func(otto.FunctionCall) interface{} {
		return func(call otto.FunctionCall) interface{} {
			return &object.AxisAlignedBox{
				Min: rm.MustResolve(call.Argument(0)).(*gnr.Vector3f),
				Max: rm.MustResolve(call.Argument(1)).(*gnr.Vector3f),
			}
		}
	},
}
