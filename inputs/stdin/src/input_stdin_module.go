/*
 * Copyright (C) 2017 Meng Shi
 */

package stdin

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
    . "github.com/rookie-xy/main/modules"
    "fmt"
)

const (
    STDIN_MODULE = INPUT_MODULE|0x01000000
    STDIN_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type InputStdin struct {
    *Module
}

func (r *InputStdin) Init(o *Option) int {
    fmt.Println("InputStdin init")
    return Ok
}

func (r *InputStdin) Main(cfg *Configure) int {
    fmt.Println("InputStdin main")
    return Ok
}

func (r *InputStdin) Exit() int {
    fmt.Println("InputStdin exit")
    return Ok
}

func (r *InputStdin) Type() *Module {
    return r.Self()
}

/*
type InputStdinContext struct {
    *Context
}

var stdinModule = String{ len("stdin_module"), "stdin_module" }
var inputStdinContext = &InputStdinContext {
    Context: &Context{
        Name: stdinModule,
    },
}

func (r *InputStdinContext) Create() unsafe.Pointer {
    return Ok
}

func (r *InputStdinContext) Insert(p *unsafe.Pointer) int {
    return Ok
}

func (r *InputStdinContext) Contexts() *Context {
    return r.Get()
}
*/

var stdinInput = String{ len("stdin_input"), "stdin_input" }
var inputStdinCommands = []Command{

    { stdinInput,
      STDIN_CONFIG,
      stdinBlock,
      0,
      0,
      nil },

    NilCommand,
}

func stdinBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := STDIN_CONFIG|CONFIG_VALUE
    Block(c, Modulers, STDIN_MODULE, flag)

    return Ok
}

var inputStdinModule = &InputStdin{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        //inputStdinContext,
        nil,
        inputStdinCommands,
        INPUT_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, inputStdinModule)
}