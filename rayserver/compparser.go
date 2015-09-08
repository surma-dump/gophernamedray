package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"math"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"

	"github.com/surma-dump/gophernamedray/gnr"
)

type composition struct {
	Name string
	Func otto.Value
}

func (j *job) runCompScript(irs [][]*gnr.InteractionResult) {
	vm := otto.New()

	injectRuntime(vm)

	if err := vm.Set("interactions", irs); err != nil {
		panic(err)
	}
	vm.Set("width", Width)
	vm.Set("height", Height)

	compositions := make([]composition, 0, 5)
	vm.Set("composition", func(call otto.FunctionCall) otto.Value {
		name, err := call.Argument(0).ToString()
		if err != nil {
			panic("First argument to composition() must be a string")
		}
		cb := call.Argument(1)
		if !cb.IsFunction() {
			panic("Second argument to composition() must be a function(x,y)")
		}
		compositions = append(compositions, composition{name, cb})
		return otto.NullValue()
	})

	if _, err := vm.Run(j.Composite); err != nil {
		panic(err)
	}

	for _, comp := range compositions {
		img := image.NewRGBA(image.Rect(0, 0, Width, Height))
		for y := 0; y < Height; y++ {
			for x := 0; x < Width; x++ {
				v, err := comp.Func.Call(otto.NullValue(), x, y)
				if err != nil {
					panic(err)
				}
				c, err := v.Export()
				if err != nil {
					panic(err)
				}
				img.Set(x, y, mapToColor(c.(map[string]interface{})))
			}
		}
		buf := &bytes.Buffer{}
		if err := png.Encode(buf, img); err != nil {
			panic(err)
		}
		j.Compositions[comp.Name] = buf.Bytes()
	}
}

var colorF = gnr.LerpCap(0, 1, 0, float64(0xff))

func mapToColor(o map[string]interface{}) color.Color {
	var c color.RGBA

	// Because JavaScript
	switch x := o["r"].(type) {
	case int64:
		c.R = uint8(colorF(float64(x)))
	case float64:
		c.R = uint8(colorF(x))
	}
	switch x := o["g"].(type) {
	case int64:
		c.G = uint8(colorF(float64(x)))
	case float64:
		c.G = uint8(colorF(x))
	}
	switch x := o["b"].(type) {
	case int64:
		c.B = uint8(colorF(float64(x)))
	case float64:
		c.B = uint8(colorF(x))
	}
	c.A = 0xff
	return c
}

func injectRuntime(vm *otto.Otto) {
	var err error
	err = vm.Set("sqrt", func(call otto.FunctionCall) otto.Value {
		f, err := call.Argument(0).ToFloat()
		if err != nil {
			panic(err)
		}
		r, err := otto.ToValue(math.Sqrt(f))
		if err != nil {
			panic(err)
		}
		return r
	})
	_, err = vm.Run(`
	gnr = {
		lerp: function(minIn, maxIn, minOut, maxOut) {
			return function(d) {
				return (d - minIn)/(maxIn - minIn)*(maxOut - minOut) + minOut;
			}
		},
		lerpCap: function(minIn, maxIn, minOut, maxOut) {
			var f = gnr.lerp(minIn, maxIn, minOut, maxOut);
			return function(d) {
				if(d < minIn) {
					d = minIn;
				}
				if(d > maxIn) {
					d = maxIn;
				}
				return f(d);
			}
		},
		vlerp: function(minIn, maxIn, minOut, maxOut) {
			var xLerp = gnr.lerp(minIn, maxIn, minOut.X, maxOut.X);
			var yLerp = gnr.lerp(minIn, maxIn, minOut.Y, maxOut.Y);
			var zLerp = gnr.lerp(minIn, maxIn, minOut.Z, maxOut.Z);
			return function(d) {
				return {
					X: xLerp(d),
					Y: yLerp(d),
					Z: zLerp(d)
				};
			};
		},
		vlerpCap: function(minIn, maxIn, minOut, maxOut) {
			var xLerp = gnr.lerpCap(minIn, maxIn, minOut.X, maxOut.X);
			var yLerp = gnr.lerpCap(minIn, maxIn, minOut.Y, maxOut.Y);
			var zLerp = gnr.lerpCap(minIn, maxIn, minOut.Z, maxOut.Z);
			return function(d) {
				return {
					X: xLerp(d),
					Y: yLerp(d),
					Z: zLerp(d)
				};
			};
		},
		vector2Color: function(v) {
			return {
				r: v.X,
				g: v.Y,
				b: v.Z
			};
		},
		vector: {
		  product: function(v1, v2) {
				return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z;
			},
			length: function(v) {
				return sqrt(gnr.vector.product(v, v));
			}
	  }
	}
	`)
	if err != nil {
		panic(err)
	}
}
