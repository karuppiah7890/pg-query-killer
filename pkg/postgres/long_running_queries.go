package postgres

import (
	"context"
	"fmt"
	"log"
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

var startOfTime time.Time

func init() {
	s, err := time.Parse("03:04", "00:00")
	if err != nil {
		log.Fatalf("error occurred while doing a time related processing: %v", err)
	}
	startOfTime = s
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

func (longRunningQuery LongRunningQuery) String() string {
	return fmt.Sprintf("Query: %+v\n", longRunningQuery.Query) +
		fmt.Sprintf("Time taken by the Query: %+v\n", timeTakenByQuerySince(longRunningQuery, startOfTime)) +
		fmt.Sprintf("Query's start time: %+v\n", longRunningQuery.QueryStartTime) +
		fmt.Sprintf("Process ID of Query: %+v", longRunningQuery.ProcessId)
}

func timeTakenByQuerySince(longRunningQuery LongRunningQuery, from time.Time) time.Duration {
	return longRunningQuery.QueryTime.Sub(from)
}
