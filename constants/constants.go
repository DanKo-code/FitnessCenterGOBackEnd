package constants

type userRole struct {
	Client string
	Admin  string
	Coach  string
}

var ROLES = userRole{
	Client: "client",
	Admin:  "admin",
	Coach:  "coach",
}

type refreshToken struct {
	HttpOnly bool
	MaxAge   int
}

type cookieSettings struct {
	RefreshToken refreshToken
}

var COOKIE_SETTINGS = cookieSettings{
	RefreshToken: refreshToken{
		HttpOnly: true,
		MaxAge:   6048e5,
	},
}

var ACCESS_TOKEN_EXPIRATION = 18e5
