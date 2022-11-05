# go-ansi
ANSI escape sequences for Golang

## Install:

```
go get github.com/bit101/go-ansi
```

## Import:

```
import "github.com/bit101/go-ansi"
```

## Functions:

### Set foreground (text) color

```
ansi.SetFg(fg ansiColor) // see below for `ansiColor` type
```

Affects the color of all text rendered with `fmt.Print*` functions. This color will remain active until it is changed or reset.

### Set background color

```
ansi.SetBg(bg ansiColor) // see below for `ansiColor` type
```

Affects the background color of all text rendered with `fmt.Print*` functions. This color will remain active until it is changed or reset.

### Set bold

```
ansi.SetBold(bold bool)
```

Affects whether text will be rendered in bold or regular. This setting will remain active until it is changed or reset.

### Set Underline

```
ansi.SetUnderline(underline bool)
```

Affects whether text will be rendered underlined or not. This setting will remain active until it is changed or reset.

### Set all properties

```
ansi.SetAll(fg ansiColor, bg ansiColor, bold bool, underline bool)
```

Sets all of the above properties in one shot.

### ResetAll()

```
ansi.ResetAll()
```

Sets all of the above properties to their default values.

## Print functions

All of the below are "one-shot" function. They can only be used to set ONE of the following: foreground color, bold, underline.

The properties selected will only be in effect for the single print call. After the call, all properties will revert to what they were set to with the above function calls.

In addition to passing an actual color, such as `ansi.Red` or `ansi.Green`, these functions will accept `ansi.Bold` and `ansi.Underline`.

Setting the background color is not supported with these methods.


### Print

```
ansi.Print(color ansiColor, message ...any)
```

Behaves exactly like `fmt.Print` but takes an `ansiColor` as a first argument.

### Print line

```
ansi.Println(color ansiColor, message ...any)
```

Behaves exactly like `fmt.Println` but takes an `ansiColor` as a first argument.

### Print formatted

```
ansi.Printf(color ansiColor, message any, args ...any)
```

Behaves exactly like `fmt.Printf` but takes an `ansiColor` as a first argument.

## The ansiColor type

`ansiColor`s are a set of predefined values that encapsule text properties - mostly colors.

### Colors

The following `ansiColor`s define actual colors. They can be used in `ansi.SetFg`, `ansi.SetBg` or any of the `ansi.Print*` functions:

```
ansi.Black
ansi.DarkGray
ansi.Red
ansi.LightRed
ansi.Green
ansi.LightGreen
ansi.Brown
ansi.Yellow
ansi.Blue
ansi.LightBlue
ansi.Purple
ansi.LightPurple
ansi.Cyan
ansi.LightCyan
ansi.LightGray
ansi.White
```

Another special color value is:

```
ansi.Default
```

This can be used in all the same places and will set either the foreground or background color to its default value as defined within the terminal emulator / shell / colorscheme.

In addition, the following values can be used in any of the `ansi.Print*` functions and do what you would expect:

```
ansi.Bold
ansi.NotBold
ansi.Underline
ansi.NotUnderline
```
