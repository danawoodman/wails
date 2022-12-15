package options

import (
	"io/fs"
	"net/http"
)

type WindowState int

const (
	WindowStateNormal WindowState = iota
	WindowStateMinimised
	WindowStateMaximised
	WindowStateFullscreen
)

type Window struct {
	// Alias is a human-readable name for the window. This can be used to reference the window in the frontend.
	Alias            string
	Title            string
	Width, Height    int
	AlwaysOnTop      bool
	URL              string
	DisableResize    bool
	Frameless        bool
	MinWidth         int
	MinHeight        int
	MaxWidth         int
	MaxHeight        int
	EnableDevTools   bool
	StartState       WindowState
	Mac              *MacWindow
	BackgroundColour *RGBA
	Assets           Assets
}

var WindowDefaults = &Window{
	Title:          "",
	Width:          1024,
	Height:         768,
	URL:            "https://wails.io",
	EnableDevTools: true,
}

type Assets struct {
	// URL to load the `index.html` file from. If this is a relative path, it will be resolved relative to the `FS` filesystem
	URL string
	// FS to use for loading assets from
	FS fs.FS
	// Handler is a custom handler to use for serving assets. If this is set, the `URL` and `FS` fields are ignored.
	Handler http.Handler
	// Middleware is a custom middleware to use for serving assets. If this is set, the `URL` and `FS` fields are ignored.
	Middleware func(http.Handler) http.Handler
}

type RGBA struct {
	Red, Green, Blue, Alpha uint8
}
