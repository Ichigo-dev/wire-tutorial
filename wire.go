// +build wireinject

package main

import "github.com/google/wire"

func NewApp() Event {
  wire.Build(NewEvent, NewGreeter, NewMessage)
  return Event{}
}
