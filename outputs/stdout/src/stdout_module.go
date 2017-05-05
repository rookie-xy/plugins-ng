/*
 * Copyright (C) 2017 Meng Shi
 */

package stdout

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
)

const (
    STDOUT_MODULE = OUTPUT_MODULE|0x01000000
    STDOUT_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Stdout struct{
    *Module
}
/*
func (r *Stdout) Init(o *Option) int {
    fmt.Println("stdout init")
    return Ok
}

func (r *Stdout) Main(cfg *Configure) int {
    fmt.Println("stdout main")
    return Ok
}

func (r *Stdout) Exit() int {
    fmt.Println("stdout exit")
    return Ok
}
*/

var stdout = String{ len("stdout"), "stdout" }
var stdoutCommands = []Command{

    { stdout,
      STDOUT_CONFIG,
      stdoutBlock,
      0,
      0,
      nil },

    NilCommand,
}

func stdoutBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := STDOUT_CONFIG|CONFIG_VALUE
    Block(c, Modulers, STDOUT_MODULE, flag)

    return Ok
}

var stdoutModule = &Module{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    stdoutCommands,
    OUTPUT_MODULE,
}

func init() {
    Modulers = Load(Modulers, &Stdout{stdoutModule})
}