clean:
	rm -f chp2

build: clean
	go build

help: build
	./chp2 -h

test-root: build
	./chp2 --files fixtures/a.txt,fixtures/b.txt,fixtures/c.txt
	./chp2 --pattern "first" --files fixtures/a.txt,fixtures/b.txt,fixtures/c.txt
	./chp2 --pattern "first" --dir fixtures

test-catrand: build
	./chp2 catrand fixtures/a.txt fixtures/b.txt fixtures/c.txt

test-catrand-error: build
	./chp2 catrand fixtures/a.txt fixtures/b.txt fixtures/c.txt fixtures/d.txt

test-grep: build
	./chp2 grepfiles "first" fixtures/a.txt fixtures/b.txt fixtures/c.txt
	./chp2 grepfiles "second" fixtures/a.txt fixtures/b.txt fixtures/c.txt
	./chp2 grepfiles "first and only" fixtures/a.txt fixtures/b.txt fixtures/c.txt
	./chp2 grepfiles "THIS DNE" fixtures/a.txt fixtures/b.txt fixtures/c.txt

test-grep-error: build
	./chp2 grepfiles "first" fixtures/a.txt fixtures/b.txt fixtures/c.txt fixtures/d.txt

test-grepdir: build
	./chp2 grepdir "third" fixtures
	./chp2 grepdir "" fixtures

test-grepdir-error: build
	./chp2 grepdir "third" non-existent-dir



# test-greprec: build
# 	./chp2 greprec "first" fixtures
# 	./chp2 greprec "second" fixtures
# 	./chp2 greprec "first and only" fixtures
# 	./chp2 greprec "THIS DNE" fixtures