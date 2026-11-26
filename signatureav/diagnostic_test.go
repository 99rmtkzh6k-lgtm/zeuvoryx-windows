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
	if !r.NativeSummaryPresent || !r.Partial || r.Errors != 2 || r.ErrorMessagesByCategory["file_busy"] != 2 || r.ErrorMessagesByCategory["access_denied"] != 1 {
		t.Fatalf("native count or message categories lost: %+v", r)
	}
	if r.LimitNotices != 2 || r.LimitNoticeReasons["max_file_size"] != 1 || r.LimitNoticeReasons["max_recursion"] != 1 {
		t.Fatalf("limits not separated: %+v", r)
	}
	if len(r.Findings) != 3 || r.Findings[2].Kind != "test_eicar" || r.Findings[0].Kind != "heuristic_or_limit" {
		t.Fatalf("findings misclassified: %+v", r.Findings)
	}
	parse(&r, "", 2, true)
	if r.NativeSummaryPresent || len(r.Findings) != 0 || r.LimitNoticeReasons != nil || r.ErrorMessagesByCategory != nil {
		t.Fatalf("stale diagnostics after reset: %+v", r)
	}
}

func TestDiagnosticWindowsCodeBoundaries(t *testing.T){
 for _,c:=range []struct{code,want string}{{"32","file_busy"},{"33","file_busy"},{"5","access_denied"},{"320","other_engine_message"},{"50","other_engine_message"}}{if got:=errorMessageCategory("LibClamAV Error: GetLastError = "+c.code);got!=c.want{t.Fatalf("code %s got %s",c.code,got)}}
}
