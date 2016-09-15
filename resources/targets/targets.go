package targets

import (
	"github.com/pressly/chi/render"
	"errors"
	"net/http"
	"time"
	"github.com/montanaflynn/stats"
)



type Target struct {
	Username   *string `json:",omitempty"`
	TargetName *string `json:", omitempty"`
	Attempts   *int64 `json:",omitempty"`
	Success    *int64 `json:",omitempty"`
	Percentage *float64 `json:",omitempty"`
	Date       *time.Time `json:",omitempty"`
}

type Targets struct {
	Targets []Target `json:",omitempty"`
	Statistics `json:",omitempty"`
}

type Statistics struct {
	StandardDeviation float64 `json:",omitempty"`
	Mean              float64 `json:",omitempty"`
	Median            float64 `json:",omitempty"`
}

func (t *Targets)generateTargetsAnalysis() (err error) {
	var percentages []float64
	for _, target := range t.Targets {
		percentages = append(percentages, *target.Percentage)
	}

	if len(percentages) == 0 {
		return nil
	}

	if t.StandardDeviation, err = stats.StandardDeviationPopulation(percentages); err != nil {
		return err
	}

	if t.Mean, err = stats.Mean(percentages); err != nil {
		return err
	}

	if t.Median, err = stats.Median(percentages); err != nil {
		return err
	}

	return nil
}

func bindTargetFromRequest(r *http.Request) (*Target, error) {
	var t *Target

	if err := render.Bind(r.Body, &t); err != nil {
		return nil, err
	}

	if t.Username == nil {
		return nil, errors.New("Username is required")
	}

	if t.Attempts == nil {
		return nil, errors.New("Attempts is required")
	}

	if t.TargetName == nil {
		return nil, errors.New("TargetName is required")
	}

	if t.Success == nil {
		return nil, errors.New("Successes is required")
	}

	return t, nil
}



