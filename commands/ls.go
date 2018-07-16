package commands

import (
	"fmt"
	b "github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/repository"
)

func runLsBug(repo repository.Repo, args []string) error {
	refs, err := repo.ListRefs(b.BugsRefPattern)

	if err != nil {
		return err
	}

	for _, ref := range refs {
		bug, err := b.ReadBug(repo, ref)

		if err != nil {
			return err
		}

		snapshot := bug.Compile()

		fmt.Printf("%s %s\t%s\n", bug.HumanId(), snapshot.Title, snapshot.Summary())
	}

	return nil
}

var lsCmd = &Command{
	Description: "Display a summary of all bugs",
	Usage:       "",
	RunMethod:   runLsBug,
}
