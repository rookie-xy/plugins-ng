/*
 * Copyright (C) 2017 Meng Shi
 */

package simple

import (
    . "github.com/rookie-xy/main/types"
)

type SimpleOption struct {
    *Module
}

func (r *SimpleOption) Init(o *Option) int {
    argv := o.GetArgv()

    for i := 1; i < o.GetArgc(); i++ {

        if argv[i][0] != '-' {
            return Error
        }

        switch argv[i][1] {

        case 'c':
	           if argv[i + 1] == "" {
                return Error
            }

            /* file://resource= */
            o.SetItem("configure", argv[i + 1])
            i++
            break

        case 'r':
	           if argv[i + 1] == "" {
	               return Error
	           }

            /* zookeeper://resource= */
            o.SetItem("resource", argv[i + 1])
            i++
            break

        case 'f':
	           if argv[i + 1] == "" {
	               return Error
	           }

            /* yaml, json, xml ... */
            o.SetItem("format", argv[i + 1])
            i++
            break

        case 't':
            o.SetItem("test", true)
	           break

        default:
            o.SetItem("invaild", "")
            o.Info("not found any option")
            break
        }
    }

    configure := NewConfigure(o.Log)

    if item := o.GetItem("format"); item != nil {
        name := item.(string)

        for _, codec := range Codecs {
            if codec.Type(name) == Ignore {
                continue
            }

            codec.Init(nil)

            code := NewCode(codec)
            //code.SetName(item.(string))
            configure.Code = &code
        }
    } else {
        return Error
    }

    o.Configure = configure

    return Ok
}

func (r *SimpleOption) Main(c *Configure) int {
    return Ignore
}

func (r *SimpleOption) Exit() int {
    return Ignore
}

func (r *SimpleOption) Type() *Module {
    return r.Self()
}

var SimpleOptionModule = &SimpleOption{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        nil,
        SYSTEM_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, SimpleOptionModule)
}
