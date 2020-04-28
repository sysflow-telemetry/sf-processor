package engine

// ActionHandler type
type ActionHandler struct {
	conf map[string]string
}

// NewActionHandler creates a new handler.
func NewActionHandler(conf map[string]string) ActionHandler {
	return ActionHandler{conf}
}

// HandleAction handles actions defined in rule.
func (s ActionHandler) HandleAction(rule Rule, r *Record) {
	for _, a := range rule.Actions {
		switch a {
		case Hash:
			r.Ctx.SetHashes(HashSet{})
		case Alert:
		case Tag:
		default:
			r.Ctx.AddRule(rule)
		}
	}
}
