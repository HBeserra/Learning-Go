# Variables and data types

The `var` keyword is used to define a variable (memory location), the variable must be used or the compiler will return a error. 
- `_` or blank mutes the compiler error for unused variable, it is used for exemple when a method returns two values and one will not be used.
- `:=` is the shot declaration operator, used to declare variables without the var keywork. It can't be used for redeclare a variable.
  
```
var HoursInDay int = 24
// or 
var HoursInDay = 24
// or
HoursInDay := 24
```

## Basic data types
|Data Type| Size|Description |
|--:|--|:--|
|`bool`| 1 bit | true or false
|`int`| 32/64 bits*| Signed interger with default arch size
|`int8`| 8 bits| Signed 8-bit integers (-128 to 127) 
|`int16`| 16 bits| Signed 16-bit integers  (-32768 to 32767)
|`int32`| 32 bits| Signed 32-bit integers (-2147483648 to 2147483647) 
|`int64`| 64 bits| Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
|`uint`|32/64 bits*| Unsigned interger with default arch size
|`uint8`| 8 bits| Unsigned 8-bit integers (0 to 255)
|`uint16`| 16 bits| Unsigned 16-bit integers (0 to 65535)
|`uint32`| 32 bits| Unsigned 32-bit integers (0 to 4294967295)
|`uint64`| 64 bits| Unsigned 64-bit integers (0 to 18446744073709551615)
|`float32`| 32 bits| IEEE-754 32-bit floating-point numbers
|`float64`| 64 bits| IEEE-754 64-bit floating-point numbers
|`complex64`| 64 bits | Complex numbers with float32 real and imaginary parts
|`complex128`| 128 bits | Complex numbers with float64 real and imaginary parts
|`byte`|8 bits| same as `uint8`
|`rune`|32 bits| same as `uint32`
|`uintptr`| - | unsigned integer to store the uninterpreted bits of a pointer

*\* depends of the cpu architecture*

## Strings

In go strings are sequences of bytes coded in UTF-8, in utf-8 not all bytes are visible so the 3 byte in a string its not necessary the 3 caractere.

## Custom data types 

You can use structs for custom data types and using OOP create methods for this new data type.   

```
type Color struct{
  R       uint8
  G       uint8
  B       uint8
}

func (c Color) Luminance() uint8 {
  return (0.2126*R + 0.7152*G + 0.0722*B)
}
```

