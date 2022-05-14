package regex

type TemplateArgs struct {
	Configure     bool
	Type          string
	UseQuery      bool
	UseCommand    bool
	Name          string
	Topic         string
	HasAuth       bool
	HasValidation bool
}

type TemplateConfigure struct {
	Namespace       string
	ApplicationPath string
	QueryPrefix     string
	CommandPrefix   string
}
