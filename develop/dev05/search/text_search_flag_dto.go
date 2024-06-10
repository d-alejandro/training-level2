package search

/*
TextSearchFlagDTO - DTO for command line flags
*/
type TextSearchFlagDTO struct {
	RowsAfterFlag   int
	RowsBeforeFlag  int
	RowsContextFlag int
	CountFlag       bool
	IgnoreCaseFlag  bool
	InvertFlag      bool
	FixedFlag       bool
	LineNumFlag     bool
}
