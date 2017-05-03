/*
 * Copyright (C) 2017 Meng Shi
 */

package plain

import (
    . "github.com/rookie-xy/main/types"
    "fmt"
)

type Plain struct {
    name string
}

func NewPlain() *Plain {
    return &Plain{}
}

var plain = &Plain{
    name: "plain",
}

func (r *Plain) New() Codec {
    plain := NewPlain()

    plain.name = "new_plain"

    return plain
}

func (r *Plain) Init(configure interface{}) int {
    //fmt.Println(configure)
    return Ok
}

func (r *Plain) Encode() int {
    fmt.Printf("name: %s, name: %s, plain: %X\n", r.name, &r.name, &r)
    return Ok
}

func (r *Plain) Decode() int {
    return Ok
}

func (r *Plain) Type(name string) int {
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = Setup(Codecs, plain)
}
