package nodejs

type Runtime struct{}

func (r *Runtime) Name() string {
	return "nodejs"
}

func (r *Runtime) Handler() string {
	return "index.handle"
}

func (r *Runtime) Shimmed() bool {
	return false
}
