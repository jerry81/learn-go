# go types

bool
string
int,int8,int16,int64
uint, uint8...64
uintptr - holds a memory address (64 bits on 64 bit machine)) - probably will not show up in practice

byte (uint8)

rune (int32, but used for characters)

float32 float64

complex64, complex128 - actually for complex numbers (imaginary plus real)

if unintialized most types default to 0, false, or "" (empty string)