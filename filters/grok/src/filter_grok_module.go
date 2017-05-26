/*
 * Copyright (C) 2017 Meng Shi
 */

package grok

import (
      "unsafe"

    . "github.com/rookie-xy/worker/types"
    . "github.com/rookie-xy/worker/modules"
)

const (
    GROK_MODULE = CHANNEL_MODULE|MAIN_MODULE
    GROK_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

var grokModule = String{ len("grok_module"), "grok_module" }
var filterGrokContext = &Context{
    grokModule,
    nil,
    nil,
}

var	grok = String{ len("grok"), "grok" }
var filterGrokCommands = []Command_t{

    { grok,
      GROK_CONFIG,
      grokBlock,
      0,
      0,
      nil },

    NilCommand,
}

func grokBlock(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := GROK_CONFIG|CONFIG_VALUE
    Block(c, Modules, GROK_MODULE, flag)

    return Ok
}

var filterGrokModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    filterGrokCommands,
    CHANNEL_MODULE,
}

func init() {
    Modules = Load(Modules, &filterGrokModule)
}