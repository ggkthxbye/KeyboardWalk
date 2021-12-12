#
#	KeyboardWalk
#	 - Small Go program to generate keyboard walks to quickly get any low hanging fruit while password cracking
#	 - Not an all inclusive key walk list.
#	 - Designed to be used with `john` by passing the wordlist via stdin
#	 - Example: ./KeyboardWalk -op=4x4 | john --stdin hashes.txt
#
#	 Author: Jason Glenn
#	 https://github.com/cyk1k
#	 https://www.linkedin.com/in/ggkthxbye/
#
#    make build
#	 bin/KeyboardWalk
#




build: clean
	go build -o bin/KeyboardWalk main.go

run: build
	bin/main

clean:
	rm -f bin/*
