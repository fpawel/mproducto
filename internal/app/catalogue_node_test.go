package app

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReadCatalogueNodes(t *testing.T) {
	xs := ReadCatalogueNodes()
	b, _ := json.MarshalIndent(xs, "", "    ")
	fmt.Println(string(b))
}
