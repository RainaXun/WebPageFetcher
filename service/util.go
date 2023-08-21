package service

import (
	"errors"
	"fmt"
	"sync"
)

// Constants
const (
	fetchStrategySaveContent = "save"
	fetchStrategyGetMetadata = "metadata"
)

// Mutex
var printMutex sync.Mutex

// Errors
var (
	EmptyInputError = errors.New("empty input")
	ExitError       = errors.New("exit")
)

// FetchArgs is a struct that holds the arguments passed in from the command line
type FetchArgs struct {
	downloadPath string
	strategy     string
	urls         []string
}

// Metadata is a struct that holds the metadata of a webpage
type Metadata struct {
	Site      string
	NumLinks  int
	NumImages int
	LastFetch string
}

// PrintMetadata prints the metadata of a webpage
func PrintMetadata(metadata Metadata) {
	printMutex.Lock()
	defer printMutex.Unlock()

	fmt.Println("site:", metadata.Site)
	fmt.Println("num_links:", metadata.NumLinks)
	fmt.Println("images:", metadata.NumImages)
	fmt.Println("last_fetch:", metadata.LastFetch)
	fmt.Println()
}
