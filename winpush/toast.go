// Package winpush provides Windows Toast Notifications for Go applications.
//
// Features:
//   - Customizable toast notifications
//   - Action buttons support
//   - Adaptive templates
//   - Silent execution (no PowerShell window)
//
// Basic usage:
//
//	notifier := winpush.Notificator{
//	    Title:   "Hello",
//	    Message: "World",
//	}
//	err := notifier.Push()
package winpush

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.New("winpush")
	tmpl.Parse(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = '{{if .AppID}}{{.AppID}}{{else}}Windows App{{end}}'

$template = @"
<toast activationType="{{.ActivationType}}" launch="{{.ActivationArguments}}" duration="{{.Duration}}">
    <visual>
        <binding template="ToastGeneric">
            {{if .Icon}}<image placement="appLogoOverride" src="{{.Icon}}" />{{end}}
            {{if .Title}}<text hint-style="title">{{.Title}}</text>{{end}}
			{{if .Subtitle}}<text hint-style="subtitle">{{.Subtitle}}</text>{{end}}
            {{if .Message}}<text hint-style="body">{{.Message}}</text>{{end}}
        </binding>
    </visual>
	{{if .Actions}}
	<actions>
		{{range .Actions}}
		<action
			activationType="{{.ActivationType}}"
			content="{{.Content}}"
			arguments="{{.Arguments}}"
			{{if .Icon}}imageUri="{{.Icon}}"{{end}}
			{{if .Placement}}placement="{{.Placement}}"{{end}}
		/>
		{{end}}
	</actions>
	{{end}}
</toast>
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
    `)
}

func (n *Notificator) applySetting() {
	for i := range n.Actions {
		if n.Actions[i].ActivationType == "" {
			n.Actions[i].ActivationType = "protocol"
		}
	}

	if n.ActivationType == "" {
		n.ActivationType = "protocol"
	}
	if n.Duration == "" {
		n.Duration = "short"
	}
	if n.Duration != "short" && n.Duration != "long" {
		n.Duration = "short"
	}
}

func (n *Notificator) bXML() (string, error) {
	n.applySetting()

	var b bytes.Buffer
	if err := tmpl.Execute(&b, n); err != nil {
		return "", ErrReadXML
	}
	return b.String(), nil
}

func (n *Notificator) executeCMD(xml string) error {
	path := filepath.Join(os.TempDir(), "GoX7specialFileToToast.ps1")
	defer os.Remove(path)
	bsp := []byte{0xEF, 0xBB, 0xBF}
	data := append(bsp, []byte(xml)...)
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		return ErrCreateFile
	}

	cmd := exec.Command("PowerShell", "-ExecutionPolicy", "Bypass", "-File", path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		return ErrExecuteToast
	}
	return nil
}

func (n *Notificator) Push() error {
	xml, err := n.bXML()
	if err != nil {
		return err
	}
	return n.executeCMD(xml)
}
