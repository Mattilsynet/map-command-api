package subject

import "github.com/Mattilsynet/mapis/gen/go/command/v1"

const (
	APPLY  Operation = "apply"
	DELETE Operation = "delete"
)

type Operation string

type CommandSubject struct {
	Kind    string
	Id      string
	Session string
	prefix  string
}

func NewCommandSubject(command *command.Command) *CommandSubject {
	kind := command.GetSpec().GetType().GetKind()
	session := command.GetSpec().GetSessionId()
	id := command.GetStatus().GetId()
	prefix := "map"
	return &CommandSubject{
		Kind:    kind,
		Session: session,
		Id:      id,
		prefix:  prefix,
	}
}

func (qs *CommandSubject) ToCommand(command *command.Command) string {
	operation := command.GetSpec().GetOperation()
	subject := qs.prefix + "." + qs.Kind + "." + operation
	return subject
}
