package multiline

import (
    . "github.com/rookie-xy/main/types"
    "fmt"
)

type Multiline struct {
    name      string

    match     string
    previous  string
    next      string
}

func NewMultiline() *Multiline {
    return &Multiline{}
}

var multiline = &Multiline{
    name: "multiline",

    match: "{^",
    previous: "what",
    next: "where",
}

func (r *Multiline) Init(configure interface{}) int {
    //fmt.Println(configure)
    return Ok
}

func (r *Multiline) Encode() int {
    fmt.Println("multiline")
    return Ok
}

func (r *Multiline) Decode() int {
    return Ok
}

func (r *Multiline) Type(name string) int {
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = Setup(Codecs, multiline)
}
