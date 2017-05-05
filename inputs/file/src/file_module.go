/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
    "fmt"
)

const (
    FILE_MODULE = INPUT_MODULE|0x03000000
    FILE_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Files struct {
    *Module
}

func (r *Files) Init(o *Option) int {
    fmt.Println("InputFile init")
    return Ok
}

func (r *Files) Main(cfg *Configure) int {
    fmt.Println("InputFile main")
    return Ok
}

func (r *Files) Exit() int {
    fmt.Println("InputFile exit")
    return Ok
}

func (r *Files) Type() *Module {
    return r.Self()
}

var file = String{ len("file"), "file" }
var fileCommands = []Command{

    { file,
      FILE_CONFIG,
      fileBlock,
      0,
      0,
      nil },

    NilCommand,
}

func fileBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := FILE_CONFIG|CONFIG_VALUE
    Block(c, Modulers, FILE_MODULE, flag)

    return Ok
}

var fileModule = &Files{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        fileCommands,
        INPUT_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, fileModule)
}
