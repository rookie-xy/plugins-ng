/*
 * Copyright (C) 2017 Meng Shi
 */

package stdin

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

type InputStdin struct {
    *Module
}

type MyStdin struct {
    status   bool
    channel  string
}

func NewMyStdin() *MyStdin {
    return &MyStdin{}
}

type InputStdinContext struct {
    *Context
}

var inputStdin = String{ len("input_stdin"), "input_stdin" }
var inputStdinContext = &InputStdinContext{
    Context: &Context{
        Name: inputStdin,
    },
}

func (r *InputStdinContext) Create() unsafe.Pointer {
    stdin := NewMyStdin()
    if stdin == nil {
        return nil
    }

    stdin.status = false
    stdin.channel = "mengshi"

    return unsafe.Pointer(stdin)
}

func (r *InputStdinContext) Contexts() *Context {
    return r.Get()
}

var (
    status = String{ len("status"), "status" }
    channel = String{ len("channel"), "channel" }
    myStdin MyStdin
)

var inputStdinCommands = []Command{

    { status,
      STDIN_CONFIG|CONFIG_VALUE,
      SetFlag,
      0,
      unsafe.Offsetof(myStdin.status),
      nil },

    { channel,
      STDIN_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(myStdin.channel),
      nil },

    NilCommand,
}

func (r *InputStdin) Init(o *Option) int {
    context := r.Context.Contexts()

    for _, v := range context.Data {
        if v != nil {
            this := (*MyStdin)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.channel)
        } else {
            break
        }
    }

    return Ok
}

func (r *InputStdin) Main(cfg *Configure) int {
    fmt.Println("Stdin main")
    return Ok
}

func (r *InputStdin) Exit() int {
    fmt.Println("Stdin exit")
    return Ok
}

func (r *InputStdin) Type() *Module {
    return r.Self()
}

var inputStdinModule = &InputStdin{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        inputStdinContext,
        inputStdinCommands,
        STDIN_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, inputStdinModule)
}