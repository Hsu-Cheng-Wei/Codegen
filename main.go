package main

import (
	"codegen/regex"
	"flag"
	"fmt"
)

func main() {
	var args regex.TemplateArgs
	flag.BoolVar(&args.Configure, "config", false, "[Option] generator template init config")
	flag.BoolVar(&args.UseQuery, "q", false, "query template")
	flag.BoolVar(&args.UseCommand, "c", false, "command template")
	flag.StringVar(&args.Name, "n", "", "template name")
	flag.StringVar(&args.Topic, "t", "", "template topic")
	flag.BoolVar(&args.HasValidation, "v", false, "[Option] generator validate template")
	flag.Parse()

	cmdArgs := flag.Args()

	for s := range cmdArgs {
		fmt.Println("Command Args: " + flag.Arg(s))
	}

	if args.UseQuery && args.UseCommand {
		fmt.Println("Query and Command only select one...")
		return
	}

	if args.UseQuery {
		args.Type = "query"
	}

	if args.UseCommand {
		args.Type = "command"
	}

	regex.Regex(&args)

	fmt.Println("Code gen complete...")
}
