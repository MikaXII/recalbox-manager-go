package routes

import (
	"github.com/kataras/iris"
)

// Configure registers the necessary routes to the app.
func Configure(app *iris.Application) {
	app.Get("/", GetHomeHandler).Name = "home"

	app.Get("/audio", GetAudioHandler).Name = "audio"
	app.Post("/audio", PostAudioHandler).Name = "audio_form"

	app.Get("/bios", GetBiosHandler).Name = "bios"
	app.Get("/bios/check", GetBiosCheckHandler).Name = "bios_check"
	app.Post("/bios/upload", iris.LimitRequestBodySize(256<<20), PostBiosUploadHandler).Name = "bios_form"
	app.Get("/bios/delete/{file:string}", GetBiosDeleteHandler).Name = "bios_delete"

	app.Get("/controllers", GetControllersHandler).Name = "controllers"
	app.Post("/controllers", PostControllersHandler).Name = "controllers_form"

	app.Get("/systems", GetSystemsHandler).Name = "systems"
	app.Post("/systems", PostSystemsHandler).Name = "systems_form"

	app.Get("/configuration", GetConfigurationHandler).Name = "configuration"
	app.Post("/configuration", PostConfigurationHandler).Name = "configuration_form"

	app.Get("/screenshots", GetScreenshotsHandler).Name = "screenshots"
	app.Get("/screenshots/delete/{file:string}", GetScreenshotsDeleteHandler).Name = "screenshots_delete"
	app.Get("/screenshots/take", GetScreenshotsTakeHandler).Name = "screenshots_take"

	app.Get("/logs", GetLogsHandler).Name = "logs"
	app.Post("/logs", PostLogsHandler).Name = "logs_form"

	app.Get("/recalbox-conf", GetRecalboxConfHandler).Name = "recalbox-conf"
	app.Post("/recalbox-conf", PostRecalboxConfHandler).Name = "recalbox-conf_form"

	app.Get("/help", GetHelpHandler).Name = "help"

	app.Get("/help/recalbox-support", GetLaunchRecalboxSupportHandler).Name = "launch-recalbox-support"

	app.Get("/os/reboot", GetRebootOsHandler).Name = "os-reboot"
	app.Get("/os/shutdown", GetShutdownOsHandler).Name = "os-shutdown"

	app.Get("/es/action/{name:string}", GetActionEsHandler).Name = "es-action"
	app.Get("/es/status", GetStatusEsHandler).Name = "es-status"

	app.Get("/monitoring", GetMonitoringHandler).Name = "monitoring"
}
