package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type command struct {
	dir    string
	name   string
	args   []string
	msg    string
	errMsg string
}

func executeCmdInKataDir(cmd command) {
	initCmd := exec.Command(cmd.name, cmd.args...)
	initCmd.Dir = cmd.dir

	if cmd.msg != "" {
		fmt.Println(cmd.msg)
	}

	err := initCmd.Run()
	if err != nil {
		fmt.Println(cmd.errMsg+":", err)
	}
}

func initGit(targetDir string, ignores []string) {

	executeCmdInKataDir(command{
		dir:    targetDir,
		name:   "git",
		args:   []string{"init"},
		msg:    "Initialising git with initial commit...",
		errMsg: "error initalising git",
	})

	createGitignore(targetDir, ignores)

	executeCmdInKataDir(command{
		dir:    targetDir,
		name:   "git",
		args:   []string{"add", "."},
		errMsg: "error adding files to git",
	})

	executeCmdInKataDir(command{
		dir:    targetDir,
		name:   "git",
		args:   []string{"commit", "-m", "Initial commit"},
		errMsg: "error committing files",
	})
}

func createGitignore(targetDir string, ignores []string) {
	gitIgnore, _ := os.Create(path.Join(targetDir, ".gitignore"))

	defer gitIgnore.Close()

	for i, ignore := range ignores {

		if i != 0 {
			ignore = "\n" + ignore
		}

		bs := []byte(ignore)
		gitIgnore.Write(bs)
	}

}

func convertToCamelCase(s string) string {
	isUpper := false

	var r string
	letters := strings.Split(s, "")

	for _, l := range letters {

		if l == " " {
			isUpper = true
			continue
		}

		if isUpper {
			r += strings.ToUpper(l)
			isUpper = false
		} else {
			r += strings.ToLower(l)
		}
	}
	return r
}

func convertLowerCamelCaseToUpper(s string) string {
	bs := []byte(s)
	firstChar := string(bs[0])
	return strings.Replace(s, firstChar, strings.ToUpper(firstChar), 1)
}

func createKataName(args []string) string {
	if len(args) == 0 {
		return promptName()
	}
	return convertToCamelCase(args[0])
}

func promptName() string {
	fmt.Print("Enter kata name: ")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n := s.Text()

	if len(n) == 0 {
		return promptName()
	}

	return convertToCamelCase(n)
}
