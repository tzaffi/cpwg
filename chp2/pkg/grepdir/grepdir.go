package grepdir

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type cancellation struct {
	wg         sync.WaitGroup
	ctx        context.Context
	cancelFunc context.CancelFunc
	cancelErr  error
	errMutex   sync.Mutex
}

func grepFile(file string, pattern string, cancel *cancellation) {
	// Open the file, read every line searching for the pattern.
	// Print the line number and the file of the matches, if any
	defer cancel.wg.Done()

	// not blocking as worse case, we just go on with the next file:
	if cancel.cancelErr != nil {
		return
	}

	osFile, err := os.Open(file)
	if err != nil {
		cancel.errMutex.Lock()
		cancel.cancelErr = fmt.Errorf("failed to open file %s: %w", file, err)
		cancel.errMutex.Unlock()
		cancel.cancelFunc()
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
}

func GrepDir(directory string, pattern string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", directory, err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	cancel := &cancellation{
		wg:         sync.WaitGroup{},
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		cancel.wg.Add(1)
		go grepFile(filepath.Join(directory, file.Name()), pattern, cancel)
	}

	go func() {
		cancel.wg.Wait()
		cancel.cancelFunc()
	}()

	<-ctx.Done()
	return cancel.cancelErr
}
