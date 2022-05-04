Pongo2gin
=========

Pongo2gin is an adapter that allows [pongo2 (v5)](https://github.com/flosch/pongo2) to be used with the
[Gin web framework](https://github.com/gin-gonic/gin).



Requirements
------------

Requires Gin 1.2 or higher and Pongo2.

Usage
-----

To use pongo2gin you need to set your router.HTMLRenderer to a new renderer
instance, this is done after creating the Gin router when the Gin application
starts up. You can use pongo2gin.Default() to create a new renderer with
default options, this assumes templates will be located in the "templates"
directory, or you can use pongo2.New() to specify a custom location.

To render templates from a route, call c.HTML just as you would with
regular Gin templates, the only difference is that you pass template
data as a pongo2.Context instead of gin.H type.

Basic Example
-------------

```go
package main

import (
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"magnax.ca/pongo2gin/v5"
)

func main() {
	router := gin.Default()

	// Use pongo2gin.Default() for default options or pongo2gin.New()
	// if you need to use custom RenderOptions.
	router.HTMLRender = pongo2gin.Default()

	router.GET("/", func(c *gin.Context) {
		// Use pongo2.Context instead of gin.H
		c.HTML(200, "hello.html", pongo2.Context{"name": "world"})
	})

	router.Run(":8080")
}
```

Template Caching
----------------

Templates will be cached if the current Gin Mode is set to anything but "debug" (`gin.DebugMode`),
this means the first time a template is used it will still load from disk, but
after that the cached template will be used from memory instead.

If the Gin Mode is set to "debug" then templates will be loaded from disk on
each request.

Caching is implemented by the Pongo2 library itself.

GoDoc
-----

https://pkg.go.dev/magnax.ca/pongo2gin/v5