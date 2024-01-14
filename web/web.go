package web

import (
	"embed"
	"io/fs"
)

//go:generate bun i
//go:generate bun run build

//go:embed dist/*
var dist embed.FS

var static fs.FS

func init() {
	var err error
	static, err = fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
}

func StaticFs() fs.FS { return static }
