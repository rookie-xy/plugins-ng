/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
    //. "github.com/rookie-xy/main/modules"
"fmt"
)

type InputFile struct {
    *Module

     host   string
     group  string
     types  string
     path   string
     name   Array

     codec  Code
}

func NewInputFile() *InputFile {
    return &InputFile{}
}

type InputFileContext struct {
    Name   String
    Data   [32]*unsafe.Pointer
}

var fileInput = String{ len("file_input"), "file_input" }
var inputFileContext = &InputFileContext{
    Name: fileInput,
}

func (r *InputFileContext) Create() unsafe.Pointer {
    file := NewInputFile()
    if file == nil {
        return nil
    }

    file.host = "127.0.0.1"
    file.group = "example"
    file.types = "de"
    file.path = "/data/logs/"
    //file.name = Array{}

    return unsafe.Pointer(file)
}

func (r *InputFileContext) GetDatas() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    host  = String{ len("host"), "host" }
    group = String{ len("group"), "group" }
    types = String{ len("type"), "type" }
    path  = String{ len("path"), "path" }
    name  = String{ len("name"), "name" }
    codec = String{ len("codec"), "codec" }

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

    { path,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.path),
      nil },

    { name,
      FILE_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(inputFile.name),
      nil },

    { codec,
      FILE_CONFIG|CONFIG_BLOCK,
      SetCodec,
      0,
      unsafe.Offsetof(inputFile.codec),
      nil },

    NilCommand,
}

func (r *InputFile) Init(o *Option) int {
    context := r.Context.GetDatas()

    for _, v := range context {
        if v != nil {
            this := (*InputFile)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.types, this.group, this.path, this.host)

            for i := 0; i < this.name.GetLength(); i++ {
                fmt.Println(this.name.GetData(i))
            }

            codec := this.codec.New()
            codec.Encode(nil)
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

var inputFileModule = &Module{
    MODULE_V1,
    CONTEXT_V1,
    inputFileContext,
    inputFileCommands,
    FILE_MODULE,
}

func init() {
    Modulers = Load(Modulers, &InputFile{Module:inputFileModule})
}
