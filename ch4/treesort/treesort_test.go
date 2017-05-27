package treesort

import (
	"fmt"
	"testing"
)

func TestTreeString(t *testing.T) {
	root := tree{value: 4}
	add(&root, 1)
	add(&root, 2)
	add(&root, 3)
	add(&root, 5)
	add(&root, 6)
	add(&root, 7)
	fmt.Printf("%s", &root)
}
