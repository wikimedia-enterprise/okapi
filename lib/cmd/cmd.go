package cmd

import (
	"flag"
)

// Context current execution context
var Context *Params = &Params{}

// Params app params
type Params struct {
	Server  *string
	DBName  *string
	Port    *string
	Task    *string
	Pointer *int
	Offset  *int
	Limit   *int
	Restart *bool
	Workers *int
}

// Parse get command line arguments
func (params *Params) Parse() {
	params.Server = flag.String("server", string(API), "What server to start ('api', 'events', 'stream', 'queue)")
	params.Restart = flag.Bool("restart", false, "Clean the task cache and start the task again")
	params.Workers = flag.Int("workers", 5, "Number of workers for task")
	params.DBName = flag.String("db_name", "*", "Specific project to run")
	params.Port = flag.String("port", "4000", "API server prot")
	params.Task = flag.String("task", "", "Task to be executed")
	params.Offset = flag.Int("offset", 0, "Offset for task to start")
	params.Limit = flag.Int("limit", 10000, "Maximum amount of items to process")
	params.Pointer = flag.Int("pointer", 0, "Unique identifier of last processed entity")
	flag.Parse()
}

// Init function to initialize on startup
func Init() error {
	Context.Parse()
	return nil
}
