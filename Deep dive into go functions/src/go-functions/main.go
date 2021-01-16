package main

import (
	"errors"
	"fmt"
	"go-functions/simplemath"
	"io"
	"math"
)

type MathExpr = string

type BadReader struct {
	err error
}

type SimpleReader struct {
	count int
}

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	if err := ReadFullFile(); err != nil {
		fmt.Printf("Something bad occured: %s\n", err)
	}
}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("A panic occured but is ok")
		} else if p != nil {
			panic("An unexpected error occured and we do not want to recover!")
		}

	}()

	defer func() {
		println("Before for-loop")
	}()

	for {
		value, readerErr := r.Read([]byte("Text that does nothing"))
		if err == io.EOF {
			println("Finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr

			return
		}

		println(value)
	}

	defer func() {
		println("after for-loop")
	}()

	return nil
}

func ReadSomethingBad() error {
	var r io.Reader = &BadReader{errors.New("My nonsense reader")}
	value, err := r.Read([]byte("Test something"))
	if err != nil {
		fmt.Printf("An error occurred %s", err)
		return err
	}

	println(value)

	return nil
}

func (br *BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

var errCatastrophicReader = errors.New("something catastrophic occured in the reader")

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count == 2 {
		panic(errors.New("Another error"))
	}

	if br.count > 3 {
		return 0, io.EOF
	}
	br.count++

	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("closing reader")

	return nil
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {

	case AddExpr:
		return simplemath.Add

	case SubtractExpr:
		return simplemath.Subtract

	case MultiplyExpr:
		return simplemath.Multiply

	default:
		return func(f, f2 float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}
