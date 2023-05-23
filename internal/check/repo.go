package check

import (
	"bufio"
	// "log"
	"os"
	"regexp"
)

// IsRepo checks if the path is git repo
func IsRepo(path string) bool {

	_, err := os.Stat(path + "/.git")

	return err == nil
}

// ParseConfig - get branch, user, remotes
func ParseConfig(path string) (string, string, []string) {
	var text, branch, user string
	var remote []string

	file, err := os.Open(path + "/.git/config")
	IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()

		re, _ := regexp.Compile(`\[branch \"`)
		if re.FindString(text) != "" {
			branch = re.ReplaceAllString(text, "")
			re, _ = regexp.Compile(`\"\]`)
			branch = re.ReplaceAllString(branch, "")
		}
		re, _ = regexp.Compile(`name =`)
		if re.FindString(text) != "" {
			user = re.ReplaceAllString(text, "")
		}
		re, _ = regexp.Compile(`\[remote \"`)
		if re.FindString(text) != "" {
			text = re.ReplaceAllString(text, "")
			re, _ = regexp.Compile(`\"\]`)
			text = re.ReplaceAllString(text, "")
			remote = append(remote, text)
		}
	}

	// log.Println("BRANCH =", branch, "USER =", user, "REMOTE =", remote)

	return branch, user, remote
}

// Branch - returns current git branch
func Branch(path string) string {
	var branch string

	file, err := os.ReadFile(path + "/.git/HEAD")
	IfError(err)

	re, _ := regexp.Compile(`ref: refs\/heads\/`)
	branch = re.ReplaceAllString(string(file), "")

	return branch
}
