package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

const (
	TemplateRoute = "/router.tmpl"
)

var routeCmd = &cobra.Command{
	Use:   "router",
	Short: "Create file router",
	Long:  "Create file router for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		ValidateExistOrCreateDirectory(PathRouter)

		var input = ""

		if len(args) >= 1 && args[0] != "" {
			input = args[0]
		}

		if len(input) == 0 {
			fmt.Println("Set name controller")
			return
		}
		path := strings.ToLower(input)
		runes := []rune(path)
		runes[0] = unicode.ToUpper(runes[0])
		name := string(runes)
		data := Data{
			ListName:   "Get" + name + "s",
			DetailName: "Get" + name,
			UpdateName: "Update" + name,
			DeleteName: "Delete" + name,
			Name:       name,
			Path:       path,
		}
		CreateRouter(name, data)

	},
}

func init() {
	rootCmd.AddCommand(routeCmd)
}

func CreateRouter(name string, data Data) {
	ValidateExistOrCreateDirectory(PathModel)
	ProcessTemplateString(
		RouterTemplate,
		PathRouter+name+".go",
		data)
	fmt.Printf("Success create router name %s\n", name)
}
