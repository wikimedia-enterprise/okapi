package task

// Params task execution params
type Params struct {
	DBName  string
	Restart bool
	Workers int
	Pointer int
	Offset  int
	Limit   int
}
