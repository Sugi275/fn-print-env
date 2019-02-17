package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(printEnv))
}

func printEnv(ctx context.Context, in io.Reader, out io.Writer) {

	msg := struct {
		LogLevel string `json:"loglevel"`
	}{
		fmt.Sprintf(os.Getenv("LOG_LEVEL")),
	}

	json.NewEncoder(out).Encode(&msg)
}
