package gitwalk

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Revlist() {
	cmd := exec.Command("git", "rev-list", "--all")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		commitHash := scanner.Text()
		fmt.Println("for commit : " + commitHash)
		diffTree(commitHash)
		fmt.Println("------------------------------------------------------------------")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
