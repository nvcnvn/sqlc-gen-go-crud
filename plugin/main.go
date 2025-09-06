package main

import (
	"github.com/sqlc-dev/plugin-sdk-go/codegen"

	golang "github.com/nvcnvn/sqlc-gen-go-curd/internal"
)

func main() {
	codegen.Run(golang.Generate)
}
