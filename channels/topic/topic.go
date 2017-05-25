package topic

import (
    . "github.com/rookie-xy/worker/types"
)

type Topic struct {
    name string
}

func NewTopic(name string) *Topic {
    return &Topic{
        name: name,
    }
}

func (r *Topic) New() Channel {
    return nil
}

func (r *Topic) Init(configure interface{}) int {
    return Ok
}

func (r *Topic) Push() int {
    return Ok
}

func (r *Topic) Pull() int {
    return Ok
}
