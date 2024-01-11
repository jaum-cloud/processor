package trigger

import (
	"github_automation/executor"
	"github_automation/github"
)

func ExecuteGistCode(gistID string) ActionFunc {
	return func(input interface{}) (output interface{}, err error) {
		code, err := github.FetchGoCodeFromGist(gistID)
		if err != nil {
			return nil, nil
		}

		result, err := executor.ExecuteGoCode(code)
		if err != nil {
			return nil, nil
		}

		return result, nil
	}
}
