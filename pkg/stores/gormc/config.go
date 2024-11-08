package gormc

type Config struct {
	Separation   bool     `json:",default=false"`
	Master       string   `json:",optional"`
	Sources      []string `json:",optional"`
	Replicas     []string `json:",optional"`
	DNS          string   `json:",optional"`
	Debug        bool     `json:",default=false"`
	MaxIdleConns int      `json:",default=10"`
	MaxOpenConns int      `json:",default=100"`
}
