package directory

import (
	"testing"
)

func TestTree(t *testing.T) {
	mainNode := CreateTree().AddSubNode("1", "")

	last := mainNode.AddSubNode("2", "2").AddSubNode("3", "4")

	root := last.GetRoot()

	if mainNode.Id != root.Id {
		t.Errorf("wrong root node")
	}
	path := last.Path()

	if path != "1->2->3" {
		t.Errorf("wrong path (%s) : must be 1->2->3", path)
	}

}
