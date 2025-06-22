package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func removeDescriptions(node *yaml.Node) {
	// Always recurse into child nodes for all kinds
	switch node.Kind {
	case yaml.DocumentNode:
		for _, n := range node.Content {
			removeDescriptions(n)
		}
	case yaml.MappingNode:
		var newContent []*yaml.Node
		for i := 0; i < len(node.Content); i += 2 {
			k := node.Content[i]
			v := node.Content[i+1]
			if k.Value == "description" {
				continue
			}
			removeDescriptions(v)
			newContent = append(newContent, k, v)
		}
		node.Content = newContent
	case yaml.SequenceNode:
		for _, v := range node.Content {
			removeDescriptions(v)
		}
	}
}

func main() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	var node yaml.Node
	if err := yaml.Unmarshal(in, &node); err != nil {
		fmt.Fprintln(os.Stderr, "yaml unmarshal error:", err)
		os.Exit(1)
	}

	removeDescriptions(&node)

	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	if err := enc.Encode(&node); err != nil {
		fmt.Fprintln(os.Stderr, "yaml encode error:", err)
		os.Exit(1)
	}
}
