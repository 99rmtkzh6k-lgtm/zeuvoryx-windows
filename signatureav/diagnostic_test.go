package signatureav

import (
	"strings"
	"testing"
)

func TestDiagnosticSeparatesLimitsAndRepeatedMessages(t *testing.T) {
	raw := strings.Join([]string{
		"LibClamAV Error: mapping: GetLastError = 32",
		"LibClamAV Error: mapping: GetLastError = 32",
		"WARNING: permission denied",
		"C:\\archive.zip: Heuristics.Limits.Exceeded.MaxFileSize FOUND",
		"C:\\deep.zip: Heuristics.Limits.Exceeded.MaxRecursion FOUND",
		"C:\\control.txt: Eicar-Test-Signature FOUND",
		"----------- SCAN SUMMARY -----------",
		"Known viruses: 100",
		"Scanned files: 7",
		"Infected files: 3",
		"Total errors: 2",
	}, "\n")
	var r Report
	parse(&r, raw, 2, false)
	if !r.Partial || r.Errors != 2 || r.ErrorMessagesByCategory["file_busy"] != 2 || r.ErrorMessagesByCategory["access_denied"] != 1 {
		t.Fatalf("native count or message categories lost: %+v", r)
	}
	if r.LimitNotices != 2 || r.LimitNoticeReasons["max_file_size"] != 1 || r.LimitNoticeReasons["max_recursion"] != 1 {
		t.Fatalf("limits not separated: %+v", r)
	}
	if len(r.Findings) != 3 || r.Findings[2].Kind != "test_eicar" || r.Findings[0].Kind != "heuristic_or_limit" {
		t.Fatalf("findings misclassified: %+v", r.Findings)
	}
	parse(&r, "", 2, true)
	if len(r.Findings) != 0 || r.LimitNoticeReasons != nil || r.ErrorMessagesByCategory != nil {
		t.Fatalf("stale diagnostics after reset: %+v", r)
	}
}
