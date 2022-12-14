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

import (
    "fmt"

    "github.com/yeungsean/ysq"
    "github.com/yeungsean/ysq/pkg/delegate"
)

func main() {
}

func castInterface() {
    slice := []int64{1,2,3,4,5,6}
    interfaceSlice := ysq.FromSlice(slice).CastToInterface().ToSlice()
    printArgs := func(args []interface{}) {
        fmt.Printf("%#v\n", args)
    }
    printArgs(interfaceSlice)
}

func getTop3Element() {
    slice := []int64{1,2,3,4,5,6,7,8,9,10}
    res := ysq.FromSlice(slice).Take(3).ToSlice()
    fmt.Println(res) // [1,2,3]
}

func pager() {
    res := ysq.FromSequence(1, 20).Skip(10).Take(5).ToSlice(5)
    fmt.Println(res) // [11,12,13,14,15]
}

func sequence() {
    res1 := ysq.FromSequence(1, 10)
    fmt.Println(res1) // [1,2,3,4,5,6,7,8,9]

    res2 := ysq.FromSequence(1, 10, 2)
    fmt.Println(res2) // [1,3,5,7,9]
}

func filter() {
    // or Where
    res := ysq.FromSequence(1, 20).Filter(func(i int) bool {
        return i < 10
    }).ToSlice(10)
    fmt.Println(res) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
}

func contains() {
    // or In
    res := ysq.FromSequence(1, 100).Contains(func(i int) bool {
        return i%2 == 0
    })
    fmt.Println(res) // true

    res = ysq.FromSequence(1, 100).In(func(i int) bool {
        return i == 10
    })
    fmt.Println(res) // true

    res = ysq.FromSequence(1, 100).Contains(func(i int) bool {
        return i > 1000
    })
    fmt.Println(res) // false
}

func all() {
    res := ysq.FromSequence(1, 100).All(func(i int) bool {
        return i < 1000
    })
    fmt.Println(res) // true

    res = ysq.FromSlice([]int{1,3,5,7,9}).All(func(i int) bool {
        return i%2 == 0
    })
    fmt.Println(res) // false
}

func mapReduce() {
    func() {
        res := ysq.FromSequence(1, 11).Select(func(v int) int {
            return v + 1
        }).Reduce(0, func(total, current int) int {
            return total + current
        })
        fmt.Println(res) // 65
    }()

    func() {
        res := ysq.FromSequence(1, 11).Select(func(v int) int {
            return v + 1
        }).SumToInt(func(current int) int {
            return current
        })
        fmt.Println(res) // 65
    }()
}

// like python
func partial() {
    tmpAction2 := func(arg1, arg2 int) {
        fmt.Println(arg1, arg2)
    }
    func() {
        var fa delegate.Action2[int, int] = tmpAction2
        delayCall := fa.Partial(5)
        delayCall(10) // print 5, 10
        delayCall(100) // print 5, 100
    }()

    tmpSumFunc2 := func(arg1, arg2 int) int {
        return arg1 + arg2
    }
    func() {
        var ff2 delegate.Func2[int, int, int] = tmpSumFunc2
        delayCall := ff2.Partial(5)
        res := delayCall(10)
        fmt.Println(res) // 15

        res = delayCall(-10)
        fmt.Println(res) // -5
    }()
}
```
