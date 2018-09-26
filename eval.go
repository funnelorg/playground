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
		done(format(funnel.Eval(url.Scope(math.Scope(builtin.Scope)), "browser", code)))
	}()
}

func format(v interface{}) interface{} {
	switch v := v.(type) {
	case stackable:
		return v.Stack()
	case error:
		return v.Error()
	case scopeWithKeys:
		result := map[interface{}]interface{}{}
		v.ForEachKeys(func(key interface{}) bool {
			result[key] = format(v.Get(key))
			return false
		})
		return result
	}
	return v
}

type stackable interface {
	Stack() string
}

type scopeWithKeys interface {
	Get(interface{}) interface{}
	ForEachKeys(fn func(interface{}) bool)
}

func main() {
	js.Global.Set("Funnel", map[string]interface{}{"Eval": Eval})
}
