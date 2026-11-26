function buildTestDiagnosis(report, source) {
    const integer = value => Number.isSafeInteger(value) && value >= 0 ? value : null;
    const number = value => Number.isFinite(value) && value >= 0 ? value : null;
    const token = (value, pattern) => typeof value === 'string' && pattern.test(value) ? value : null;
    const timestamp = value => token(value, /^20\d\d-\d\d-\d\dT\d\d:\d\d:\d\d(?:\.\d{1,9})?(?:Z|[+-]\d\d:\d\d)$/);
    const inventory = report.inventory_314 || {};
    const engine = Array.isArray(report.signature_reports) ? report.signature_reports[0] || {} : {};
    const errors = report.error_categories || {};
    const facts = {
        commit: token(report.build_commit, /^[a-f0-9]{7,40}$/),
        versao: token(report.engine_version, /^\d+\.\d+\.\d+(?:-lab)?$/),
        modo: token(report.scan_mode, /^(quick|full|custom)$/),
        inicio: timestamp(report.started),
        registro: timestamp(report.snapshot_at || report.finished),
        estado: report.in_progress === true ? 'em andamento' : report.cancelled === true ? 'cancelado' :
            report.partial === true ? 'final parcial' : timestamp(report.finished) ? 'final' : 'indeterminado',
        duracao_ms: integer(report.duration_ms),
        inventario: {arquivos: integer(inventory.files), bytes: integer(inventory.bytes), falhas: integer(inventory.errors), excluidos: integer(inventory.excluded)},
        scanner: {arquivos: integer(report.files), bytes: integer(report.bytes), ignorados: integer(report.skipped),
            compartilhamento: integer(errors.sharing_violation), bloqueio: integer(errors.lock_violation), acesso_negado: integer(errors.access_denied)},
        clamav: {estado: token(report.signature_state, /^(pending|running|finished|partial|unavailable)$/),
            codigo: integer(engine.exit_code), arquivos: integer(engine.scanned_files), erros: integer(engine.engine_reported_errors),
            limites: integer(engine.engine_limit_notices), lidos_mib: number(engine.engine_data_read_mib), examinados_mib: number(engine.engine_data_scanned_mib),
            duracao_ms: integer(report.signature_duration_ms)},
        amostras_diagnostico: source === 'diag' && Array.isArray(report.samples) ? report.samples.length : null
    };
    const hypotheses = [];
    function add(causa, prioridade, confianca, verificar, alternativas, impacto, teste) {
        hypotheses.push({causa, prioridade, confianca, verificar, alternativas, impacto, teste_de_aceite: teste});
    }
    if (facts.estado === 'em andamento' && facts.clamav.estado === 'running') {
        add('Etapa ClamAV ainda em execução no snapshot', 'alta', 'estado confirmado; causa de demora indeterminada',
            'Comparar horário e fase com o JSON final da mesma sessão e verificar a duração do ClamAV.',
            ['Atualizar o diagnóstico após a execução', 'Se persistir, medir avanço e tempo por etapa em corpus controlado'],
            'Encerrar a etapa por prazo preserva resultados parciais, com cobertura reduzida.',
            'Distinguir execução, prazo excedido e cancelamento em testes de regressão.');
    }
    if (facts.clamav.codigo === 2 || facts.clamav.erros > 0 || facts.clamav.limites > 0) {
        add('Cobertura ClamAV reduzida por erros ou limites registrados', 'alta', 'alta para a limitação; causa por arquivo a confirmar',
            'Conferir o resumo e os códigos no log local, sem compartilhar caminhos pessoais.',
            ['Classificar erros de acesso e limites separadamente', 'Reavaliar arquivos grandes com orçamento de memória, disco e tempo'],
            'Elevar limites pode degradar desempenho em computadores com pouca RAM.',
            'Testar ZIP benigno grande, arquivo bloqueado e resumo parcial sem contabilizá-los como infecção.');
    }
    if (facts.scanner.compartilhamento > 0 || facts.scanner.bloqueio > 0 || facts.scanner.acesso_negado > 0) {
        add('Arquivos ocupados, bloqueados ou sem acesso', 'alta', 'alta para as categorias registradas',
            'Reproduzir com um arquivo benigno aberto por outro processo e verificar a classificação do erro.',
            ['Repetir bloqueios transitórios com prazo curto e cancelável', 'Registrar exclusão de cobertura se o acesso continuar indisponível'],
            'Repetições podem aumentar a duração. Preservar controles de acesso e proteção.',
            'Confirmar que arquivo bloqueado fica visível no relatório e que cancelar interrompe novas tentativas.');
    }
    if (facts.estado === 'cancelado') {
        add('Encerramento por cancelamento', 'alta', 'estado confirmado; origem indeterminada',
            'Conferir a origem do cancelamento no histórico local e o estado de cada motor.',
            ['Registrar motivo e fase do cancelamento', 'Preservar relatório parcial para revisão'],
            'Uma sessão cancelada deixa parte da cobertura indisponível.',
            'Testar cancelamento manual e prazo do ClamAV com estados finais distintos.');
    }
    if (facts.inventario.falhas > 0 || facts.inventario.excluidos > 0) {
        add('Inventário com falhas ou exclusões', 'média', 'alta para os contadores registrados',
            'Confrontar o escopo e a contagem por fase com um conjunto de arquivos conhecido.',
            ['Separar bytes inventariados, lidos e indisponíveis', 'Mostrar o saldo sem atribuir causa quando faltarem dados'],
            'O espaço ocupado do volume usa outro denominador.',
            'Reconciliar bytes e arquivos sem tratar o inventário parcial como disco integral.');
    }
    if (!hypotheses.length) {
        add('Amostra insuficiente para atribuir uma falha', 'a confirmar', 'indeterminada',
            'Comparar a execução com a linha de base, os resultados finais e o CI do commit.',
            ['Coletar evidência da fase que falhou antes de editar código'],
            'Uma correção sem reprodução pode introduzir regressão.',
            'Demonstrar a falha e confirmar a correção em teste controlado.');
    }
    return {esquema: 1, objetivo: 'Investigar e corrigir o ZEUVORYX em LAB', fonte: source, fatos: facts, hipoteses: hypotheses,
        pendencias: ['Confirmar branch e CI do commit no repositório', 'Conferir separadamente falhas simultâneas e dados ausentes'],
        instrucao: 'Leia o repositório, o CI e os memorandos antes de editar. Confirme a causa com evidência, corrija no LAB e teste regressões. Preserve Defender e revise a validação física antes de promover.'};
}
