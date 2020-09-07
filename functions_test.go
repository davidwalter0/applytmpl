package applytmpl

import (
	"fmt"
	"reflect"
	"testing"
)

type Result struct {
	want []string
	got  []string
}

func Test_GenerateChar(t *testing.T) {
	fmt.Printf("ok\n")

	var compareResult []Result = []Result{
		Result{want: []string{}, got: GeneratorChar()},
		Result{want: []string{"a", "b", "c", "d", "e"}, got: GeneratorChar(5)},
		Result{want: []string{"b", "c", "d", "e", "f", "g", "h", "i", "j"}, got: GeneratorChar(1, 9)},
		Result{want: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"}, got: GeneratorChar(15)},
		Result{want: []string{"b", "e", "h", "k"}, got: GeneratorChar(1, "10", 3)},
		Result{want: []string{"a", "f", "k", "p", "u", "z"}, got: GeneratorChar(0, 26, 5)},
		Result{want: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, got: GeneratorChar("a", "z")},
	}

	for _, result := range compareResult {
		if result.want != nil && result.got != nil && !reflect.DeepEqual(result.want, result.got) {
			t.Errorf("wanted %v got %v", result.want, result.got)
		}
	}
	// fmt.Println(GeneratorChar())
	// fmt.Println(GeneratorChar(5))
	// fmt.Println(GeneratorChar(1, 8))
	// fmt.Println(GeneratorChar(15))
	// fmt.Println(GeneratorChar(1, "9", 3))
	// fmt.Println(GeneratorChar(0, 25, 5))
	// fmt.Println(GeneratorChar("a", "z"))
}

func Test_Generate(t *testing.T) {
	fmt.Printf("ok\n")
	fmt.Println(Generator())
	fmt.Println(Generator(uint32(5)))
	fmt.Println(Generator(5))
	fmt.Println(Generator(1, 8))
	fmt.Println(Generator(15))
	fmt.Println(Generator(1, "9", 3))
	fmt.Println(Generator(0, 25, 5))
	fmt.Println(GeneratorChar("a", "f", 2))
}

func Test_Ascii(t *testing.T) {
	a := Generator(3)
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a, 'a'+i, uint8('a'+i), string('a'+i))
	}
}
