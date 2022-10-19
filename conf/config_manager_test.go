package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInit(t *testing.T) {

	os.Setenv("APP_LOG_LEVEL", "warn")
	os.Setenv("PZN_ELASTIC_HOST", "https://es-pzn.internal.doubtnut.com/")
	os.Setenv("PZN_ELASTIC_PORT", "80")
	_, fileTypeError := New("configs", "dev.txt")
	assert.NotNil(t, fileTypeError)

	_, doesNotExistError := New("configs", "random.yml")
	assert.NotNil(t, doesNotExistError)

	_, err := New("test_configs", "dev.yml")
	assert.NotNil(t, err)

}
