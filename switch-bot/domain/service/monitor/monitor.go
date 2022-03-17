package monitor

type Monitor interface {
	CurrentDialogContains(text []string) bool
}
