package tree_sitter_logscale_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_logscale "github.com/jakewilliami/tree-sitter-logscale/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_logscale.Language())
	if language == nil {
		t.Errorf("Error loading Logscale grammar")
	}
}
