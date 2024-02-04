package main

import (
	"fmt"
	"strings"

	"github.com/PanosVasilopoulos92/learn-to-code-go/sec34-testing-and-benchmarking/251-benchmark-examples/cat/mystr"
)

const s = `Sunlight whispers through the leaves,
A gentle dance on dappled ground,
Where winter's hush begins to leave,
And spring's first hints are softly found.

A robin's song, a joyful trill,
A melody that fills the air,
As snowdrops peek, defying chill,
With promises of blossoms fair.

A butterfly, on painted wings,
Flitting free on breezes warm,
A splash of color that life brings,
Transforming nature's sleeping form.

So let us breathe the air so sweet,
And feel the sun upon our face,
Embrace the hope that winter's beat,
Cannot forevermore erase.

For even in the depths of cold,
A spark of warmth remains below,
And soon, a story will unfold,
Where beauty starts to gently grow.
`

var xs []string

func main() {
	xs = strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Println("--------")

	// fmt.Println(strings.Contains(s, "story"))
	// fmt.Println(strings.ToTitle(s))
	// fmt.Println(strings.Clone(s))

	fmt.Printf("\n%s\n", mystr.Cat(xs))

	fmt.Printf("\n%s\n", mystr.Concat(xs))

}
