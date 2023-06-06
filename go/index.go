package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/pixelbin-dev/pixelbin-go/v2/sdk/platform"
)


func formatJSON(jsonString string) string {
	// Replace quotes around keys

	jsonString = strings.ReplaceAll(jsonString, `"`, "'")

	// Add indentation
	lines := strings.Split(jsonString, "\n")
	for i, line := range lines {
		lines[i] = "  " + line
	}
	indentedString := strings.Join(lines, "\n")

	// Replace : with =
	jsFormatted := strings.ReplaceAll(indentedString, ":", ":")

	return jsFormatted
}

func main() {
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )

    config.SetOAuthClient()


    pixelbin := platform.NewPixelbinClient(config)


    params := platform.GetTransformationContextXQuery{
        URL: "IMAGE_URL.png",
    }
    result, err := pixelbin.Transformation.GetTransformationContext(params)

    if err != nil {
        fmt.Println(err)
    }


	jsonBytes, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling map to JSON:", err)
		return
	}

	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}