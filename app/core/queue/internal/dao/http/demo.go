package http

type DemoHttp struct {
}

func newDemoHttp() *DemoHttp {
	return &DemoHttp{}
}

func (m *DemoHttp) Hello() (err error) {
	return nil
}
