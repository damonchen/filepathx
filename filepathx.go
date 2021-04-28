package filepathx

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// Glob support double star and single star
func Glob(pattern string) ([]string, error) {
	if !strings.Contains(pattern, "**") {
		return filepath.Glob(pattern)
	}

	patterns := strings.Split(pattern, "**")
	newPatterns := make([]string, len(patterns))
	for i, p := range patterns {
		p = strings.TrimRight(p, "/")
		newPatterns[i] = p
	}

	return expand(newPatterns)
}

func expand(patterns []string) ([]string, error) {
	// all matches will put in array
	matches := []string{""}
	for _, pattern := range patterns {
		m, err := expandMatches(matches, pattern)
		if err != nil {
			return nil, err
		}
		matches = m
	}

	return matches, nil
}

func expandMatches(matches []string, pattern string) ([]string, error) {
	var r []string
	for _, match := range matches {
		paths, err := filepath.Glob(match + pattern)
		if err != nil {
			return nil, err
		}

		for _, path := range paths {
			// TODO: cache the result
			err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				r = append(r, path)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}
