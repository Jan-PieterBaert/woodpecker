package repo

import (
	"os"
	"text/template"

	"github.com/laszlocph/drone-oss-08/cli/drone/internal"
	"github.com/urfave/cli"
)

var repoSyncCmd = cli.Command{
	Name:      "sync",
	Usage:     "synchronize the repository list",
	ArgsUsage: " ",
	Action:    repoSync,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: tmplRepoList,
		},
	},
}

func repoSync(c *cli.Context) error {
	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	repos, err := client.RepoListOpts(true, true)
	if err != nil || len(repos) == 0 {
		return err
	}

	tmpl, err := template.New("_").Parse(c.String("format") + "\n")
	if err != nil {
		return err
	}

	org := c.String("org")
	for _, repo := range repos {
		if org != "" && org != repo.Owner {
			continue
		}
		tmpl.Execute(os.Stdout, repo)
	}
	return nil
}
