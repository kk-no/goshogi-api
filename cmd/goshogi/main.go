package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/kk-no/yaneurago/usi"
)

func main() {
	ctx := context.Background()

	engine := usi.NewEngine()
	if err := engine.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer engine.Disconnect(ctx)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			line := scanner.Text()
			engine.SendCommand(ctx, line)
			if line == "quit" {
				break
			}
		}
	}
}
