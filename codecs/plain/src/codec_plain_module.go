/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
      "fmt"
    . "github.com/rookie-xy/main/types"
    . "github.com/rookie-xy/main/modules"
)

type CodecPlain struct {
    *Module
     plain  string
}

func NewCodecPlain() *CodecPlain {
    return &CodecPlain{}
}

func (r *CodecPlain) Init(o *Option) int {
    context := r.Context.Contexts()

    for _, v := range context.Data {
        if v != nil {
            this := (*CodecPlain)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.plain)
        } else {
            break
        }
    }

    fmt.Println("codec plain init")

    return Ok
}

func (r *CodecPlain) Main(c *Configure) int {
    fmt.Println("codec plain main")
    return Ok
}

func (r *CodecPlain) Exit() int {
    fmt.Println("codec plain exit")
    return Ok
}

func (r *CodecPlain) Type() *Module {
    return r.Self()
}

type CodecPlainContext struct {
    *Context
}

var plainCodec = String{ len("plain_codec"), "plain_codec" }
var codecPlainContext = &CodecPlainContext{
    Context: &Context{
        Name: plainCodec,
    },
}

func (r *CodecPlainContext) Create() unsafe.Pointer {
    codecPlain := NewCodecPlain()
    if codecPlain == nil {
        return nil
    }

    codecPlain.plain = "data"

    return unsafe.Pointer(codecPlain)
}

func (r *CodecPlainContext) Contexts() *Context {
    return r.Get()
}

var (
    plain  = String{ len("plain"), "plain" }
    codecPlain CodecPlain
)

var codecPlainCommands = []Command{

    { plain,
      USER_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(codecPlain.plain),
      nil },

    NilCommand,
}

var codecPlainModule = &CodecPlain{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        codecPlainContext,
        codecPlainCommands,
        CODEC_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, codecPlainModule)
}
