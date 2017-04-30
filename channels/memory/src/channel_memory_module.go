/*
 * Copyright (C) 2017 Meng Shi
 */

package memory

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

type ChannelMemory struct {
    *Module

     name     string
     size     int
}

func NewChannelMemory() *ChannelMemory {
    return &ChannelMemory{}
}

type ChannelMemoryContext struct {
    *Context
}

var memoryChannel = String{ len("memory_channel"), "memory_channel" }
var channelMemoryContext = &ChannelMemoryContext{
    Context: &Context{
        Name: memoryChannel,
    },
}

func (r *ChannelMemoryContext) Create() unsafe.Pointer {
    memory := NewChannelMemory()
    if memory == nil {
        return nil
    }

    memory.name = "channel name"
    memory.size = 1024

    return unsafe.Pointer(memory)
}

func (r *ChannelMemoryContext) Contexts() *Context {
    return r.Get()
}

var (
    name = String{ len("name"), "name" }
    size = String{ len("size"), "size" }
    channelMemory ChannelMemory
)

var channelMemoryCommands = []Command{

    { name,
      MEMORY_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(channelMemory.name),
      nil },

    { size,
      MEMORY_CONFIG|CONFIG_VALUE,
      SetNumber,
      0,
      unsafe.Offsetof(channelMemory.size),
      nil },

    NilCommand,
}

var channelMemoryModule = &ChannelMemory{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        channelMemoryContext,
        channelMemoryCommands,
        MEMORY_MODULE,
    },
}

func (r *ChannelMemory) Init(o *Option) int {
    context := r.Context.Contexts()

    for _, v := range context.Data {
        if v != nil {
            this := (*ChannelMemory)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.name, this.size)
        } else {
            break
        }
    }

    return Ok
}

func (r *ChannelMemory) Main(cfg *Configure) int {
    fmt.Println("Memory main")
    return Ok
}

func (r *ChannelMemory) Exit() int {
    fmt.Println("Memory exit")
    return Ok
}

func (r *ChannelMemory) Type() *Module {
    return r.Self()
}

func init() {
    Modulers = Load(Modulers, channelMemoryModule)
}