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
	if err := ReadFullFile(); err == io.EOF {
		println("Successfully read file")
	} else if err != nil {
		println("Something bad occured")
	}
}

func ReadFullFile() error {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
	}()

	defer func() {
		println("before for-loop")
	}()

	for {
		value, err := r.Read([]byte("Text that does nothing"))
		if err == io.EOF {
			println("Finished reading file, breaking out of loop")
			break
		} else if err != nil {
			return err
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

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	// if br.count == 2 {
	// 	panic("something catastrophic occured in the reader")
	// }

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
