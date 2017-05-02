/*
 * Copyright (C) 2017 Meng Shi
 */

package multiline
/*
import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
    . "github.com/rookie-xy/main/modules"
"fmt"
)

const (
    MULTILINE_MODULE = CODEC_MODULE|0x01000000
    MULTILINE_CONFIG = USER_CONFIG|CONFIG_MAP
)

type Multiline struct {
    *Module
}

func (r *Multiline) Init(o *Option) int {
    fmt.Println("Multiline init")
    return Ok
}

func (r *Multiline) Main(cfg *Configure) int {
    fmt.Println("Multiline main")
    return Ok
}

func (r *Multiline) Exit() int {
    fmt.Println("Multiline exit")
    return Ok
}

func (r *Multiline) Type() *Module {
    return r.Self()
}

var	multiline = String{ len("multiline"), "multiline" }
var multilineCommands = []Command{

    { multiline,
      MULTILINE_CONFIG,
      multilineBlock,
      0,
      0,
      nil },

    NilCommand,
}

func multilineBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := MULTILINE_CONFIG|CONFIG_VALUE
    Block(c, Modulers, MULTILINE_MODULE, flag)

    return Ok
}

var multilineModule = &Multiline{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        multilineCommands,
        CODEC_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, multilineModule)
}
*/