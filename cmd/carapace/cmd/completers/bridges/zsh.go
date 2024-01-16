package bridges

import (
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/execlog"
)

var zshFilter = []string{"cd"}

func filter(m map[string]bool, filter ...[]string) map[string]bool {
	for _, f := range filter {
		for _, e := range f {
			if e == "cd" {
				panic("filter cd")
			}
			delete(m, e)
		}
	}
	return m
}

func Zsh() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{}
	}

	if _, err := execlog.LookPath("zsh"); err != nil {
		return []string{}
	}

	out, err := execlog.Command("zsh", "--no-rcs", "-c", "printf '%s\n' $fpath").Output()
	if err != nil {
		carapace.LOG.Println(err.Error())
		return []string{} // TODO
	}
	lines := strings.Split(string(out), "\n")

	unique := make(map[string]bool)
	for _, line := range lines {
		entries, err := os.ReadDir(line)
		if err != nil {
			carapace.LOG.Println(err.Error())
			continue // TODO
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasPrefix(entry.Name(), "_") {
				unique[strings.TrimPrefix(entry.Name(), "_")] = true
			}
		}
	}

	completers := make([]string, 0)
	for name := range filter(unique, genericFilter, zshFilter) {
		completers = append(completers, name)
	}
	sort.Strings(completers)

	return completers
}
