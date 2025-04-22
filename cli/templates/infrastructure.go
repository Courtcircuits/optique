package templates

const INFRASTRUCTURE_TPL = `package {{.Name}}

type {{.NameCapitalized}} struct {}

func New{{.NameCapitalized}}() (*{{.NameCapitalized}}, error) {
  panic("implement me")
}

func (m *{{.NameCapitalized}}) Setup() error {
  panic("implement me")
}

func (m *{{.NameCapitalized}}) Shutdown() error {
  panic("implement me")
}
`
