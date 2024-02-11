# Main Points
**1. Statically Typed Language**
```
n = 1
n = "string"

> Error!
```

```
var Name type
var Name = "knownType"  // string
var Number = 1          // integer
```

**2. Strongly Typed Language**
```
a = 1
b = "two"
c = a + b

> Error!
```

**3. Compiled**
- Compile before running the program, unlike the language that needs an interpreter like Python
- Points 1 + 2. The Compiler can check errors in the code during the Compile time
- code -> 1010101011101010101 executable program .exe
- Make GO faster e.g. 100 times

**4. Fast Compile Time**
- make it almost like hot-reload or restarting a Python server

**5. Built-in Concurrency**
- Goroutine

**6. Simplicity**
- avoid unnecessary variables and imports
- garbage collector

---
# Installation
`https://go.dev/doc/install`
- command `go ...`
- Add binary `PATH` for your console: 
  - `PATH=$PATH:/usr/local/go/bin`
  - The PATH environment variable is a list of directories that your operating system searches through to find executable files when you type a command in the terminal or command prompt.
  - Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.
- `GOPATH`
  - GOPATH is an environment variable that specifies the location of your Go workspace. The Go workspace typically has three subdirectories:
    - The `src` directory contains your Go source code files
    - the `pkg` directory contains compiled package files
    - the `bin` directory contains executable binaries.
  - When you compile or run Go programs, the Go toolchain looks in the directories specified by GOPATH to find the necessary source files and packages.
  - Multiple workspace directories can be specified in the GOPATH, separated by the system's path separator 
    (such as a colon`:` on Unix-like systems or a semicolon`;` on Windows).
- `$GOBIN`
  - The go install command installs binaries to $GOBIN, which defaults to $GOPATH/bin.

---
# GO Project
**1. initialize Go project**
```
- module/
  - package/
  - package/
  - package/
    - file.go
    - file.go
    - file.go
```
Initialize a package management (Module)
```
(workspace): mkdir github.com/user/module
(workspace): cd github.com/user/module

(workspace/module): go mod init

(workspace/module): ls
go.mod
```
Inside go module file contains `name`, `go version`, and installed packages.
```
# go.mod
module github.com/user/module

go 1.20
```
**2. Go find main package -> main()**

Create main.go, go will try to find the entry point of the program which will look at 
the `main` package and the `main()` function.
```
# shell
(workspace/module): mkdir -p cmd/basics/
(workspace/module): cd cmd/basics/

(workspace/module): touch main.go

x---------------------------------------x
package main 

func main() {
    // run code here
}
x---------------------------------------x
```
Try to build the main.go file into an executable file
```
# shell
(workspace/module): go build cmd/basics/main.go

(workspace/module): ls
cmd     go.mod      main

(workspace/module): ./main
Hello World!
```
To run go build and executable file in a single command
```
(workspace/module): go run cmd/basics/main.go
Hello World!
```

---
# Const, Var, Types
```
* const C
* var V

* bool
* int   int16   int32   int64
* uint  uint8   uint16  uint32
* float32 (1+8+23   single-precision  3.141 593)
* float64 (1+11+52  double-precision  3.141 592 741 012 573 242)
* rune  string
```

**Bits size in memory can store a larger value**
```
var num int  ==> 32 | 64 bits depends on the os
var num int8 ==> (-128, 0-127)
var num uint8 ==> (0,   255)
```

**Number Overflow Error**

Build time overflow
```
var int8 = 127 + 1 

<caution!> Overflow Error!
```

Run-time unexpectation
```
var int8 = 127
int8 = int8 + 1

<caution!> Weird result = -127

 127 = 0 111 1111
   1 = 0 000 0001  
-----------------
-128 = 1 000 0000
```

**Floating precision**
* use `32 bits` where precision isn't critical
* use `64 bits` for scientific calculations, financial applications, or any domain where precision is crucial.
```
var num float32 = 12345678.9 >> 12345679.000 000
var num float64 = 12345678.9 >> 12345678.900 000

float32 = 1.0 / 3.0 >> 0.33333334326744079590
float64 = 1.0 / 3.0 >> 0.33333333333333331483
```

