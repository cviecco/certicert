package server

import (
	"fmt"
	"net/http"
)

func setupSecurityHeaders(w http.ResponseWriter) error {
	// All common security headers go here
	w.Header().Set("Strict-Transport-Security", "max-age=315360")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1")
	w.Header().Set("Content-Security-Policy", "default-src 'self' ;style-src 'self' maxcdn.bootstrapcdn.com fonts.googleapis.com 'unsafe-inline'; font-src maxcdn.bootstrapcdn.com fonts.gstatic.com fonts.googleapis.com")
	return nil
}

func (s *Server) getRemoteUserName(w http.ResponseWriter, r *http.Request) (string, error) {
	// If you have a verified cert, no need for cookies
	if r.TLS != nil {
		if len(r.TLS.VerifiedChains) > 0 {
			clientName := r.TLS.VerifiedChains[0][0].Subject.CommonName
			return clientName, nil
		}
	}

	return "", fmt.Errorf("Only mTLS supported")
	/*
		setupSecurityHeaders(w)

		remoteCookie, err := r.Cookie(authCookieName)
		if err != nil {
			s.logger.Debugf(1, "Err cookie %s", err)
			s.oauth2DoRedirectoToProviderHandler(w, r)
			return "", err
		}



		   s.cookieMutex.Lock()
		   defer s.cookieMutex.Unlock()
		   authInfo, ok := s.authCookie[remoteCookie.Value]

		   	if !ok {
		   		//s.oauth2DoRedirectoToProviderHandler(w, r)
		   		return "", errors.New("Cookie not found")
		   	}

		   	if authInfo.ExpiresAt.Before(time.Now()) {
		   		//s.oauth2DoRedirectoToProviderHandler(w, r)
		   		return "", errors.New("Expired Cookie")
		   	}

		   return authInfo.Username, nil
	*/
}
