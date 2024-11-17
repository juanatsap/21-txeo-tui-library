package models

type Panel int

const (
	TabsPanel Panel = iota
	LeftSidebarPanel
	CentralPanel
	RightSidebarPanel
)

func (p Panel) String() string {
	switch p {
	case TabsPanel:
		return "Tabs"
	case LeftSidebarPanel:
		return "Left Sidebar"
	case CentralPanel:
		return "Central"
	case RightSidebarPanel:
		return "Right Sidebar"
	default:
		return "?"
	}
}

func (p Panel) Int() int {
	return int(p)
}

func (p *Panel) Next() {
	*p = (*p + 1) % 4

	if *p == 0 {
		*p = 4
	}
}

func (p *Panel) Prev() {
	*p = (*p - 1) % 4
}

func (p *Panel) Reset() {
	*p = TabsPanel
}
