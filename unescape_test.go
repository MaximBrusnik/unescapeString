package unescapeString

import "testing"

func TestUnescapeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
		errMsg   string
	}{
		// Основные случаи
		{
			name:     "empty string",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "simple string",
			input:    "abc",
			expected: "abc",
			wantErr:  false,
		},
		{
			name:     "string and digits",
			input:    "a2b3",
			expected: "aabbb",
			wantErr:  false,
		},
		{
			name:     "escape characters",
			input:    `a\2b\3`,
			expected: "a2b3",
			wantErr:  false,
		},
		{
			name:     "mixed escape and expansion",
			input:    `a\12b3`,
			expected: "a11bbb",
			wantErr:  false,
		},
		{
			name:     "multiple digits count",
			input:    "a10",
			expected: "aaaaaaaaaa",
			wantErr:  false,
		},

		// Крайние случаи
		{
			name:     "escape at end",
			input:    `abc\`,
			expected: "",
			wantErr:  true,
			errMsg:   "escape character at end of string",
		},
		{
			name:     "digit at start",
			input:    "1abc",
			expected: "",
			wantErr:  true,
			errMsg:   "digit without preceding character",
		},
		{
			name:     "only digits",
			input:    "123",
			expected: "",
			wantErr:  true,
			errMsg:   "digit without preceding character",
		},
		{
			name:     "zero repeat count",
			input:    "a0",
			expected: "",
			wantErr:  true,
			errMsg:   "invalid repeat count",
		},
		{
			name:     "negative repeat count",
			input:    "a-1",
			expected: "a-",
			wantErr:  false, // "-1" трактуется как символы '-', '1'
		},
		{
			name:     "unicode characters",
			input:    "я3",
			expected: "яяя",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnescapeString(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnescapeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err.Error() != tt.errMsg {
					t.Errorf("UnescapeString() error message = %v, want %v", err.Error(), tt.errMsg)
				}
				return
			}

			if got != tt.expected {
				t.Errorf("UnescapeString() = %v, want %v", got, tt.expected)
			}
		})
	}
}
