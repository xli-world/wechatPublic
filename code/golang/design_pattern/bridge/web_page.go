package bridge

type WebPage interface {
	getContent() string
}

type About struct {
	theme theme
}

func (a *About) getContent() string {
	return "About page in " + a.theme.getColor()
}

type Careers struct {
	theme theme
}

func (c *Careers) getContent() string {
	return "Careers page in" + c.theme.getColor()
}
