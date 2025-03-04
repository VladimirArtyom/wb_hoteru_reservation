package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)


func NoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)
	
	cookie := http.Cookie{
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Secure: false, // Still in dev mode, changer cela plus tard
		Path: "/",
	}

	csrfHandler.SetBaseCookie(cookie)
	return csrfHandler
}