**Choosing the right data type**
```
    0-255  0-255  0-255
rgb(uint8, uint8, uint8)
```

**Operations: + - * / %**
- Cannot be a different type, we need to cast
```
floatNum32 + float32(intNum32)
```

- int will round it down
```
fmt.Print(3/2)          >> 1
fmt.Print(float32(3/2)) >> 1  // float32(int(1))
```

**UTF8 Encode**
- string (collection of runes)
```
var my string = "one-line"
var my string = `multiple-
line`
var my string = "can" + " " + "concat"

len() will count number of byte (not char)
len("a")      = 1         1 bytes -> utf8
len("ABCD")   = 4         4 bytes -> utf8
len("δ")      = 2 !       2 bytes -> utf8

If you want to count the string
utf8.RuneCountInString("δ")  = 1
```

**Rune (represents character)**
```
var myRune rune = 'a'
myRune          = 97
```

**Default values**
```
var myString string           >> ""
var myBool bool               >> false
var myNumber int|float|rune   >> 0
```

**Constants**
- Can't change
- Have to set immediately

---
# Functions
- privateFunction
- PublicFunction
- No overload
- No default arguments
- But can do this `fields ...types`
- Return multiple type `(type1, type2)`
  - var x, y int = makePoint(num, num)
- Design to return `error` type with default value type is `nil` otherwise panic!
  - `err = errors.New("Something went wrong")`

```
func privateFunc(input1 type) (outputType, error) {
    return ...
}
```

---
# Control flows
**If-else statement**
```
if err != nil {
  ...
} else if remainder == 0 {
  ...
} else {
  ...
}


if 1==1 && 2 == 2 || 3 != 3 {
  ...
}
```
**Switch**
- no need to break
```
switch{
  case err != nil:
    ...
  case remainder == 0:
    ...
  default:
    ...
}


switch remainder{
case 0:
  ...
case 1,2:  // 1 || 2
  ...
default:
  ...
}
```

---
# Collections
## Arrays: [3]int32
- Fixed length (Cannot change)
- Same type value
- Indexable (starts from 0)
- Contiguous in Memory
- The default value is `[]` with cap=0
```
// fixed length & same type value
[LENGTH]Type

var intArr = [3]int32            [0,0,0] as default
intArr := [3]int32{0, 0, 0}      self define
intArr := [...]int32{0, 0, 0}    AutoCount+self define


// Indexable
inArr[0]                  >> 0
inArr[1:3]                >> [0: 0]
inArr[0] = 123            >> [123, 0, 0]

int32                     >> 4 bytes (8*4=32)
[3]int3                   >> 12 bytes


// Contiguous in Memory
&intArr[0]                >> 0x14000122004 (hex)
&intArr[1]                >> 0x14000122008 (+4 bytes)
&intArr[2]                >> 0x1400012200c (+32 bits)

Accessing by index is predictable O(1)
```

## Slices: []int32
- Wraps array
- To give general and convenient interfaces
- Appending to the slice will create a new underlying array for you
- Need to reallocation if the size is exceeding the capacity (Slow operation), basically copy to a new address.
- `len()`  current member length in the slice
- `cap()`  capacity of the underlying array for this slice
- The default value is `[]` with cap=0

```
[]int32{1,2,3}      >> len=3, cap=3             <-- cap=underlying
append(mySlice,4)   >> len=4, cap=6 (whatever)  <-- re-allocate

                 |----capacity----|
underlying_array [1, 2, 3, 4, *, *]
                 |--length--|


// Use make()
                    len cap
slice := make([]int, 5, 10)


// Spread operator
slice2 := []int32{1, 2}
slice2 =  append(slice2, slice3...)

[1, 2, 3, 4, 5, 6]
```

## Map: map[string]int32
- key-value pair
- use a hash-table data structure
- store values in a bucket (a linked list)
- optional initialize capacity
- resizing if exceeding the capacity
- The default value is `nil`! Have to use `make()` instead
```
myMap := map[string]uint8{"key1": 1, "key2": 2}
myMap["Key1"]       >> 1
myMap["Key2"]       >> 2


// delete
delete(myMap, "Key1")


// Use make()
myMap := make(map[string]uint8)


// Default-Valued for unknown key
m := make(map[string]int)
m["key"]                  >> 0  (default value of int)

m, ok := m["key"]         >> 0, false


// Map Nil Error
var myMap map[string]int
myMap["5"] = 1

<Panic !>
A nil map doesn't point to an initialized map.
so attempting to read from or write to it will result 
in a runtime panic.
```

