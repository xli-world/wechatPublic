package bridge

type theme interface {
	getColor() string
}

type DarkTheme struct {
}

func (d *DarkTheme) getColor() string {
	return "Dark Black"
}

type LightTheme struct {
}

func (l *LightTheme) getColor() string {
	return "Off white"
}

type AquaTheme struct {
}

func (a *AquaTheme) getColor() string {
	return "Light Blue"
}
