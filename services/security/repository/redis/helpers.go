package redis

import "github.com/mreza0100/golog"

type helpers struct {
	lgr *golog.Core
}

func (h *helpers) addPrefixS(token string) string {
	return SESSION_PREFIX + token
}

func (h *helpers) addPrefixes(tokens ...string) []string {
	for idx, t := range tokens {
		tokens[idx] = SESSION_PREFIX + t
	}
	return tokens
}
