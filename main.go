package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Field struct {
	Height, Width int
	Field         [][]byte
}

func (f Field) Show() {
	for _, str := range f.Field {
		for _, val := range str {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
	fmt.Println("\n")
}

func Init(h, w int) *Field {
	f := Field{
		Height: h,
		Width:  w,
		Field:  make([][]byte, h),
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < h; i++ {
		f.Field[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			if rand.Intn(4) == 0 {
				f.Field[i][j] = '#'
				continue
			}
			f.Field[i][j] = '.'
		}
	}
	return &f
}

type Pair struct {
	V byte
	N int
}

func (f *Field) Neighbours() [][]Pair {

	tempField := make([][]Pair, f.Height)
	for i := range f.Field {
		tempField[i] = make([]Pair, f.Width)
		for j := range f.Field[i] {
			res := 0
			if j+1 < f.Width && f.Field[i][j+1] == '#' {
				res++
			}
			if j-1 >= 0 && f.Field[i][j-1] == '#' {
				res++
			}
			if i-1 >= 0 && f.Field[i-1][j] == '#' {
				res++
			}
			if i+1 < f.Height && f.Field[i+1][j] == '#' {
				res++
			}
			tempField[i][j] = Pair{
				V: f.Field[i][j],
				N: res,
			}
		}
	}
	return tempField
}

func (f *Field) Next() {
	TF := f.Neighbours()
	for i := range f.Field {
		for j := range f.Field[i] {
			switch {
			case TF[i][j].V == '#' && (TF[i][j].N < 2 || TF[i][j].N > 3):
				f.Field[i][j] = '.'
			case TF[i][j].V == '.' && TF[i][j].N == 3:
				f.Field[i][j] = '#'
			}
		}
	}
}

func main() {
	f := Init(10, 10)
	f.Show()
	for i := 0; i < 10; i++ {
		f.Next()
		f.Show()
		time.Sleep(3)
	}
}
