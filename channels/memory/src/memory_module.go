/*
 * Copyright (C) 2017 Meng Shi
 */

package memory

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
    "fmt"
)

const (
    MEMORY_MODULE = CHANNEL_MODULE|0x01000000
    MEMORY_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Memory struct {
    *Module
}

func (r *Memory) Init(o *Option) int {
    fmt.Println("ChannelMemory init")
    return Ok
}

func (r *Memory) Main(cfg *Configure) int {
    fmt.Println("ChannelMemory main")
    return Ok
}

func (r *Memory) Exit() int {
    fmt.Println("ChannelMemory exit")
    return Ok
}

func (r *Memory) Type() *Module {
    return r.Self()
}

var	memory = String{ len("memory"), "memory" }
var memoryCommands = []Command{

    { memory,
      MEMORY_CONFIG,
      memoryBlock,
      0,
      0,
      nil },

    NilCommand,
}

func memoryBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := MEMORY_CONFIG|CONFIG_VALUE
    Block(c, Modulers, MEMORY_MODULE, flag)

    return Ok
}

var memoryModule = &Memory{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        memoryCommands,
        CHANNEL_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, memoryModule)
}
