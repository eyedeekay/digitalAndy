package main

import (
        "encoding/binary"
	"golang.org/x/mobile/exp/f32"
)

type Point struct {
        X, Y float32
        Z float32
}

func NewPoint(x, y float32) *Point {
    p := new(Point)
    p.X = x
    p.Y = y
    p.Z = 0.0
    return p
}

func (p *Point) GLPoint() []byte {
        return f32.Bytes(binary.LittleEndian, p.X, p.Y, p.Z)
}

func NewTestTriangle() []byte {
        //a, b, c Point
        var a = NewPoint(0.0, 0.4)
        var b = NewPoint(0.0, 0.0)
        var c = NewPoint(0.4, 0.4)

        var test = a.GLPoint()
        test = append(test, b.GLPoint()...)
        test = append(test, c.GLPoint()...)
        return test
}
