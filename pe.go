package engine

import (
	"bytes"
	"debug/pe"
	"path/filepath"
	"strings"
)

// analyzePE is an explainable, conservative static feature detector, not a trained ML model.
// No scanned executable is executed by this analyzer.
func analyzePE(path string, data []byte) *Finding {
	if len(data) < 2 || !bytes.Equal(data[:2], []byte{'M', 'Z'}) {
		return nil
	}
	f, err := pe.NewFile(bytes.NewReader(data))
	if err != nil {
		return nil
	}
	defer f.Close()
	var score int
	var reasons []string
	imports, err := f.ImportedSymbols()
	if err == nil {
		seen := map[string]bool{}
		for _, imp := range imports {
			seen[strings.ToLower(imp)] = true
		}
		has := func(parts ...string) bool {
			for key := range seen {
				for _, p := range parts {
					if strings.Contains(key, p) {
						return true
					}
				}
			}
			return false
		}
		// These APIs have many legitimate uses. Require combinations, not isolated calls.
		if has("virtualallocex", "ntallocatevirtualmemory") && has("writeprocessmemory", "ntwritevirtualmemory") && has("createremotethread", "ntcreatethreadex") {
			score += 45
			reasons = append(reasons, "PE: combinação de APIs de injeção remota (indicador, não prova)")
		}
		if has("isdebuggerpresent") && has("checkremotedebuggerpresent") {
			score += 10
			reasons = append(reasons, "PE: verificações múltiplas contra depuração")
		}
	}
	ext := strings.ToLower(filepath.Ext(path))
	if ext == ".pdf" || ext == ".jpg" || ext == ".png" || ext == ".txt" {
		score += 35
		reasons = append(reasons, "Assinatura PE em arquivo com extensão não executável")
	}
	if score == 0 {
		return nil
	}
	return &Finding{Risk: score, Reasons: reasons}
}
