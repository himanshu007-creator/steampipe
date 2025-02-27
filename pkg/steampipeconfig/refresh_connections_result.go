package steampipeconfig

import (
	"fmt"
	"github.com/turbot/steampipe/pkg/utils"
	"strings"
)

// RefreshConnectionResult is a structure used to contain the result of either a RefreshConnections or a NewLocalClient operation
type RefreshConnectionResult struct {
	UpdatedConnections bool
	Warnings           []string
	Error              error
}

func (r *RefreshConnectionResult) AddWarning(warning string) {
	r.Warnings = append(r.Warnings, warning)
}

func (r *RefreshConnectionResult) ShowWarnings() {
	for _, w := range r.Warnings {
		fmt.Println(w)
	}
}

func (r *RefreshConnectionResult) Merge(other *RefreshConnectionResult) {
	if other == nil {
		return
	}
	if other.UpdatedConnections {
		r.UpdatedConnections = other.UpdatedConnections
	}
	if other.Error != nil {
		r.Error = other.Error
	}
	r.Warnings = append(r.Warnings, other.Warnings...)
}

func (r *RefreshConnectionResult) String() any {
	var op strings.Builder
	if len(r.Warnings) > 0 {
		op.WriteString(fmt.Sprintf("%s:\n\t%s", utils.Pluralize("Warning", len(r.Warnings)), strings.Join(r.Warnings, "\n\t")))
	}
	if r.Error != nil {
		op.WriteString(fmt.Sprintf("Error: %s\n", r.Error.Error()))
	}
	op.WriteString(fmt.Sprintf("UpdatedConnections: %v\n", r.UpdatedConnections))
	return op.String()
}
