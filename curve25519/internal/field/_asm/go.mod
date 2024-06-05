module asm

go 1.18

require (
	github.com/mmcloughlin/avo v0.5.0
	golang.org/x/crypto v0.24.0
)

require (
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
)

replace golang.org/x/crypto => ../../../..
