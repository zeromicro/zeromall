package http

type Dao struct {
	Demo *DemoHttp
}

func New() *Dao {
	return &Dao{
		Demo: newDemoHttp(),
	}
}

func (m *Dao) Close() {
	return
}
