package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func readVars(envDir string) (envVars []string, err error) {

	files, err := ioutil.ReadDir(envDir)
	if err != nil {
		return
	}

	for _, f := range files {
		var envVar []byte

		envVar, err = ioutil.ReadFile(path.Join(envDir, f.Name()))
		if err != nil {
			return
		}
		envVars = append(envVars, fmt.Sprintf("%s=%s", f.Name(), envVar))
	}

	return
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "go-envdir <envdir> <program>",
		Short: "Starts program with environment variables from envdir",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			envDir := args[0]
			program := args[1]

			envVars, err := readVars(envDir)

			if err != nil {
				log.Fatalf("Vars read error: %s\n", err)
			}

			c := exec.Command(program)
			c.Env = append(c.Env, envVars...)

			c.Stderr = os.Stderr
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin

			//err = c.Run()
			if err != nil {
				log.Fatalf("Execution error: %s\n", err)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
