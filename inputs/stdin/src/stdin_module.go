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

type Stdin struct {
    *Module
}

func (r *Stdin) Init(o *Option) int {
    fmt.Println("InputStdin init")
    return Ok
}

func (r *Stdin) Main(cfg *Configure) int {
    fmt.Println("InputStdin main")
    return Ok
}

func (r *Stdin) Exit() int {
    fmt.Println("InputStdin exit")
    return Ok
}

func (r *Stdin) Type() *Module {
    return r.Self()
}

var stdin = String{ len("stdin"), "stdin" }
var stdinCommands = []Command{

    { stdin,
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

var stdinModule = &Stdin{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        stdinCommands,
        INPUT_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, stdinModule)
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
