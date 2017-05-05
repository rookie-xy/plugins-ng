/*
 * Copyright (C) 2017 Meng Shi
 */

package stdout

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

type OutputStdout struct {
    *Module

     status   bool
     channel  string
}

func NewOutputStdout() *OutputStdout {
    return &OutputStdout{}
}

type OutputStdoutContext struct {
    Name   String
    Data   [32]*unsafe.Pointer
}

var stdoutOutput = String{ len("stdout_output"), "stdout_output" }
var outputStdoutContext = &OutputStdoutContext{
    Name: stdoutOutput,
}

func (r *OutputStdoutContext) Create() unsafe.Pointer {
    stdout := NewOutputStdout()
    if stdout == nil {
        return nil
    }

    stdout.channel = "zhang yue"
    stdout.status = false

    return unsafe.Pointer(stdout)
}

func (r *OutputStdoutContext) GetDatas() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    stdoutStatus = String{ len("status"), "status" }
    stdoutChannel = String{ len("channel"), "channel" }
    outputStdout OutputStdout
)

var outputStdoutCommands = []Command{

    { stdoutStatus,
      STDOUT_CONFIG|CONFIG_VALUE,
      SetFlag,
      0,
      unsafe.Offsetof(outputStdout.status),
      nil },

    { stdoutChannel,
      STDOUT_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(outputStdout.channel),
      nil },

    NilCommand,
}

func (r *OutputStdout) Init(o *Option) int {
    context := r.Context.GetDatas()

    for _, v := range context {
        if v != nil {
            this := (*OutputStdout)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.channel, this.status)
        } else {
            break
        }
    }

    return Ok
}

func (r *OutputStdout) Main(cfg *Configure) int {
    fmt.Println("output main")
    return Ok
}

func (r *OutputStdout) Exit() int {
    fmt.Println("output exit")
    return Ok
}

var outputStdoutModule = &Module{
    MODULE_V1,
    CONTEXT_V1,
    outputStdoutContext,
    outputStdoutCommands,
    STDOUT_MODULE,
}

func init() {
    Modulers = Load(Modulers, &OutputStdout{Module:outputStdoutModule})
}
