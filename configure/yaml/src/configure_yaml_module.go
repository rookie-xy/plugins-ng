/*
 * Copyright (C) 2017 Meng Shi
 */

package yaml

import (
      "gopkg.in/yaml.v2"
    . "github.com/rookie-xy/main/types"
)

type YamlConfigure struct {
    *Module
    *Configure
}

func (yc *YamlConfigure) Marshal(in interface{}) ([]byte, error) {
    return nil, nil
}

func (yc *YamlConfigure) Unmarshal(in []byte, out interface{}) int {
    if err := yaml.Unmarshal(in, out); err != nil {
        return Error
    }

    return Ok
}

func (yc *YamlConfigure) Init(option *Option) int {
    item := option.GetItem("format")
    if item == nil {
        return Error
    }

    if format := item.(string);
       format == "yaml" || format == "yml" {
        /* Ok */
    } else {
        return Ignore
    }

    if yc.Configure == nil {
        if c := option.Configure; c != nil {
            yc.Configure = c

        } else {
            yc.Configure = NewConfigure(option.Log)
        }
    }

    yc.NewParser(yc)

    return Ok
}

func (yc *YamlConfigure) Main(_ *Configure) int {
    return Ok
}

func (yc *YamlConfigure) Exit() int {
    return Ok
}

func (yc *YamlConfigure) Type() *Module {
    return yc.Self()
}

var YamlConfigureModule = &YamlConfigure{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        nil,
        SYSTEM_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, YamlConfigureModule)
}
