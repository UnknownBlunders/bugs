package bugs

import "strings"

const (
	VerbHelp   = "help"
	VerbList   = "list"
	VerbCreate = "create"
	VerbUpdate = "update"
)

type Action struct {
	Verb      string
	BugTitle  string
	BugID     string
	BugStatus string
}

func ParseArgs(args []string) Action {
	if len(args) < 2 {
		return Action{
			Verb: VerbHelp,
		}
	}

	switch args[1] {
	case "list":
		return Action{
			Verb: VerbList,
		}
	case "create":
		if len(args) < 3 {
			return Action{
				Verb: VerbHelp,
			}
		}
		return Action{
			Verb:     VerbCreate,
			BugTitle: strings.Join(args[2:], " "),
		}
	case "update":
		if len(args) != 4 {
			return Action{
				Verb: VerbHelp,
			}
		}
		return Action{
			Verb:      VerbUpdate,
			BugID:     args[2],
			BugStatus: args[3],
		}
	default:
		return Action{
			Verb: VerbHelp,
		}
	}
}
