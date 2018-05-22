package layouts

import (
	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func New(app *iris.Application) iris.Handler {
	return func(ctx iris.Context) {
		currentPath := ctx.Path()
		menuEntries := []structs.MenuItem{
			{
				Link:     app.GetRoute("home").FormattedPath,
				Glyph:    "home",
				Label:    ctx.Translate("Accueil"),
				IsActive: app.GetRoute("home").FormattedPath == currentPath,
			},
			{
				Link:     "/monitoring",
				Glyph:    "signal",
				Label:    ctx.Translate("Monitoring"),
				IsActive: false,
			},
			{
				Link:     app.GetRoute("audio").FormattedPath,
				Glyph:    "volume-up",
				Label:    ctx.Translate("Audio"),
				IsActive: app.GetRoute("audio").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("bios").FormattedPath,
				Glyph:    "compact-disc",
				Label:    ctx.Translate("BIOS"),
				IsActive: app.GetRoute("bios").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("controllers").FormattedPath,
				Glyph:    "gamepad",
				Label:    ctx.Translate("Contrôleurs"),
				IsActive: app.GetRoute("controllers").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("systems").FormattedPath,
				Glyph:    "hdd",
				Label:    ctx.Translate("Systèmes"),
				IsActive: app.GetRoute("systems").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("configuration").FormattedPath,
				Glyph:    "cog",
				Label:    ctx.Translate("Configuration"),
				IsActive: app.GetRoute("configuration").FormattedPath == currentPath,
			},
			{
				Link:     "/roms",
				Glyph:    "save",
				Label:    ctx.Translate("ROMs"),
				IsActive: false,
			},
			{
				Link:     "/screenshots",
				Glyph:    "images",
				Label:    ctx.Translate("Screenshots"),
				IsActive: false,
			},
			{
				Link:  "/help",
				Glyph: "question-circle",
				Label: ctx.Translate("Dépannage"), Children: []structs.MenuItem{
					{
						Link:     "/logs",
						Glyph:    "file",
						Label:    ctx.Translate("Logs"),
						IsActive: false,
					},
					{
						Link:     "/recalbox-conf",
						Glyph:    "file",
						Label:    "recalbox.conf",
						IsActive: false,
					},
					{
						Link:     "/help",
						Glyph:    "question-circle",
						Label:    ctx.Translate("Dépannage"),
						IsActive: false,
					},
				}},
		}

		for k, v := range menuEntries {
			if v.IsActive {
				menuEntries[k].ActiveClass = "active"
			}
		}

		menuLanguages := make(map[string]string)
		languages := []structs.Language{}
		err := viper.UnmarshalKey("availableLanguages", &languages)

		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		for _, v := range languages {
			menuLanguages[v.Locale] = v.Name
		}

		currentLang := menuLanguages[ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())]

		ctx.ViewLayout("layouts/default.pug")
		ctx.ViewData("RecalboxManagerTitle", ctx.Translate("Recalbox Manager"))
		ctx.ViewData("MenuEntries", menuEntries)
		ctx.ViewData("CurrentLang", currentLang)
		ctx.ViewData("AvailableLang", menuLanguages)
		ctx.ViewData("AppVersion", viper.GetString("app.version"))

		ctx.Next()
	}
}

// Configure creates a new layout middleware and registers that to the app.
func Configure(app *iris.Application) {
	h := New(app)
	app.Use(h)
}
