package grep

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

type doneChan chan string

func grepFile(file, pattern string, finish doneChan, cancelCauseFunc context.CancelCauseFunc) {
	// Open the file, read every line searching for the pattern.
	// Print the line number and the file of the matches, if any
	osFile, err := os.Open(file)
	if err != nil {
		cancelCauseFunc(fmt.Errorf("failed to open file %s: %w", file, err))
		return
	}
	defer osFile.Close()

	scanner := bufio.NewScanner(osFile)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Printf("match [%s] in %s @ L%d\n", pattern, file, lineNumber)
		}
		lineNumber++
	}

	finish <- osFile.Name()
}

func Grep(files []string, pattern string) error {
	finish := make(doneChan, len(files))
	ctx, cancelCauseFunc := context.WithCancelCause(context.Background())

	for _, file := range files {
		go grepFile(file, pattern, finish, cancelCauseFunc)
	}

	for range len(files) {
		select {
		case <-finish:
			// fmt.Printf("Finished file: %s\n", finishedFile)
		case <-ctx.Done():
			fmt.Printf("Cancelled with cause: %s\n", context.Cause(ctx))
			return ctx.Err()
		}
	}
	return nil
}
