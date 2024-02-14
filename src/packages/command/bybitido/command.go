package bybitido

import (
	"time"

	"github.com/hungdoo/bot/src/common"
	"github.com/hungdoo/bot/src/packages/bybitapi"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/utils"
)

type IdoCommand struct {
	command.Command
	Id string `bson:"_id,unique"`
}

func (c *IdoCommand) Validate(data []string) error {
	return nil
}

func (c *IdoCommand) SetData(newValue []string) (err error) {
	return nil
}

func (c *IdoCommand) Execute(mustReport bool, subcommand string) (string, *common.ErrorWithSeverity) {
	switch subcommand {
	case "latest":
		latest, err := bybitapi.GetLatestProject(bybitapi.IDO_URL)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		return utils.PrettyPrint(latest), nil
	case "all":
		a, err := bybitapi.FetchProjects(bybitapi.IDO_URL)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		return utils.PrettyPrint(a), nil
	default:
		p, err := bybitapi.GetUpcomingProjects(bybitapi.IDO_URL, time.Now())
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		// Process the results
		if len(p) != 0 {
			return utils.PrettyPrint(p), nil
		}
		return "", nil
	}
}
