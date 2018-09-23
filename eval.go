// +gopherjs
// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package main

import (
	"github.com/funnelorg/funnel"
	"github.com/funnelorg/funnel/math"
	"github.com/gopherjs/gopherjs/js"
)

func Eval(code string) interface{} {
	return funnel.Eval(math.New(), "browser", code)
}

func main() {
	js.Global.Set("Funnel", map[string]interface{}{"Eval": Eval})
}
