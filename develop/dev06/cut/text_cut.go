package cut

type TextCut struct {
	dto *TextCutFlagDTO
}

func NewTextCut(dto *TextCutFlagDTO) *TextCut {
	return &TextCut{dto}
}

func (receiver *TextCut) Cut(inputRows []string) []string {
	return nil
}
