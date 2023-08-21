package service

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strings"
)

func readInput(reader *bufio.Reader) (FetchArgs, error) {
	args, err := reader.ReadString('\n')
	if err != nil {
		return FetchArgs{}, err
	}
	args = strings.TrimSpace(args)

	if len(args) == 0 {
		return FetchArgs{}, EmptyInputError
	}

	if args == "exit" {
		exitChan <- struct{}{}
		return FetchArgs{}, ExitError
	}

	arg, err := parseArgs(args)
	if err != nil {
		return FetchArgs{}, err
	}
	return arg, nil
}

func parseArgs(arg string) (FetchArgs, error) {
	args := strings.Split(arg, " ")
	if len(args) < 2 {
		fmt.Println("Error: invalid input")
		return FetchArgs{}, fmt.Errorf("invalid input")
	}
	if filepath.Base(args[0]) != "fetch" {
		fmt.Println("Error: invalid input")
		return FetchArgs{}, fmt.Errorf("invalid input")
	}

	pathToSave := filepath.Dir(args[0])
	strategy := fetchStrategySaveContent
	args = args[1:]
	if args[0] == "--metadata" {
		strategy = fetchStrategyGetMetadata
		args = args[1:]
	}

	for idx := range args {
		if !strings.HasPrefix(args[idx], "http://") && !strings.HasPrefix(args[idx], "https://") {
			args[idx] = "https://" + args[idx]
		}
	}

	return FetchArgs{
		downloadPath: pathToSave,
		strategy:     strategy,
		urls:         args,
	}, nil
}
