/*
 * Copyright (C) 2017 Meng Shi
 */

package stdin

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

type Stdin struct {
    *Module
    *File

     status   bool
     channel  string
}

func NewStdin() *Stdin {
    return &Stdin{}
}

type StdinContext struct {
    *Context
}

var stdinPlugin = String{ len("stdin_plugin"), "stdin_plugin" }
var stdinContext = &StdinContext{
    Context: &Context{
        Name: stdinPlugin,
    },
}

func (r *StdinContext) Create() unsafe.Pointer {
    stdin := NewStdin()
    if stdin == nil {
        return nil
    }

    stdin.status = false
    stdin.channel = "mengshi"

    return unsafe.Pointer(stdin)
}

func (r *StdinContext) Insert(p *unsafe.Pointer) int {
    this := (*Stdin)(unsafe.Pointer(uintptr(*p)))
    if this == nil {
        fmt.Println("stdin context")
        return Error
    }

    fmt.Println(this.channel)

    return Ok
}

func (r *StdinContext) Contexts() *Context {
    return r.Get()
}

var (
    status = String{ len("status"), "status" }
    channel = String{ len("channel"), "channel" }
    mystdin Stdin
)

var stdinCommands = []Command{

    { status,
      STDIN_CONFIG|CONFIG_VALUE,
      SetFlag,
      0,
      unsafe.Offsetof(mystdin.status),
      nil },

    { channel,
      STDIN_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(mystdin.channel),
      nil },

    NilCommand,
}

func (r *Stdin) Init(o *Option) int {
    ctx := r.Context.Contexts()
    fmt.Println(len(ctx.Data))

    for _, v := range ctx.GetDatas() {
        if v != nil {
            this := (*Stdin)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                fmt.Println("stdin context")
                return Error
            }

            fmt.Println(this.channel)
        } else {
            break
        }
    }

    fmt.Println("Stdin init")
    return Ok
}

func (r *Stdin) Main(cfg *Configure) int {
    fmt.Println("Stdin main")
    return Ok
}

func (r *Stdin) Exit() int {
    fmt.Println("Stdin exit")
    return Ok
}

func (r *Stdin) Type() *Module {
    return r.Self()
}

var stdinModule = &Stdin{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        stdinContext,
        stdinCommands,
        STDIN_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, stdinModule)
}