package global

type Runnable func(args ...string)

var (
	actions = make(map[string]Runnable, 0)
)

func Register(name string, action Runnable) {
	actions[name] = action
}

func GetAction(name string) (Runnable, bool) {
	if a, ok := actions[name]; ok {
		return a, ok
	}
	return nil, false
}
