package goblin

// Options ...
type Options struct {
	Format      string
	Compression string
	Backend     string

	Log Logger

	ext string
}

type Logger struct {
	Level string
	JSON  bool
}
