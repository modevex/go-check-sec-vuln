package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

func main() {
	client := sentry.Init(sentry.ClientOptions{})
	fmt.Printf("%+v\n", client)
}
