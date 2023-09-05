/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	code bool
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new challenge directory with specified name",
	Long: `Creates a new directory with the specified name and establishes the prerequisite development
	and testing files required to write and test a given challenge function.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("Name is required during execution - `go-challenge add <name>`")
			os.Exit(1)
		}

		// Only processing one name argument...
		name := args[0]

		// Make the directory for use
		// TODO: What permissions are best?
		err := os.Mkdir(name, 0755)

		if err != nil {
			return fmt.Errorf("%w", err)
		}

		// Establish the version go to use via runtime
		// This feels needlessly complex
		goVer := runtime.Version()
		goVer = strings.Trim(goVer, "go")
		modVer := strings.SplitAfterN(goVer, ".", 3)
		modVer[1] = strings.Trim(modVer[1], ".")
		goVer = modVer[0] + modVer[1]

		modFile := []byte(fmt.Sprintf("module %v\n\ngo %v\n", name, goVer))
		err = os.WriteFile(name+"/go.mod", modFile, 0644)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		title := camelCase(name)

		mainFile := []byte(fmt.Sprintf("package main\n\nfunc %v(){\n\n}\n\nfunc main() {\n\t%v()\n}\n", title, title))
		err = os.WriteFile(name+"/main.go", mainFile, 0644)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		mainTestFile := []byte(fmt.Sprintf("package main\n\nimport \"testing\"\n\nfunc Test%v(t *testing.T) {\n\n}\n", cases.Title(language.AmericanEnglish, cases.NoLower).String(title)))
		err = os.WriteFile(name+"/main_test.go", mainTestFile, 0644)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		if code {
			cmd := exec.Command("code", "./"+name)
			_, err = cmd.Output()
			if err != nil {
				return fmt.Errorf("%w", err)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolVarP(&code, "code", "c", false, "Open vscode automatically")
}

// Starting point:
// https://stackoverflow.com/questions/70083837/how-to-convert-a-string-to-camelcase-in-go
// This does not actually create a camelcase title - need to separate string - capitalize - and combine
func camelCase(s string) string {

	// Replace all underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Replace all dashes with spaces
	s = strings.ReplaceAll(s, "-", " ")

	// Remove all characters that are not alphanumeric or spaces or underscores
	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

	// split on spaces
	strs := strings.Split(s, " ")
	var newstr []string

	for _, v := range strs {
		newstr = append(newstr, cases.Title(language.AmericanEnglish, cases.NoLower).String(v))
	}
	// Iterate over each item in the array
	s = strings.Join(newstr, "")

	// Lowercase the first letter
	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}
