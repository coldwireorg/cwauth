package cwauth

func LogoutURL() string {
	return *Address + "oauth2/sessions/logout"
}
