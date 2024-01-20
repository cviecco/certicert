package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"

	"github.com/Cloud-Foundations/golib/pkg/log"
)

type Server struct {
	//config       *configuration.Configuration
	//htmlWriters  []HtmlWriter
	//htmlTemplate *template.Template
	logger      log.DebugLogger
	cookieMutex sync.Mutex
	//authCookie   map[string]AuthCookie
	//staticConfig *staticconfiguration.StaticConfiguration
	//userInfo     userinfo.UserInfo
	netClient    *http.Client
	accessLogger log.DebugLogger
	tlsConfig    *tls.Config
	serviceMux   *http.ServeMux
	isReady      bool
}

func StartServer(config *AppConfigFile, logger log.DebugLogger) (*Server, error) {
	return nil, fmt.Errorf("not implemented")
}
