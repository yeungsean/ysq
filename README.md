# ysq

![Build Status](https://github.com/yeungsean/ysq/workflows/CI/badge.svg)
[![License](https://img.shields.io/github/license/yeungsean/ysq)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yeungsean/ysq)](https://goreportcard.com/report/github.com/yeungsean/ysq)
[![Coverage Status](https://img.shields.io/coveralls/github/yeungsean/ysq.svg)](https://coveralls.io/r/yeungsean/ysq)

## Go Version

go≥1.19


## Install

```shell
go get -v github.com/yeungsean/ysq
```


## Example Usage

```go
package main

import "fmt"

func main() {
}

func castInterface() {
    slice := []int64{1,2,3,4,5,6}
    interfaceSlice := FromSlice(slice).CastToInterface().ToSlice()
    printArgs := func(args []interface{}) {
        fmt.Printf("%#v\n", args)
    }
    printArgs(interfaceSlice)
}

func getTop3Element() {
    slice := []int64{1,2,3,4,5,6,7,8,9,10}
    res := FromSlice(slice).Take(3).ToSlice()
    fmt.Println(res) // [1,2,3]
}

func pager() {
    res := FromSequence(1, 20).Skip(10).Take(5).ToSlice(5)
    fmt.Println(res) // [11,12,13,14,15]
}

func sequence() {
    res1 := FromSequence(1, 10)
    fmt.Println(res1) // [1,2,3,4,5,6,7,8,9,10]

    res2 := FromSequence(1, 10, 2)
    fmt.Println(res2) // [1,3,5,7,9]
}

func filter() {
    // or Where
    res := FromSequence(1, 20).Filter(func(i int) bool {
		return i < 10
	}).ToSlice(10)
    fmt.Println(res) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
}

func contains() {
    // or In
    res := FromSequence(1, 100).Contains(func(i int) bool {
        return i%2 == 0
    })
    fmt.Println(res) // true

    res = FromSequence(1, 100).In(func(i int) bool {
        return i == 10
    })
    fmt.Println(res) // true

    res = FromSequence(1, 100).Contains(func(i int) bool {
        return i > 1000
    })
    fmt.Println(res) // false
}

func all() {
    res := FromSequence(1, 100).All(func(i int) bool {
        return i < 1000
    })
    fmt.Println(res) // true

    res = FromSlice([]int{1,3,5,7,9}).All(func(i int) bool {
        return i%2 == 0
    })
    fmt.Println(res) // false
}
```
