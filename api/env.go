package api

import (
	"basic/pkg/dynamo"

	"go.uber.org/zap"
)

// Env structure to inject dependencies of global scopes
type Env struct {
	Db  *dynamo.DB
	Log *zap.SugaredLogger
}
