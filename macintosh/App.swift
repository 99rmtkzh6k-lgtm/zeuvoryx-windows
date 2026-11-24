import SwiftUI
import AppKit

struct ScanFinding: Decodable { let path: String; let kind: String; let sha256: String }
struct ScanReport: Decodable { let files: Int; let bytes: Int64; let skipped: Int; let partial: Bool; let findings: [ScanFinding] }
struct LinkReport: Decodable { let host: String; let signals: [String]; let assessment: String }
struct CheckReport: Decodable { let update: Bool; let statistics: Bool; let policy: Bool }

struct ContentView: View {
    @State private var language = "pt"
    @State private var busy = false
    @State private var report: ScanReport?
    @State private var status = ""
    @State private var checked = false
    @State private var link = ""
    @State private var linkReport: LinkReport?
    func tr(_ pt: String, _ en: String, _ es: String) -> String {
        language == "en" ? en : language == "es" ? es : pt
    }
    var body: some View {
        HStack(alignment: .top, spacing: 22) {
            VStack {
                if let url = Bundle.main.url(forResource: "zeuvoryx_brand", withExtension: "jpg"),
                   let art = NSImage(contentsOf: url) {
                    Image(nsImage: art).resizable().aspectRatio(contentMode: .fit).frame(maxWidth: 290, maxHeight: 450)
                }
                Text("CREATED & ENGINEERED BY").font(.caption)
                Text("Leonardo A Botelho").font(.headline)
            }
            VStack(alignment: .leading, spacing: 14) {
                Text("ZEUVORYX Macintosh").font(.largeTitle.bold())
                Text("3.0.5 LAB BETA").foregroundColor(.orange)
                Picker("Language / Idioma", selection: $language) {
                    Text("Português").tag("pt"); Text("English").tag("en"); Text("Español").tag("es")
                }.pickerStyle(.segmented)
                Text(tr("Scanner experimental somente leitura. Não substitui a proteção do macOS.", "Experimental read-only scanner. Not a replacement for macOS security.", "Escáner experimental de solo lectura. No sustituye la protección de macOS."))
                HStack {
                    Button(tr("Examinar arquivo ou pasta","Scan file or folder","Examinar archivo o carpeta")) { choose() }.disabled(busy)
                    Button(tr("Diagnóstico local","Local diagnostics","Diagnóstico local")) { run(["selftest"]) }.disabled(busy)
                }
                HStack {
                    TextField(tr("URL para análise local", "URL for local assessment", "URL para análisis local"), text: $link)
                    Button(tr("Analisar link", "Assess link", "Analizar enlace")) { run(["link", link]) }
                        .disabled(busy || link.trimmingCharacters(in: .whitespacesAndNewlines).isEmpty || link.utf8.count > 8192)
                }
                if let l = linkReport {
                    Text(tr("Domínio: ", "Host: ", "Dominio: ") + l.host).textSelection(.enabled)
                    Text(tr("Indicadores estruturais: ", "Structural signals: ", "Indicadores estructurales: ") + String(l.signals.count))
                    Text(tr("Análise local, sem acessar o site. Ausência de indicadores não garante segurança.", "Local analysis, no website access. No signals does not guarantee safety.", "Análisis local, sin acceder al sitio. Ausencia de indicadores no garantiza seguridad.")).font(.caption)
                    if language == "pt" { ForEach(Array(l.signals.enumerated()), id: \.offset) { item in Text(item.element).font(.caption) } }
                }
                if busy { ProgressView().controlSize(.small) }
                if checked { Text(tr("Diagnóstico dos componentes aprovado (dados fictícios).","Component diagnostics passed (test data).","Diagnóstico de componentes aprobado (datos ficticios).")) }
                if !status.isEmpty { Text(status).foregroundColor(.orange) }
                if let r = report {
                    Text(tr("Arquivos","Files","Archivos") + ": " + String(r.files) + "  |  " + tr("Ignorados","Skipped","Omitidos") + ": " + String(r.skipped))
                    if r.partial { Text(tr("Análise parcial","Partial scan","Análisis parcial")).foregroundColor(.orange) }
                    Text(tr("Alertas não são confirmação de malware.","Alerts do not confirm malware.","Las alertas no confirman malware."))
                    ScrollView {
                        VStack(alignment: .leading) {
                            ForEach(Array(r.findings.enumerated()), id: \.offset) { item in
                                Text(item.element.kind + " — " + item.element.path).font(.caption).textSelection(.enabled)
                            }
                        }.frame(maxWidth: .infinity, alignment: .leading)
                    }.frame(minHeight: 100)
                }
                Spacer()
                Text(tr("Sem telemetria, atualizações automáticas, remoção ou proteção residente próprias nesta beta.","No telemetry, automatic updates, removal or resident protection in this beta.","Sin telemetría, actualizaciones automáticas, eliminación ni protección residente en esta beta.")).font(.caption)
            }
        }.padding(24).frame(minWidth: 850, minHeight: 540)
        .onAppear {
            let code = Locale.current.language.languageCode?.identifier ?? "en"
            language = ["pt","es"].contains(code) ? code : "en"
        }
    }
    func choose() {
        let p = NSOpenPanel()
        p.canChooseDirectories = true; p.canChooseFiles = true; p.allowsMultipleSelection = false
        if p.runModal() == .OK, let u = p.url { run(["scan", u.path]) }
    }
    func run(_ arguments: [String]) {
        guard !busy else { return }
        guard let exe = Bundle.main.url(forAuxiliaryExecutable: "zeuvoryx-engine") else { status = tr("Motor indisponível.", "Engine unavailable.", "Motor no disponible."); return }
        busy = true; status = ""; checked = false; report = nil; linkReport = nil
        DispatchQueue.global(qos: .utility).async {
            let process = Process(); let pipe = Pipe()
            process.executableURL = exe; process.arguments = arguments; process.standardOutput = pipe
            process.standardError = FileHandle.nullDevice
            do {
                try process.run()
                let data = pipe.fileHandleForReading.readDataToEndOfFile()
                process.waitUntilExit()
                guard process.terminationStatus == 0 else { throw CocoaError(.fileReadUnknown) }
                if arguments.first == "selftest" {
                    let r = try JSONDecoder().decode(CheckReport.self, from: data)
                    DispatchQueue.main.async { checked = r.update && r.statistics && r.policy; busy = false }
                } else if arguments.first == "link" {
                    let r = try JSONDecoder().decode(LinkReport.self, from: data)
                    DispatchQueue.main.async { linkReport = r; busy = false }
                } else {
                    let r = try JSONDecoder().decode(ScanReport.self, from: data)
                    DispatchQueue.main.async { report = r; busy = false }
                }
            } catch {
                DispatchQueue.main.async { status = tr("Operação não concluída.","Operation failed.","Operación no completada."); busy = false }
            }
        }
    }
}
@main struct ZeuvoryxMacintosh: App {
    var body: some Scene { WindowGroup { ContentView() } }
}
