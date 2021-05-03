package domain

type RouteVars map[string]string

func (v RouteVars) Get(key string) (string, bool) {
	value, ok := v[key]
	return value, ok
}
