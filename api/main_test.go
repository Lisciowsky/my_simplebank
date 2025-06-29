package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// func newTestServer(t *testing.T, store db.Store) *Server {
// 	config := util.Config{
// 		TokenSymetricKey:    util.RandomString(32),
// 		AccessTokenDuration: time.Minute,
// 	}

// 	server, err := NewServer(store, config)
// 	require.NoError(t, err)

// 	return server
// }

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
