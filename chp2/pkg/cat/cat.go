package cat

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type doneChan chan string

func printFile(file string, finish doneChan, cancelCauseFunc context.CancelCauseFunc) {
	// Open the file, read every line, printing each line to the console
	// When done, close the file and send the file name to the finish channel
	osFile, err := os.Open(file)
	if err != nil {
		cancelCauseFunc(fmt.Errorf("failed to open file %s: %w", file, err))
		return
	}
	defer osFile.Close()

	scanner := bufio.NewScanner(osFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	finish <- osFile.Name()
}

func CatRand(files []string) error {
	finish := make(doneChan, len(files))
	ctx, cancelCauseFunc := context.WithCancelCause(context.Background())

	for _, file := range files {
		go printFile(file, finish, cancelCauseFunc)
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
