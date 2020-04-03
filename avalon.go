package main
import (
	"fmt"
    "github.com/rwynn/monstache/monstachemap"
)
// a plugin to convert document values to uppercase
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document

	if tags, ok := doc["tags"].(map[string]interface{}); ok {
		strTags := ""
        for key := range tags {
			strTags = strTags + " " + key
		}
		doc["tags"] = strTags
    } else {
		fmt.Printf("record not a map[string]interface{}: %v\n", doc["tags"])
		doc["tags"] = nil
    }

    output = &monstachemap.MapperPluginOutput{Document: doc}
    return
}