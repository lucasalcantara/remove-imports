package analyzer

import "strings"

type ScopeComments struct {
	inCommentBlock bool
}

func (s *ScopeComments) InCommentLine(text string) bool {
	s.inCommentBlock = strings.Contains(text, "/*") || s.inCommentBlock
	isCommentLine := strings.Contains(text, "//")

	if s.inCommentBlock {
		defer s.updateInCommentBlock(text)
	}

	return s.inCommentBlock || isCommentLine
}

func (s *ScopeComments) updateInCommentBlock(text string) {
	s.inCommentBlock = !strings.Contains(text, "*/")
}
