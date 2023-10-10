package api

import (
	"os"
	"testing"
	"time"

	db "github.com/Arodrigow/simple_bank_project/db/sqlc"
	"github.com/Arodrigow/simple_bank_project/util"
	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		AcessTokenDuration: time.Minute,
	}

	server := NewServer(config, store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
