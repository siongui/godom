==========================
`DOM Manipulation`_ in Go_
==========================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/godom/wasm?status.png
   :target: https://godoc.org/github.com/siongui/godom/wasm

.. image:: https://api.travis-ci.org/siongui/godom.png?branch=master
   :target: https://travis-ci.org/siongui/godom

.. image:: https://goreportcard.com/badge/github.com/siongui/godom/wasm
   :target: https://goreportcard.com/report/github.com/siongui/godom/wasm

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/godom/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/godom.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D

Make `DOM Manipulation`_ in Go_ as similar to JavaScript_ as possible via
WebAssembly_. For DOM Manipulation via GopherJS_, visit root_ directory.

  - `Ubuntu 18.04`_
  - Go_ (1.11 or later)

.. contents:: **Table of Content**

Install
+++++++

.. code-block:: bash

  $ GOARCH=wasm GOOS=js go get -u github.com/siongui/godom/wasm


Why?
++++


Why not use `syscall/js`_ directly?
###################################

Because the code written directly using `syscall/js`_ without any wrapper is
really ugly. For example, if you want to *getElementById*, you need to write:

.. code-block:: go

  import (
  	"syscall/js"
  )

  foo := js.Global().Get("document").Call("getElementById", "foo")

With *godom*, you write:

.. code-block:: go

  import (
  	. "github.com/siongui/godom/wasm"
  )

  foo := Document.GetElementById("foo")

which looks like *JavaScript* and more readable.


What if the method/property is not implemented in *godom*?
##########################################################

*godom* is only a wrapper for `syscall/js`_ package. If something is not
implemented, you can still use methods in `syscall/js`_ to call or get the
method/property you need. For example, if the *Play()* method of the audio
element is not implemented, you can use `syscall/js`_ *Call* method to call
*play* method directly:

.. code-block:: go

  import (
  	. "github.com/siongui/godom/wasm"
  )

  a := Document.GetElementById("foo")
  a.Call("play")


Code Example
++++++++++++

- `Frontend Programming in Go`_: If you have no experience of Go WebAssembly
  before, read this.


Reference
+++++++++

.. [1] `WebAssembly · golang/go Wiki · GitHub <https://github.com/golang/go/wiki/WebAssembly>`_


.. %s/o \*Object/v Value/gc

.. _DOM Manipulation: https://www.google.com/search?q=DOM+Manipulation
.. _Go: https://golang.org/
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _syscall/js: https://godoc.org/syscall/js
.. _GopherJS: http://www.gopherjs.org/
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _root: https://github.com/siongui/godom
.. _Ubuntu 18.04: http://releases.ubuntu.com/18.04/
.. _UNLICENSE: http://unlicense.org/
.. _Frontend Programming in Go: https://siongui.github.io/2017/12/04/frontend-programming-in-go/
