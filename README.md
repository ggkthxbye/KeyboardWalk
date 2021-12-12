# KeyboardWalk
- Small **Go** program to generate keyboard walks to quickly get any low hanging fruit while password cracking
- Not an all inclusive key walk list.
- Designed to be used with `john` (JtR) by passing the wordlist via stdin
- **Example:** *`./KeyboardWalk -op=4x4 | john --stdin hashes.txt`*

## Build
```
make build
```
## Run
```
└─$ ./KeyboardWalk -h
./KeyboardWalk is a small Go program to generate keyboard walks to quickly get any low hanging fruit while password cracking.
Usage of ./KeyboardWalk:
  -op string
        Optionally, explicitly specify keywalk operation to perform. [4x4, 3x4, 4x5, 3x5, 4x3] (default "all")
Example: 
        ./KeyboardWalk -op=4x3 | john --stdin hashes.txt   
```
```
┌──(kali㉿kali)-[~/projects/KeyboardWalk/bin]
└─$ ./KeyboardWalk
1qaz2wsxZXCV$#@!
1qaz2wsxZXCVvcxz
1qaz2wsxZXCVfdsa
1qaz2wsxZXCVrewq
1qaz2wsxZXCV4321
1qaz2wsxZXCVVFR$
1qaz2wsxZXCVCDE#
1qaz2wsxZXCVXSW@
...
```
```
./KeyboardWalk -op=4x4 | john --stdin hashes.txt
```
## ME
**Author: Jason Glenn**  
[https://github.com/cyk1k]  
[https://www.linkedin.com/in/ggkthxbye/]