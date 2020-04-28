package engine

// ActionHandler type
type ActionHandler struct{}

// HandleAction handles actions defined in rule.
func (s ActionHandler) HandleAction(rule Rule, rec Record) {
	for _, a := range rule.Actions {
		switch a {
		case Hash:
			rec.Ctx.SetHashes(HashSet{})
		case Alert:
		case Tag:
		default:
			rec.Ctx.AddRule(rule)
		}
	}
}