---
# Loop
**Iterate arrays and slices**
```
for index, val := range [3]int{4, 5, 6} {
  ...
}

$ go run main.go 
Index: 0, Value: 4 
Index: 1, Value: 5 
Index: 2, Value: 6
```

**Iterate a map**
```
notOrderedMap := map[string]uint8{
  "Adam":  23,
  "Sarah": 45,
}

for name, age := range notOrderedMap {
  fmt.Printf("Name: %v, Age: %v\n", name, age)
}

$ go run main.go 
Name: Adam, Age: 23
Name: Sarah, Age: 45

$ go run main.go 
Name: Sarah, Age: 45
Name: Adam, Age: 23
```

**A simple for-loop**
```
for i:=0; i<10; i++ {
    ...
}
```

**While Loop**
```
for i<10 {
    ...
}

for {
  break
}
```

**Performance matters**
```
// With pre-allocation is faster
slice1 := make([]int, 0, 1000000)  // pre-set capacity
t0 := time.Now()
for len(slice1) < 1000000 {
  slice1 = append(slice1, 1)
}
fmt.Println(time.Since(t0))
>> 2.41925ms


// Without pre-allocation is faster
slice1 := []int{}
t0 := time.Now()
for len(slice1) < 1000000 {
  slice1 = append(slice1, 1)
}
fmt.Println(time.Since(t0))
>> 801.750625ms
```

---
# Float
- Floating is non-trivial

---
# Strings, Runes, Bytes
## Strings
- Represents characters
- Is an array of bytes (number)
- Use encoding-decoding (mapping value in the table)
- Indexable (But you will get a number)
- Immutable, Cannot modify it e.g. myStr[0] = "x"
- The default value is `""`
```
myString := "rèsumè"

myString[0]           %v >> 114,  %T >> uint8

// Iterate byte array
        decoded decoded
utf-8   value   binary    encoded
r       114     1110010   [0]1110010
è       233     11101000  [110]000011 [10]101000
s       115     1110011
u       117     1110101
m       109     1101101
è       233     11101000
```

**ASCII Table**
- 7 bits (unsigned)
- Support 127 chars
```
0 null
8 backspace
10 new line
33 !
48 0
65 A
97 a
126 ~
127 DEL
```

**Support more characters problem?**
- I want more bits
- UTF-32 (32-bits) can support 1,114,112 chars
- But it wasted a lot of memory e.g.
```
Character     UTF-32
a             00000000 00000000 00000000 01100001
              |-------- wasted --------|

smile emoji   00000000 00000001 11110110 00001010
```

**UTF-8**
- support ASCII
- support many country characters
- support emoji representative
- group the same characters as the same code (Unicode)
- solve the issue by allowing dynamic bytes, depending on the char.

```
// Dynamically bytes (1-4)
Character     UTF-32
a             01100001
smile emoji   11110000 10011111 10011000 10001010


// How?
Code point <-> UTF-8 conversion

startsByte1   range      Byte1   Byte2   Byte3   Byte4
0             0-127      0x      -       -       -
110           128-2037   110x    10x     -       -
1110          x-65,535   1110x   10x     10x     -
11110         x-~1.1m    11110x  10x     10x     10x
                         ^- identifier

// Example

(char)    é
(base)    233                       <-- encoded
          ^- between 128-2037, Use pattern '110xxxxx 10xxxxxx'
(binary)  11101001 => 000011|101001 <-- encoding
(utf-8)   [110]xxxxx  [10]xxxxxx    <-- decoding
          [110]00011  [10]101001
          1100011 10101001          >> lookup = 'é'


// So, iterating string is iterating over the byte array
// and it does decode the character into the value for us

char   find mapped val     decode      base10
'é' -> 1100011 10101001 -> 11101001 -> 233

ré = [(1-4 bytes), (1-4 bytes)]   // actual binary
   = [01110010, 11000011]         // decoded binary
   = [114, 233]                   // decoded base10
   = [U+0072, U+00E9]             // unicode point
   = [0x72, 0xE9]                 // hex
   = 0x72E9                       // concatenate

résume = 0x72E973756D65


// len()
// len() is couting the member of array,
// in this case "string" is a byte array.

where "rèsumè" have 6 characters
len("rèsumè") = 8 
        mean  = 8 bytes
                because "string" is a byte array
              = [1, 1, 1, 1, 1, 1, 1, 1]
                 r  | è | s  u  m  | è |
```

