/*
 * Copyright (C) 2017 Meng Shi
 */

package line

import (
    . "github.com/rookie-xy/worker/types"
    "fmt"
)

type Line struct {
    name string
}

func NewLine() *Line {
    return &Line{}
}

var line = &Line{
    name: "Line",
}

func (r *Line) New() Codec {
    line := NewLine()

    line.name = "new_line"

    return line
}

func (r *Line) Init(configure interface{}) int {
    //fmt.Println(configure)
    return Ok
}

func (r *Line) Encode(in interface{}) (interface{}, error) {
    fmt.Printf("name: %s, name: %s, line: %X\n", r.name, &r.name, &r)
    return nil, nil
}

func (r *Line) Decode(in []byte) (interface{}, error) {
    return nil, nil
}

func (r *Line) Type(name string) int {
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = append(Codecs, line)
}
