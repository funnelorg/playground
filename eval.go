// +js
// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package main

import (
	"github.com/funnelorg/funnel"
	"github.com/funnelorg/funnel/builtin"
	"github.com/funnelorg/funnel/math"
	"github.com/funnelorg/funnel/url"
	"github.com/gopherjs/gopherjs/js"
)

func Eval(code string, done func(interface{})) {
	go func() {
		result := funnel.Eval(url.Scope(math.Scope(builtin.Scope)), "browser", code)
		if s, ok := result.(stackable); ok {
			result = s.Stack()
		} else if err, ok := result.(error); ok {
			result = map[string]interface{}{"Error": err.Error()}
		}
		done(result)
	}()
}

type stackable interface {
	Stack() string
}

func main() {
	js.Global.Set("Funnel", map[string]interface{}{"Eval": Eval})
}
