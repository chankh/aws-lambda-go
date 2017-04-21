package main

import (
	"fmt"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(req Request, ctx *runtime.Context) (string, error) {
	name := req.Name
	return fmt.Sprintf("Hello %s!", name), nil
}

type Request struct {
	Name string `json:"name"`
}
