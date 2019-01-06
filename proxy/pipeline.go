package proxy

type Pipeline struct {
	Log chan string
}

func InitPipeline() *Pipeline {
	return &Pipeline{
		Log: make(chan string),
	}
}
