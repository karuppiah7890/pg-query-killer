package postgres

import (
	"context"
	"fmt"
	"time"
)

type LongRunningQuery struct {
	ProcessId      int       `bun:"pid"`
	Query          string    `bun:"query"`
	QueryStartTime time.Time `bun:"query_start"`
	QueryTime      time.Time `bun:"query_time"`
	WaitEvent      string    `bun:"wait_event"`
	WaitEventType  string    `bun:"wait_event_type"`
}

type LongRunningQueries []LongRunningQuery

func (c *Client) GetListOfLongRunningQueries(duration time.Duration) (LongRunningQueries, error) {
	longRunningQueries := make(LongRunningQueries, 0)

	err := c.db.NewRaw(longRunningQuerySql(duration)).
		Scan(context.TODO(), &longRunningQueries)

	if err != nil {
		return nil, err
	}

	return longRunningQueries, nil
}

func longRunningQuerySql(duration time.Duration) string {
	return fmt.Sprintf("SELECT pid, query, query_start, wait_event, wait_event_type, now() - query_start AS query_time FROM pg_stat_activity WHERE (now() - query_start) > interval '%d seconds' and state = 'active'", int(duration.Seconds()))
}
