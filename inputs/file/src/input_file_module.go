/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

type InputFile struct {
    *Module

     host     string
     group    string
     types    string
     dir      string
     pattern  string
}

func NewInputFile() *InputFile {
    return &InputFile{}
}

type InputFileContext struct {
    *Context
}

var fileInput = String{ len("file_input"), "file_input" }
var inputFileContext = &InputFileContext{
    Context: &Context{
        Name: fileInput,
    },
}

func (r *InputFileContext) Create() unsafe.Pointer {
    file := NewInputFile()
    if file == nil {
        return nil
    }

    file.host = "127.0.0.1"
    file.group = "example"
    file.types = "de"
    file.dir = "/data/logs/"
    file.pattern = "file.log"

    return unsafe.Pointer(file)
}

func (r *InputFileContext) Contexts() *Context {
    return r.Get()
}

var (
    host    = String{ len("host"), "host" }
    group   = String{ len("group"), "group" }
    types   = String{ len("type"), "type" }
    dir     = String{ len("dir"), "dir" }
    pattern = String{ len("pattern"), "pattern" }

    inputFile InputFile
)

var inputFileCommands = []Command{

    { host,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.host),
      nil },

    { group,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.group),
      nil },

    { types,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.types),
      nil },

    { dir,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.dir),
      nil },

    { pattern,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.pattern),
      nil },

    NilCommand,
}

func (r *InputFile) Init(o *Option) int {
    context := r.Context.Contexts()

    for _, v := range context.Data {
        if v != nil {
            this := (*InputFile)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.types, this.pattern, this.group, this.dir, this.host)
        } else {
            break
        }
    }

    return Ok
}

func (r *InputFile) Main(cfg *Configure) int {
    fmt.Println("File main")
    return Ok
}

func (r *InputFile) Exit() int {
    fmt.Println("File exit")
    return Ok
}

func (r *InputFile) Type() *Module {
    return r.Self()
}

var inputFileModule = &InputFile{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        inputFileContext,
        inputFileCommands,
        FILE_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, inputFileModule)
}
