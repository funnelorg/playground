// +js
// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package main

import (
	"github.com/funnelorg/funnel"
	"github.com/funnelorg/funnel/math"
	"github.com/funnelorg/funnel/url"	
	"github.com/gopherjs/gopherjs/js"
)

func Eval(code string, done func(interface{})) {
	go func() {
		result := funnel.Eval(url.New(math.New()), "browser", code)
		if err, ok := result.(error); ok {
			result = map[string]interface{}{"Error": err.Error()}
		}
		done(result)
	}()
}

func main() {
	js.Global.Set("Funnel", map[string]interface{}{"Eval": Eval})
}
