//go:build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/eco-migrate/migrate/v4/database/clickhouse"
)
