package migrations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(-1)
	}
	exitVal := m.Run()

	//tearDown
	os.Exit(exitVal)
}
func setup() error {
	return nil
}
func Test_ImportData(t *testing.T) {
	err := Migrate()
	assert.NoError(t, err)
	if !isProd() {
		count := 0
		chassis.DB().Table("gopkg_package").Count(&count)
		assert.NotEmpty(t, count)
		assert.Equal(t, 1, count)
	}
}
