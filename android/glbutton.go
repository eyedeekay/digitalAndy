package main

//import (
//        "encoding/binary"
//        "golang.org/x/mobile/exp/app/debug"
//	"golang.org/x/mobile/exp/f32"
//	"golang.org/x/mobile/exp/gl/glutil"
//	"golang.org/x/mobile/gl"
//)

type Button struct {
        TL, TR, BL, BR Point
        Text string
//        col    gl.Uniform
}


/* A ReactRect is a button which only changes it's own properties when triggered
by an event. It might not stick around but changing the name will force me to
go over all the related code before it'll compile again. */
func NewReactRect(x, y float32, w, h float32) *Button {
        b := new(Button)
        b.TL = *NewPoint(x,y)
        b.TR = *NewOffsetPoint(x, y, w, 0)
        b.BL = *NewOffsetPoint(x, y, 0, h)
        b.BR = *NewOffsetPoint(x, y, w, h)
        return b
}

func NewReactRectFromPoint(xy Point, w, h float32) *Button{
        b := new(Button)
        b.TL = xy
        b.TR = *AddAPoint(b.TL, w, 0)
        b.BL = *AddAPoint(b.TL, 0, h)
        b.BR = *AddAPoint(b.TL, w, h)
        return b
}

func (b *Button) GLButton() []byte{
        var rvalue = b.TL.GLPoint()
        rvalue = ANewPoint(rvalue, b.TR)
        rvalue = ANewPoint(rvalue, b.BL)
        rvalue = ANewPoint(rvalue, b.BR)
        return rvalue
}
