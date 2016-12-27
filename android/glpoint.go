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

func NewOffsetPoint(x,y, w, h float32) *Point{
        p := new(Point)
        p.X = x+w
        p.Y = y+h
        p.Z = 0.0
        return p
}

func AddAPoint(XY Point, w, h float32) *Point{
        p := new(Point)
        p.X = XY.X+w
        p.Y = XY.Y+h
        p.Z = 0.0
        return p
}
func (p *Point) AddAPoint(w, h float32) *Point{
        q := new(Point)
        q.X = p.X+w
        q.Y = p.Y+h
        q.Z = 0.0
        return p
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

func ANewPoint(appendto []byte, pointb Point) []byte {
        var rvalue = appendto
        rvalue = append(rvalue, pointb.GLPoint()...)
        return rvalue
}

