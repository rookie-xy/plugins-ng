package multiline

import (
      "unsafe"
      "fmt"
    . "github.com/rookie-xy/main/types"
)

type CodecMultiline struct {
    *Module

     match     string
     previous  string
     next      string
}

func NewCodecMultiline() *CodecMultiline {
    return &CodecMultiline{}
}

type CodecMultilineContext struct {
    *Context
}

var multilineCodec = String{ len("multiline_codec"), "multiline_codec" }
var codecMultilineContext = &CodecMultilineContext{
    Context: &Context{
        Name: multilineCodec,
    },
}

func (r *CodecMultilineContext) Create() unsafe.Pointer {
    multiline := NewCodecMultiline()
    if multiline == nil {
        return nil
    }

    multiline.match = "abc"
    multiline.previous = "a"
    multiline.next = "c"

    return unsafe.Pointer(multiline)
}

func (r *CodecMultilineContext) Contexts() *Context {
    return r.Get()
}

var (
    match    = String{ len("match"), "match" }
    previous = String{ len("previous"), "previous" }
    next     = String{ len("next"), "next" }
    codecMultiline  CodecMultiline
)

var codecMultilineCommands = []Command{

    { match,
      MULTILINE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(codecMultiline.match),
      nil },

    { previous,
      MULTILINE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(codecMultiline.previous),
      nil },

    { next,
      MULTILINE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(codecMultiline.next),
      nil },

    NilCommand,
}

var codecMultilineModule = &CodecMultiline{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        codecMultilineContext,
        codecMultilineCommands,
        MULTILINE_MODULE,
    },
}

func (r *CodecMultiline) Init(o *Option) int {
    context := r.Context.Contexts()

    for _, v := range context.Data {
        if v != nil {
            this := (*CodecMultiline)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.match, this.previous, this.next)
        } else {
            break
        }
    }

    return Ok
}

func (r *CodecMultiline) Main(cfg *Configure) int {
    fmt.Println("CodecMultiline main")
    return Ok
}

func (r *CodecMultiline) Exit() int {
    fmt.Println("CodecMultiline exit")
    return Ok
}

func (r *CodecMultiline) Type() *Module {
    return r.Self()
}

func init() {
    Modulers = Load(Modulers, codecMultilineModule)
}
