package plugins

import "testing"

func Test_loadPluginsConfig(t *testing.T) {
	p := &plugins{}
	err := p.loadPluginsConfig("../")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", p.info)
}