## Runes
- rune is an alias for int32, which represents a Unicode code point. (encode)
- rune itself is a 32-bit integer, the actual memory representation of a character in a Go string can vary depending on its Unicode code point and the encoding used as UTF-8 encoding
- Immutable, it is used for mapping purposes!
- **iterate over a string is a rune!**
- The default value is `0`
```
[]rune("résumé")    >> [114, 233, 115, 117, 109, 233]


// Iterating over a string as runes
str := "résumé"
for _, r := range str {
    fmt.Printf("%c ", r)
}
```

**Concatenate**

- Copying (appending) the content of both strings into the new string
- Creating new slices of runes every time of the loop (SLOW)
- Runtime needs to allocate memory for each intermediate string created.
- Involves memory allocation and deallocation, (costly in a loop)
- Creates multiple intermediate strings that need to be garbage-collected eventually.

```
var catString string
for _, v := range []string{"h", "e", "l", "l", "o"} {
  catString += v
}

v-------------- slow operation
{"h"}
{"h", "e"}
{"h", "e", "l"}
{"h", "e", "l", "l"}
{"h", "e", "l", "l", "0"}
```

**Use strBuilder instead of + for concatenation**
- 'Strings.Builder', on the other hand, pre-allocates memory for the final string, reducing the need for frequent memory allocations.
- Minimizes the number of allocations and reallocations by dynamically resizing the buffer as needed.
- Reduced Garbage Collection Overhead.
```
var strBuilder strings.Builder
for _, v := range []string{"h", "e", "l", "l", "o"} {
  strBuilder.WriteString(strSlice[i])
}
fmt.Println(strBuilder.String())
```

---
# Struct
- Use to define your own type
- Can hold Mixed-type
- Mutable
- Default values are `children's default value`
```
type Person struct{
  name   string
  age    uint8
  father Person     <-- Mixed-type
}

// Default value
var myPerson Person
{
  name: ""         <-- Default string
  age:  0          <-- Default uint8
}
```

Initialize value
```
Person{name: "Pop", age: 30}
Person{"Pop", 30}

// set value
myPerson.name = "New Pop"   <-- Mutable
```

## Choosing Func Value vs. Pointer Receiver
- Use a value receiver if the method doesn't need to modify the struct and operates purely on its value. Value receivers are useful when you want to ensure that the original struct remains unmodified.
- Use a pointer receiver if the method needs to modify the struct or if you want to avoid copying the struct when calling the method, especially for large structs. Pointer receivers are useful when you want changes made within the method to be reflected in the original struct.

## Struct Composition
```
type Namable struct {
	name string
}

func (n *Namable) GetName() string {
	return n.name
}


type Person struct {
	age int
	int
	Namable
}


person := Person{10, 11, Namable{"test"}}
fmt.Print(person.int)
fmt.Print(person.name)
fmt.Print(person.GetName())
```

## Anonymous Struct
* cannot be reused
```
var myData := struct{
  x int
  y int
}{1, 2}
```

---
# Interface
**Problem to solve**
```
type gasEngine struct{
  ...
}

func (e gasEngine) milesLeft() uint8{
  ...
}
```
```
type electricEngine struct{
  ...
}

func (e electricEngine) milesLeft() uint8{
  ...
}
```
```
                 v---- too strict type!
func canMakeIt(e gasEngine, miles uint8){
  if miles <= e.milesLeft() {
    fmt.Println("You can make it there!")
  }else {
    fmt.Println("You need to fuel up first!")
  }
}
```

**Solution**
- Decoupling the type by its behavior
- We only need its behavior
```
type Engine interface{
  milesLeft() uint8     <-- signature
}

func canMakeIt(e Engine, miles uint8){
}
```

