const assert = require('node:assert/strict');
const fs = require('node:fs');
const vm = require('node:vm');

const context = vm.createContext({});
vm.runInContext(fs.readFileSync('lab/test_diagnosis.js', 'utf8'), context);
const rawPath = 'C:\\Users\\PrivateName\\private.zip';
const result = context.buildTestDiagnosis({
    scan_mode: 'full', engine_version: '3.0.15-lab', cancelled: true,
    error_categories: {sharing_violation: 7},
    inventory_314: {files: 10, bytes: 2048, errors: 1},
    signature_reports: [{exit_code: 2, scanned_files: -1, engine_limit_notices: 9, engine_output: rawPath}],
    root: rawPath, errors: [rawPath]
}, 'current');
assert.equal(result.fatos.estado, 'cancelado');
assert.equal(result.fatos.clamav.arquivos, null);
assert.equal(result.fatos.clamav.limites, 9);
assert.equal(result.hipoteses.length, 4);
assert.equal(JSON.stringify(result).includes('PrivateName'), false);
assert.equal(context.buildTestDiagnosis({in_progress: true, signature_state: 'running'}, 'current').fatos.estado, 'em andamento');
assert.equal(context.buildTestDiagnosis({samples: [{}]}, 'diag').fatos.amostras_diagnostico, 1);
if (process.argv[2]) {
    const html = fs.readFileSync(process.argv[2], 'utf8');
    for (const id of ['z315BuildDiagnosis', 'z315CopyDiagnosis', 'z315SaveDiagnosis']) {
        assert.equal(html.split('id="' + id + '"').length, 2, id);
    }
}
