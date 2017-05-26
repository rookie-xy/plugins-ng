/*
 * Copyright (C) 2017 Meng Shi
 */

package yaml

import (
    yml "gopkg.in/yaml.v2"
    .   "github.com/rookie-xy/worker/types"
)

type Yaml struct {
    name string
}

func NewYaml() *Yaml {
    return &Yaml{}
}

var yaml = &Yaml{
    name: "yaml",
}

func (r *Yaml) New() Codec {
    yaml := NewYaml()

    yaml.name = "yaml"

    return yaml
}

func (r *Yaml) Init(configure interface{}) int {
    return Ignore
}

func (r *Yaml) Encode(in interface{}) (interface{}, error) {
    out, error := yml.Marshal(in);
    if error != nil {
        return nil, error
    }

    return out, nil
}

func (r *Yaml) Decode(in []byte) (interface{}, error) {
    var out interface{}

    if e := yml.Unmarshal(in, &out); e != nil {
        return nil, e
    }

    return out, nil
}

func (r *Yaml) Type(name string) int {
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = append(Codecs, yaml)
}
