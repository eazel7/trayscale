// Code generated by gtkbuilder. DO NOT EDIT.

package main

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type MainWindow struct {
	*adw.ApplicationWindow

	ToastOverlay   *adw.ToastOverlay
	StatusSwitch   *gtk.Switch
	MainMenuButton *gtk.MenuButton
	PeersStack     *gtk.Stack
}

func NewMainWindow(app *gtk.Application) *MainWindow {
	ToastOverlay := adw.NewToastOverlay()
	ToastOverlay0 := adw.NewLeaflet()
	ToastOverlay00 := gtk.NewBox(0, 0)
	ToastOverlay000 := adw.NewHeaderBar()
	StatusSwitch := gtk.NewSwitch()
	MainMenuButton := gtk.NewMenuButton()
	ToastOverlay001 := gtk.NewStackSidebar()
	ToastOverlay01 := gtk.NewBox(0, 0)
	ToastOverlay010 := adw.NewHeaderBar()
	ToastOverlay0100 := gtk.NewBox(0, 0)
	PeersStack := gtk.NewStack()
	parent := adw.NewApplicationWindow(app)

	ToastOverlay.SetChild(ToastOverlay0)

	ToastOverlay0.Append(ToastOverlay00)
	ToastOverlay0.Append(ToastOverlay01)

	ToastOverlay00.SetObjectProperty("orientation", 1)
	ToastOverlay00.SetObjectProperty("width-request", 360)
	ToastOverlay00.Append(ToastOverlay000)
	ToastOverlay00.Append(ToastOverlay001)

	ToastOverlay000.SetObjectProperty("show-end-title-buttons", false)
	ToastOverlay000.PackStart(StatusSwitch)
	ToastOverlay000.PackEnd(MainMenuButton)

	MainMenuButton.SetObjectProperty("icon-name", "open-menu-symbolic")
	MainMenuButton.SetObjectProperty("primary", true)

	ToastOverlay001.SetObjectProperty("stack", PeersStack)
	ToastOverlay001.SetObjectProperty("vexpand", true)
	ToastOverlay001.SetObjectProperty("width-request", 270)

	ToastOverlay01.SetObjectProperty("hexpand", true)
	ToastOverlay01.SetObjectProperty("orientation", 1)
	ToastOverlay01.Append(ToastOverlay010)
	ToastOverlay01.Append(PeersStack)

	ToastOverlay010.SetObjectProperty("show-start-title-buttons", false)
	ToastOverlay010.SetTitleWidget(ToastOverlay0100)

	PeersStack.SetObjectProperty("transition-type", 7)
	PeersStack.SetObjectProperty("vexpand", true)

	parent.SetObjectProperty("content", ToastOverlay)
	parent.SetObjectProperty("default-height", 600)
	parent.SetObjectProperty("default-width", 800)
	parent.SetObjectProperty("title", "Trayscale")
	parent.SetContent(ToastOverlay)

	return &MainWindow{
		ApplicationWindow: parent,

		ToastOverlay:   ToastOverlay,
		StatusSwitch:   StatusSwitch,
		MainMenuButton: MainMenuButton,
		PeersStack:     PeersStack,
	}
}

type PeerPage struct {
	*adw.StatusPage

	IPGroup                 *adw.PreferencesGroup
	OptionsGroup            *adw.PreferencesGroup
	AdvertiseExitNodeRow    *adw.ActionRow
	AdvertiseExitNodeSwitch *gtk.Switch
	MiscGroup               *adw.PreferencesGroup
	ExitNodeRow             *adw.ActionRow
	ExitNodeSwitch          *gtk.Switch
	OnlineRow               *adw.ActionRow
	Online                  *gtk.Image
	LastSeenRow             *adw.ActionRow
	LastSeen                *gtk.Label
	CreatedRow              *adw.ActionRow
	Created                 *gtk.Label
	LastWriteRow            *adw.ActionRow
	LastWrite               *gtk.Label
	LastHandshakeRow        *adw.ActionRow
	LastHandshake           *gtk.Label
	RxBytesRow              *adw.ActionRow
	RxBytes                 *gtk.Label
	TxBytesRow              *adw.ActionRow
	TxBytes                 *gtk.Label
}

