package graph

import "practice/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PlayerStore map[string]model.Player
	TeamStore   map[string]model.Team
}
