import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

    {{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/cloudernative/go-zero/core/stores/builder"
	"github.com/cloudernative/go-zero/core/stores/sqlc"
	"github.com/cloudernative/go-zero/core/stores/sqlx"
	"github.com/cloudernative/go-zero/core/stringx"
)
