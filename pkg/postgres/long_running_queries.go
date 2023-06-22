package postgres

import (
	"context"
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

func (c *Client) GetListOfLongRunningQueries() (LongRunningQueries, error) {
	longRunningQueries := make(LongRunningQueries, 0)

	err := c.db.NewRaw("SELECT pid, query, query_start, wait_event, wait_event_type, now() - query_start AS query_time FROM pg_stat_activity WHERE (now() - query_start) > interval '5 seconds' and state = 'active'").
		Scan(context.TODO(), &longRunningQueries)

	if err != nil {
		return nil, err
	}

	return longRunningQueries, nil
}
