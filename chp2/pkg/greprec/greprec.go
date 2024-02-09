package greprec

import (
	"bufio"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type grepContext struct {
	wg         sync.WaitGroup
	ctx        context.Context
	cancelFunc context.CancelFunc
	cancelErr  error
	errMutex   sync.Mutex
}

func grepFile(file string, pattern string, cancel *grepContext) {
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

func GrepRec(directory string, pattern string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return fmt.Errorf("directory %s does not exist: %w", directory, err)
	}

	fileSystem := os.DirFS(directory)

	paths := make([]string, 0)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory %s: %w", directory, err)
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	ctx2, cancelFunc := context.WithCancel(context.Background())
	ctx := &grepContext{
		wg:         sync.WaitGroup{},
		ctx:        ctx2,
		cancelFunc: cancelFunc,
	}

	for _, path := range paths {
		ctx.wg.Add(1)
		go grepFile(filepath.Join(directory, path), pattern, ctx)
	}

	go func() {
		ctx.wg.Wait()
		ctx.cancelFunc()
	}()

	<-ctx.ctx.Done()
	return ctx.cancelErr
}
