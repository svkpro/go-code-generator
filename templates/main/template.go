package template

type MainTpl struct {
	PackageName string
	ServiceName string
}

func New(pn string, sn string) *MainTpl {
	t := &MainTpl{
		PackageName: pn,
		ServiceName: sn,
	}

	return t
}
