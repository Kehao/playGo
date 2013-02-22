package main

import (
  "fmt"
)
type I interface {
  Get() int
  Put(int)
}

func g(any interface{}) int {
  //any.(I).Put(1)
  return any.(I).Get()
}

type S struct { i int }
func (p *S) Get() int { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct { i int }
func (p *R) Get() int { return p.i }
func (p *R) Put(v int) { p.i = v }

type Test struct { i int }

func f(p I) {
  p.Put(1)
}

func main() {
  var s S
  //t := new(Test)
  f(&s)
  get := g(&s)
  fmt.Println(get)
  //var s1 S
  //get1 := g(s1)
  //fmt.Println(get1)
}
