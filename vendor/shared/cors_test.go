/*
The MIT License (MIT)

Copyright (c) 2016 Henri Koski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Source: https://github.com/heppu/simple-cors
*/

package shared

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCors(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "https://golang.org/", nil)
	req.Header.Set(origin, "localhost")
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wo := w.Header().Get(allowOrigin)
		ro := r.Header.Get(origin)
		if wo != ro {
			t.Fatalf("Wrong Allow-Origin '%s' expected '%s'", wo, ro)
		}

		wm := w.Header().Get(allowMethods)
		if wm != methods {
			t.Fatalf("Wrong Allow-Methods '%s' expected '%s'", wm, methods)
		}

		wh := w.Header().Get(allowHeaders)
		if wh != headers {
			t.Fatalf("Wrong Allow-Header '%s' expected '%s'", wh, headers)
		}
	})

	Handler(h).ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Fatal(http.StatusFound)
	}
}
