package telegramclient_test

import (
	"testing"

	telegramclient "github.com/br0-space/bot-telegramclient"
)

type escapeMarkdownTest struct {
	in  string
	out string
}

var escapeMarkdownTests = []escapeMarkdownTest{
	{"_", "\\_"},
	{"*", "\\*"},
	{"[", "\\["},
	{"]", "\\]"},
	{"(", "\\("},
	{")", "\\)"},
	{"~", "\\~"},
	{"`", "\\`"},
	{">", "\\>"},
	{"#", "\\#"},
	{"+", "\\+"},
	{"-", "\\-"},
	{"=", "\\="},
	{"|", "\\|"},
	{"{", "\\{"},
	{"}", "\\}"},
	{".", "\\."},
	{"!", "\\!"},
	{"\\", "\\\\"},
}

func TestEscapeMarkdown(t *testing.T) {
	t.Parallel()

	for _, test := range escapeMarkdownTests {
		if out := telegramclient.EscapeMarkdown(test.in); out != test.out {
			t.Errorf("EscapeMarkdown(%q) = %q, want %q", test.in, out, test.out)
		}
	}
}
