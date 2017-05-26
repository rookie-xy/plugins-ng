package multiline

import (
    . "github.com/rookie-xy/worker/types"
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

func (r *Multiline) New() Codec {
    multiline := NewMultiline()

    multiline.name = "new_multiline"
    multiline.match = r.match
    multiline.previous = r.previous
    multiline.next = r.next

    return multiline
}

func (r *Multiline) Init(configure interface{}) int {
    //fmt.Println(configure)
    return Ok
}

func (r *Multiline) Encode(in interface{}) (interface{}, error) {
    fmt.Println(r.name)
    return nil, nil
}

func (r *Multiline) Decode(in []byte) (interface{}, error) {
    return nil, nil
}

func (r *Multiline) Type(name string) int {
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = append(Codecs, multiline)
}
