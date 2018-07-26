==========================
`DOM Manipulation`_ in Go_
==========================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/godom?status.png
   :target: https://godoc.org/github.com/siongui/godom

.. image:: https://api.travis-ci.org/siongui/godom.png?branch=master
   :target: https://travis-ci.org/siongui/godom

.. image:: https://goreportcard.com/badge/github.com/siongui/godom
   :target: https://goreportcard.com/report/github.com/siongui/godom

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/godom/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/godom.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D

Make `DOM Manipulation`_ in Go_ as similar to JavaScript_ as possible via
GopherJS_. For DOM Manipulation via WebAssembly_, visit wasm_ directory.

  - `Ubuntu 18.04`_
  - Go_

.. contents:: **Table of Content**


Why?
++++


Why not use GopherJS directly?
##############################

Because the code written directly by GopherJS without any wrapper is really
ugly. For example, if you want to *getElementById*, you need to write:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  foo := js.Global.Get("document").Call("getElementById", "foo")

With *godom*, you write:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  foo := Document.GetElementById("foo")

which looks like *JavaScript* and more readable.

Why not use existing `go-js-dom`_?
##################################

Because it's too restricted, and sometimes need to do a lot of type casting.
For example, if you have an *audio* element with id *foo* and you want to call
the *Play()* method, you need to write the following code:

.. code-block:: go

  import (
  	"honnef.co/go/js/dom"
  )

  a := dom.GetWindow().Document().GetElementByID("foo").(*dom.HTMLAudioElement)
  a.Play()

If you use *querySelectorAll* to select a lot of such elements, you need to do a
lot of type casting, which is really disturbing.

With *godom*, you can write like this:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  a := Document.GetElementById("foo")
  a.Play()


What if the method/property is not implemented in *godom*?
##########################################################

*godom* is only a wrapper for GopherJS. If something is not implemented, you can
still use the GopherJS methods to call or get the method/property you need.
For example, if the *Play()* method of the audio element is not implemented, you
can use GopherJS *Call* method to call *play* method directly:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  a := Document.GetElementById("foo")
  a.Call("play")


Code Example
++++++++++++

- `Frontend Programming in Go`_: If you have no experience of GopherJS before,
  read this.
- `Synonyms - Go and JavaScript`_: If you have some experience about GopherJS,
  this serves as references for quick start.


Issues
++++++

null test
#########

Test if event.state is null in ``popstate`` event listener:

.. code-block:: go

  	ih := Document.QuerySelector("#infoHistory")

  	Window.AddEventListener("popstate", func(e Event) {
  		if e.Get("state") == nil {
  			ih.SetInnerHTML("Entry Page")
  		} else {
  			ih.SetInnerHTML(e.Get("state").String())
  		}
  	})


HTML dataset (data-* attribute)
###############################

Assume we have the following element:

.. code-block:: html

  <p id="foo" data-content="content of person 1"></p>

You can access the ``data-content`` as follows:

.. code-block:: go

  p := Document.QuerySelector("#foo")
  content := p.Dataset().Get("content").String()


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [3] | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Google search <https://www.google.com/search?q=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `add a method to an external package - Google search <https://www.google.com/search?q=add+a+method+to+an+external+package>`_

.. [4] `[Golang] Add Method to Existing Type in External Package <https://siongui.github.io/2017/02/11/go-add-method-function-to-type-in-external-package/>`_

.. [5] `JavaScript Remove All Children of a DOM Element <https://siongui.github.io/2012/09/26/javascript-remove-all-children-of-dom-element/>`_

.. [6] `How to do insert After() in JavaScript without using a library? - Stack Overflow <http://stackoverflow.com/a/32135318>`_

.. [7] `javascript element position <https://www.google.com/search?q=javascript+element+position>`_

       `javascript - Retrieve the position (X,Y) of an HTML element - Stack Overflow <http://stackoverflow.com/questions/442404/retrieve-the-position-x-y-of-an-html-element>`_

.. [8] `javascript check class exists - Google search <https://www.google.com/search?q=javascript+check+class+exists>`_

       `javascript - Test if an element contains a class? - Stack Overflow <http://stackoverflow.com/a/5898748>`_

.. [9] | `Who is using GopherJS? : golang <https://www.reddit.com/r/golang/comments/5urqny/who_is_using_gopherjs/>`_
       | `GopherJS 1.8-1 is released : golang <https://www.reddit.com/r/golang/comments/5upkkc/gopherjs_181_is_released/>`_

.. [10] `Go Report Card | Go project code quality report cards <https://goreportcard.com/>`_
.. [11] `Shields.io: Quality metadata badges for open source projects  <https://shields.io/>`_

.. [12] `HTML DOM Style object <https://www.w3schools.com/jsref/dom_obj_style.asp>`_

.. [13] | `javascript is focused - Google search <https://www.google.com/search?q=javascript+is+focused>`_
        | `javascript is focused - DuckDuckGo search <https://duckduckgo.com/?q=javascript+is+focused>`_
        | `javascript is focused - Ecosia search <https://www.ecosia.org/search?q=javascript+is+focused>`_
        | `javascript is focused - Qwant search <https://www.qwant.com/?q=javascript+is+focused>`_
        | `javascript is focused - Bing search <https://www.bing.com/search?q=javascript+is+focused>`_
        | `javascript is focused - Yahoo search <https://search.yahoo.com/search?p=javascript+is+focused>`_
        | `javascript is focused - Baidu search <https://www.baidu.com/s?wd=javascript+is+focused>`_
        | `javascript is focused - Yandex search <https://www.yandex.com/search/?text=javascript+is+focused>`_

.. _DOM Manipulation: https://www.google.com/search?q=DOM+Manipulation
.. _Go: https://golang.org/
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _GopherJS: http://www.gopherjs.org/
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _wasm: wasm
.. _Ubuntu 18.04: http://releases.ubuntu.com/18.04/
.. _Go 1.8: https://golang.org/dl/
.. _go-js-dom: https://github.com/dominikh/go-js-dom
.. _UNLICENSE: http://unlicense.org/
.. _Frontend Programming in Go: https://siongui.github.io/2017/12/04/frontend-programming-in-go/
.. _Synonyms - Go and JavaScript: https://siongui.github.io/2017/12/07/synonyms-go-and-javascript/

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
