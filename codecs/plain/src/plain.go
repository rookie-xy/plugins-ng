/*
 * Copyright (C) 2017 Meng Shi
 */

package plain

import (
    . "github.com/rookie-xy/worker/types"
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

func (r *Plain) Encode(in interface{}) (interface{}, error) {
    fmt.Printf("name: %s, name: %s, plain: %X\n", r.name, &r.name, &r)
    return nil, nil
}

func (r *Plain) Decode(in []byte) (interface{}, error) {
    return nil, nil
}

func (r *Plain) Type(name string) int {
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = append(Codecs, plain)
}