---
# Pointer
* Where the value itself is a memory address
* pointer's value is 32 bits - 64 bits depending on the OS
* We cannot set the value into the nil memory pointer: 'panic: runtime error: invalid memory address or nil pointer dereference.'
* `&obj`: memory address value to the object
* `*obj`: dereference (the memory address value) to get/set the object 
* `*int32`: typing of pointer that is pointing to int32 memory address
* Slice is a pointer to underlying array

## TLDR: Example
**Style 1. Copy Value**
* To Retain immutable
* For a small object, A few copy cost
```
person := Person{age: 10}

var p Person = person
p.age = 9

fmt.Printf("%v", person)
>> {10}
```

**Style 2. Reference Pointer**
* To mutate object
* For a large object that's expensive cost
```
person := Person{age: 10}

var p *Person = &person
p.age = 9

fmt.Printf("%v", person)
>> {9}
```

**Reminder: Slice is a pointer to the underlying array**
* Copy the slice value, will copy the memory location to the underlying array!
```
var slice     = []int32{1,2,3}
var sliceCopy = slice         <--- copy
sliceCopy[2]  = 4

slice     >> [1 2 4]
sliceCopy >> [1 2 4]  !!!

// Look at the variable itself
%p &slice          >> 0x400000c018
%p &sliceCopy      >> 0x400000c030

// Look at the underlying array
%p slice           >> 0x400000e090
%p sliceCopy       >> 0x400000e090
%p &slice[0]       >> 0x40000a4000
%p &sliceCopy[0]   >> 0x40000a4000
```

**Passing the parameters**
* is copied value by default
```
var numArr = [2]int32{1, 2}

%p &num[0]        >> 0x400000e090
%p &num[1]        >> 0x400000e094


func sum(num [2]int32) {
    %p &num[0]        >> 0x400000e0b0
    %p &num[1]        >> 0x400000e0b4
}


func sum(num *[2]int32) {
    %p &num[0]        >> 0x400000e090
    %p &num[1]        >> 0x400000e094
}
```

## Create a Pointer from scratch
```
var i int32

x-----------------------x
variable  value   address
-------------------------
i         0       0x1b08
                  0x1b09
                  0x1b0a
                  0x1b0b  (int32 = 4 bytes)
------------------------
                  :
```

Example of Pointer in 64 bits OS (8 bytes)
```
var p *int32

x-----------------------x
variable  value   address
-------------------------
p         nil     0x1b00
                  0x1b01
                  0x1b02
                  0x1b03
                  0x1b04
                  0x1b05
                  0x1b06
                  0x1b07  (64 bits = 8 bytes)
------------------------
                  :
```

Assigning reference value to a pointer
```
var p *int32 = new(int32)

x-----------------------x
variable  value   address
-------------------------
p         0x1b0c  0x1b00
                  0x1b01
                  0x1b02
                  0x1b03
                  0x1b04
                  0x1b05
                  0x1b06
                  0x1b07  (64 bits = 8 bytes)
------------------------
i         0       :
                  0x1b0b
------------------------
          0       0x1b0c
                  0x1b0d 
                  0x1b0e 
                  0x1b0f   (int32 = 4 bytes)
```

Referencing & Dereferencing
```
p               >> 0x1b0c
*p              >> 0


// example
// set this value at the memory location 
// that p is pointing to
*p = 10
```

## Create a Pointer from the reference
* `&ref`
```
var i int32
var p *int32 = &i


x-----------------------x
variable  value   address
-------------------------
p         0x1b08  0x1b00 <-|
                  0x1b01   |
                  0x1b02   |
                  0x1b03   |
                  0x1b04   |
                  0x1b05   |
                  0x1b06   |
                  0x1b07   |
------------------------   |
i         0       0x1b08 <-|
                  0x1b09
                  0x1b0a
                  0x1b0b
------------------------
```

Dereference and set the value
* `*p`
```
*p = 10


x-----------------------x
variable  value   address
-------------------------
                  :
i         10      0x1b08
                  :
```

# Resources
* [Learn GOLang Fast](https://www.youtube.com/watch?v=8uiZC0l4Ajw)
