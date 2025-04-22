package templates

const APPLICATION_TPL = `package {{.Name}}

type {{.NameCapitalized}} struct {}

func New{{.NameCapitalized}}() (*{{.NameCapitalized}}, error) {
  panic("implement me")
}

func (m *{{.NameCapitalized}}) Ignite() error {
  panic("implement me")
}

func (m *{{.NameCapitalized}}) Stop() error {
  panic("implement me")
}
`
