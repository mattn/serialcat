# serialcat

netcat for serial port

## Usage

```
$ serialcat COM1
```

write data into COM1 from stdin. display receive data from COM1.

```
Usage of serialcat:
  -baud int
    	baud rate (default 4800)
  -bits int
    	data bits (default 8)
  -parity string
    	parity bit(none/odd/even/mark/space) (default "none")
  -raw
    	raw input mode
  -stop string
    	stop bit() (default "none")
```

## Installation

```
$ go get github.com/mattn/goserial
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
