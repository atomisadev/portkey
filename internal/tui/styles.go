package tui

import "github.com/charmbracelet/lipgloss"

const (
	ColorRosewater = "#f5e0dc"
	ColorFlamingo  = "#f2cdcd"
	ColorPink      = "#f5c2e7"
	ColorMauve     = "#cba6f7"
	ColorRed       = "#f38ba8"
	ColorMaroon    = "#eba0ac"
	ColorPeach     = "#fab387"
	ColorYellow    = "#f9e2af"
	ColorGreen     = "#a6e3a1"
	ColorTeal      = "#94e2d5"
	ColorSky       = "#89dceb"
	ColorSapphire  = "#74c7ec"
	ColorBlue      = "#89b4fa"
	ColorLavender  = "#b4befe"
	ColorText      = "#cdd6f4"
	ColorSubtext1  = "#bac2de"
	ColorOverlay0  = "#6c7086"
	ColorSurface0  = "#313244"
	ColorBase      = "#1e1e2e"
	ColorMantle    = "#181825"
	ColorCrust     = "#11111b"
)

var (
	AppStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorMauve)).
			Foreground(lipgloss.Color(ColorText))

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorLavender)).
			Bold(true).
			PaddingBottom(1)

	SelectedTitle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color(ColorMauve)).
			Foreground(lipgloss.Color(ColorMauve)).
			Padding(0, 0, 0, 1)

	SelectedDesc = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPink)).
			Padding(0, 0, 0, 2)
)
