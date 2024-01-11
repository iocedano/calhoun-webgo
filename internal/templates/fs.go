package template

import "embed"

// FS holds our static web server content.
//
//go:embed pages/*.gohtml
//go:embed widgets/*.gohtml
//go:embed layout.gohtml
var FS embed.FS
