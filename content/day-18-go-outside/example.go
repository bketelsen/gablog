package main

import (
	"github.com/tHinqa/outside"
	"runtime"
)

type Wand struct{}
type Bool bool

var New func() *Wand
var Destroy func(m *Wand) //*Wand
var Read func(m *Wand, filename string) (Bool, error)
var Sketch func(m *Wand, radius, sigma, angle float64) (Bool, error)
var Equalize func(m *Wand) (Bool, error)
var Write func(m *Wand, filename string) (Bool, error)

func (ok Bool) Error(e error) (Bool, error) {
	if ok {
		return ok, nil
	}
	panic("Failed with: " + e.Error())
}

var allApis = outside.Apis{
	{"NewMagickWand", &New},
	{"DestroyMagickWand", &Destroy},
	{"MagickReadImage", &Read},
	{"MagickSketchImage", &Sketch},
	{"MagickEqualizeImage", &Equalize},
	{"MagickWriteImage", &Write},
}

func init() {
	var dll string
	if runtime.GOOS == "windows" {
		dll = "CORE_RL_wand_.dll"
	}
	if runtime.GOOS == "linux" {
		dll = "libMagickWand.so.5"
	}
	outside.AddDllApis(dll, false, allApis)
}

func (m *Wand) Read(filename string) *Wand {
	Read(m, filename)
	return m
}
func (m *Wand) Sketch(radius, sigma, angle float64) *Wand {
	Sketch(m, radius, sigma, angle)
	return m
}
func (m *Wand) Equalize() *Wand {
	Equalize(m)
	return m
}
func (m *Wand) Write(filename string) *Wand {
	Write(m, filename)
	return m
}

func main() {
	m := New()
	defer Destroy(m)
	m.Read("i.png").Sketch(0, 30, 60).Equalize().Write("o.jpeg")
}
