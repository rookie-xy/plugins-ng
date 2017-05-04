/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "github.com/fsnotify/fsnotify"

    . "github.com/rookie-xy/main/types"

    "log"
    "fmt"
)

const (
    RESOURCE = "/data/service"
)

type FileConfigure struct {
    *Module
    *Configure

     watcher     *fsnotify.Watcher
     resource     string

     Notice       chan *Event
}

func (fc *FileConfigure) SetResource(resource string) int {
    if resource == "" {
        return Error
    }

    fc.resource = resource

    return Ok
}

func (fc *FileConfigure) GetResource() string {
    if fc.resource == "" {
        return ""
    }

    return fc.resource
}

func (fc *FileConfigure) SetConfigure() int {
    if fc.File == nil {
        fc.Error("file configure set error")
        return Error
    }

    flag := false
    if fc.Reader() == Error {
        fc.Error("configure read file error")
        flag = true
    }

    if fc.Closer() == Error {
        fc.Warn("file close error: %d\n", 10)
        return Error
    }

    if flag {
        return Error
    }

    return Ok
}

func (fc *FileConfigure) GetConfigure() int {
    if fc.File == nil {
        fc.File = NewFile(fc.Log)
    }

    resource := fc.GetResource()
    if resource == "" {
        return Error
    }

    if fc.Open(resource) == Error {
        fc.Error("configure open file error")
        return Error
    }

    return Ok
}

func (fc *FileConfigure) Clear() {
    return
}

func (fc *FileConfigure) Init(option *Option) int {
    if c := option.Configure; c != nil {
        fc.Configure = c
    } else {
        fc.Configure = NewConfigure(option.Log)
    }

    var (
        protocol,
        resource string
    )

    item := option.GetItem("configure")
    if item == nil {
        return Error
    }

    if protocol = item.(string); protocol != "file" {
        return Ignore
    }

    if fc.SetName(protocol) == Error {
        return Error
    }


    if item = option.GetItem("resource"); item != nil {
        resource = item.(string)

    } else {
        return Error
    }

    if fc.SetResource(resource) == Error {
        return Error
    }

    if watcher, error := fsnotify.NewWatcher(); error != nil {
        return Error
    } else {
        fc.watcher = watcher
    }

    if error := fc.watcher.Add(resource); error != nil {
        fmt.Println(resource, error)
        return Error
    }

    if fc.NewConfigurer(fc) == Error {
        return Error
    }

    return Ok
}

func (fc *FileConfigure) Main(configure *Configure) int {
    flag := Error

    if fc.GetConfigure() == Error {
        return flag
    }

    if fc.SetConfigure() == Error {
        return flag
    }

    notice := NewEvent()

    if flag == Error {
        notice.SetOpcode(LOAD)
        notice.SetName("load")
        configure.Event <- notice
    }

    defer fc.watcher.Close()

    quit := false

    for {
        select {

        case event := <-fc.watcher.Events:
            if event.Op & fsnotify.Write == fsnotify.Write {
                notice.SetOpcode(RELOAD)
                notice.SetName("reload")
                configure.Event <- notice
            }

        case err := <-fc.watcher.Errors:
            log.Println("error:", err)
        /*
        case e := <-fc.Notice.GetNotice():
            if op := e.GetOpcode(); op == SYSTEM_MODULE {
                quit = true
            }
            */
        }

        if quit {
            break
        }
    }

    fc.Clear()

    return Ok
}

func (fc *FileConfigure) Exit() int {
    //fileConfigure.Event <- 1
    fc.Clear()
    //fileConfigure.Quit()
    return Ok
}

func (fc *FileConfigure) Type() *Module {
    return fc.Self()
}

var FileConfigureModule = &FileConfigure{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        nil,
        SYSTEM_MODULE,
    },

    resource:  RESOURCE,
    Notice:    make(chan *Event),
}

func init() {
    Modulers = Load(Modulers, FileConfigureModule)
}