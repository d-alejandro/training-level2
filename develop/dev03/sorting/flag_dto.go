package sorting

/*
FlagDTO - DTO for command line flags
*/
type FlagDTO struct {
	ColumnNumberFlag    int
	EndingSpaceFlag     bool
	UniqueFlag          bool
	DescendingOrderFlag bool
	SortByNumberFlag    bool
	MonthNameFlag       bool
	SortCheckFlag       bool
	HumanNumericFlag    bool
}
