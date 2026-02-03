//go:build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/maozi01/eco-migrate/database/clickhouse"
)
