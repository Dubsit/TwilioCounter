package main

import (
    "github.com/sfreiberg/gotwilio"
    "github.com/hoisie/web"
    "fmt"
    "strconv"
    "strings"
)

var ryan int
var john int

func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
}

func hello(ctx *web.Context, val string) string {
    for k, v := range ctx.Params {
        if(k=="sms" && strings.Contains(v, "Ryan")) {
            ryan = ryan + 1
        }
        if(k=="sms" && strings.Contains(v, "John")){
            john = john + 1
        }
    }
        fmt.Println()
        return "at " + strconv.Itoa(ryan) + " for Ryan and " + strconv.Itoa(john) + " for John"
}