func NewPeerPage() *PeerPage {
	parent0 := adw.NewClamp()
	parent00 := gtk.NewBox(0, 0)
	IPGroup := adw.NewPreferencesGroup()
	OptionsGroup := adw.NewPreferencesGroup()
	AdvertiseExitNodeRow := adw.NewActionRow()
	AdvertiseExitNodeSwitch := gtk.NewSwitch()
	MiscGroup := adw.NewPreferencesGroup()
	ExitNodeRow := adw.NewActionRow()
	ExitNodeSwitch := gtk.NewSwitch()
	OnlineRow := adw.NewActionRow()
	Online := gtk.NewImage()
	LastSeenRow := adw.NewActionRow()
	LastSeen := gtk.NewLabel("")
	CreatedRow := adw.NewActionRow()
	Created := gtk.NewLabel("")
	LastWriteRow := adw.NewActionRow()
	LastWrite := gtk.NewLabel("")
	LastHandshakeRow := adw.NewActionRow()
	LastHandshake := gtk.NewLabel("")
	RxBytesRow := adw.NewActionRow()
	RxBytes := gtk.NewLabel("")
	TxBytesRow := adw.NewActionRow()
	TxBytes := gtk.NewLabel("")
	parent := adw.NewStatusPage()

	parent0.SetChild(parent00)

	parent00.SetObjectProperty("orientation", 1)
	parent00.SetObjectProperty("spacing", 12)
	parent00.Append(IPGroup)
	parent00.Append(OptionsGroup)
	parent00.Append(MiscGroup)

	IPGroup.SetObjectProperty("title", "Tailscale IPs")

	OptionsGroup.SetObjectProperty("title", "Options")
	OptionsGroup.Add(AdvertiseExitNodeRow)

	AdvertiseExitNodeRow.SetObjectProperty("title", "Advertise exit node")
	AdvertiseExitNodeRow.AddSuffix(AdvertiseExitNodeSwitch)

	AdvertiseExitNodeSwitch.SetObjectProperty("margin-bottom", 12)
	AdvertiseExitNodeSwitch.SetObjectProperty("margin-top", 12)

	MiscGroup.SetObjectProperty("title", "Misc.")
	MiscGroup.Add(ExitNodeRow)
	MiscGroup.Add(OnlineRow)
	MiscGroup.Add(LastSeenRow)
	MiscGroup.Add(CreatedRow)
	MiscGroup.Add(LastWriteRow)
	MiscGroup.Add(LastHandshakeRow)
	MiscGroup.Add(RxBytesRow)
	MiscGroup.Add(TxBytesRow)

	ExitNodeRow.SetObjectProperty("icon-name", "security-high-symbolic")
	ExitNodeRow.SetObjectProperty("title", "Use as exit node")
	ExitNodeRow.AddSuffix(ExitNodeSwitch)

	ExitNodeSwitch.SetObjectProperty("margin-bottom", 12)
	ExitNodeSwitch.SetObjectProperty("margin-top", 12)

	OnlineRow.SetObjectProperty("title", "Online")
	OnlineRow.AddSuffix(Online)

	LastSeenRow.SetObjectProperty("title", "Last seen")
	LastSeenRow.AddSuffix(LastSeen)

	CreatedRow.SetObjectProperty("title", "Created at")
	CreatedRow.AddSuffix(Created)

	LastWriteRow.SetObjectProperty("title", "Last write")
	LastWriteRow.AddSuffix(LastWrite)

	LastHandshakeRow.SetObjectProperty("title", "Last handshake")
	LastHandshakeRow.AddSuffix(LastHandshake)

	RxBytesRow.SetObjectProperty("title", "Bytes received")
	RxBytesRow.AddSuffix(RxBytes)

	TxBytesRow.SetObjectProperty("title", "Bytes sent")
	TxBytesRow.AddSuffix(TxBytes)

	parent.SetChild(parent0)

	return &PeerPage{
		StatusPage: parent,

		IPGroup:                 IPGroup,
		OptionsGroup:            OptionsGroup,
		AdvertiseExitNodeRow:    AdvertiseExitNodeRow,
		AdvertiseExitNodeSwitch: AdvertiseExitNodeSwitch,
		MiscGroup:               MiscGroup,
		ExitNodeRow:             ExitNodeRow,
		ExitNodeSwitch:          ExitNodeSwitch,
		OnlineRow:               OnlineRow,
		Online:                  Online,
		LastSeenRow:             LastSeenRow,
		LastSeen:                LastSeen,
		CreatedRow:              CreatedRow,
		Created:                 Created,
		LastWriteRow:            LastWriteRow,
		LastWrite:               LastWrite,
		LastHandshakeRow:        LastHandshakeRow,
		LastHandshake:           LastHandshake,
		RxBytesRow:              RxBytesRow,
		RxBytes:                 RxBytes,
		TxBytesRow:              TxBytesRow,
		TxBytes:                 TxBytes,
	}
}
