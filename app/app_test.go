package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_App(t *testing.T) {
	app, err := New()
	assert.NoError(t, err, "should be no error getting app")
	app.Start()
	app.EventDistributor.AddRiotEventListener("translation-manager", app.TranslationManager.IncomingEvents)
	app.EventDistributor.AddTranslatedEventListener("merge-manager", app.MergeManager.EventStream)
	select {}
}
