package anagramm_service

import (
	"anagramm/pkg/logger"
	"go.uber.org/fx"
	"strings"
)

var Module = fx.Provide(New)

type Param struct {
	fx.In
	Logger logger.ILogger
}

type Service interface {
	Load([]string) error
	Get(string) ([]string, error)
}

type service struct {
	logger logger.ILogger
	dict   []string
}

func New(p Param) Service {
	return &service{
		logger: p.Logger,
	}
}

func (s *service) loadDict(items []string) error {
	for _, v := range items {
		s.dict = append(s.dict, strings.ToLower(v))
	}
	return nil
}

func (s *service) getDict() []string {
	return s.dict
}

func (s *service) Load(dict []string) error {
	s.logger.Debug("load dictionary: %v", dict)
	return s.loadDict(dict)
}

func (s *service) Get(word string) ([]string, error) {
	s.logger.Debug("get all anagrams of '%s'", word)
	res := findAll(s.getDict(), word)
	return res, nil
}

func findAll(dict []string, word string) []string {
	var res = make([]string, 0)

	for _, v := range dict {
		if len(v) == len(word) {
			if isAnagram(strings.ToLower(word), v) {
				res = append(res, v)
			}
		}
	}

	return res
}

func isAnagram(s, target string) bool {
	chars := make(map[rune]int)
	for _, v := range s {
		chars[v]++
	}

	for _, v := range target {
		chars[v]--
	}

	for k, _ := range chars {
		if chars[k] != 0 {
			return false
		}
	}

	return true
}
