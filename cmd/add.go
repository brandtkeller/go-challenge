/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	code bool
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new challenge directory with specified name",
	Long: `Creates a new directory with the specified name and establishes the prerequisite development
	and testing files required to write and test a given challenge function.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Name is required during execution - `go-challenge add <name>`")
			os.Exit(1)
		}
		run(args[0])
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

func run(name string) {
	// Make the directory for use
	err := os.Mkdir(name, 0755)
	check(err)

	// Establish the version go to use via runtime
	// This feels needlessly complex
	goVer := runtime.Version()
	goVer = strings.Trim(goVer, "go")
	modVer := strings.SplitAfterN(goVer, ".", 3)
	modVer[1] = strings.Trim(modVer[1], ".")
	goVer = modVer[0] + modVer[1]

	modFile := []byte(fmt.Sprintf("module %v\n\ngo %v\n", name, goVer))
	err = os.WriteFile(name+"/go.mod", modFile, 0644)
	check(err)

	title := strings.Title(name)

	mainFile := []byte(fmt.Sprintf("package main\n\nfunc %v(){\n\n}\n\nfunc main() {\n\n}\n", title))
	err = os.WriteFile(name+"/main.go", mainFile, 0644)
	check(err)

	mainTestFile := []byte(fmt.Sprintf("package main\n\nimport \"testing\"\n\nfunc Test%v(t *testing.T) {\n\n}\n", title))
	err = os.WriteFile(name+"/main_test.go", mainTestFile, 0644)
	check(err)

	if code {
		cmd := exec.Command("code", "./"+name)
		_, err = cmd.Output()
		check(err)
	}

}
