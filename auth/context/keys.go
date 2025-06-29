package context

type contextKey string

const authUserKey contextKey = "authUser"

func AuthUserKey() contextKey {
	return authUserKey
}
