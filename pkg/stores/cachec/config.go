package cachec

type Config struct {
	StrongConsistency bool `json:",default=true"`
	DisableCacheRead  bool `json:",default=false"`
}
