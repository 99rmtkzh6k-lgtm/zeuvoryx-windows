# ZEUVORYX Antivirus — Centro de ayuda

Base pública de soporte para ZEUVORYX Windows 3.0.6 LAB R1, candidato experimental de seguridad complementaria. No es una versión de producción, una certificación independiente ni un sustituto de Microsoft Defender.

[Português](pt-br.md) · [English](en.md) · [Español](es.md) · [日本語](ja.md)

## ¿Qué es ZEUVORYX?

Escáner de seguridad para Windows x64 desarrollado en Go, con análisis de archivos, indicadores SHA-256, heurística, cuarentena reversible e IA consultiva experimental. Macintosh 3.0.5 LAB es un candidato independiente en validación.

## Instalación y actualización

Utilice solo paquetes de prueba autorizados y compruebe la versión y SHA-256 con la referencia de GitHub. Extraiga el ZIP, termine o cancele el análisis, cierre la aplicación, haga una copia de seguridad de la configuración y cuarentena e instale sobre una edición LAB compatible. Mantenga Defender y SmartScreen activos. El CI comprobó la migración de 3.0.5 LAB a 3.0.6 LAB; otras migraciones no están verificadas.

## La aplicación no abre o se cierra

No reinstale repetidamente. Compruebe la versión, Windows x64, Microsoft Edge WebView2 Runtime, espacio y memoria disponibles. Consulte el Monitor de confiabilidad (perfmon /rel) y el Visor de eventos (eventvwr.msc), registros Aplicación y Sistema cerca del fallo. No borre la cuarentena ni desactive la seguridad de Windows. Comparta solo información anonimizada.

## El análisis completo se bloquea o termina

Los cierres en equipos físicos siguen en investigación. En CI, la interfaz instalada 3.0.6 permaneció abierta 15 segundos tras confirmar su inicio y pasaron 60 pruebas de un conjunto pequeño controlado. Esto no demuestra estabilidad en todo el disco. Si vuelve a ocurrir, detenga las pruebas en su equipo de trabajo y anote hora, versión, modo, CPU/RAM y último progreso. No ejecute archivos sospechosos.

## Errores, archivos omitidos y alertas

El acceso denegado, archivos bloqueados o desaparecidos pueden impedir la lectura. Un informe parcial o cero detecciones no demuestra ausencia de malware. Una heurística o puntuación de IA es un indicio, no una confirmación. No elimine archivos del sistema solo por una alerta.

## IA, tiempo real, ClamAV y actualizaciones

La IA actual es consultiva, sin aprendizaje local continuo validado. Protección residente propia, aprendizaje continuo, actualizaciones automáticas de producción y estadísticas globales siguen en desarrollo. ClamAV es opcional y requiere instalación separada. La autoprueba no demuestra un servicio mundial operativo. No hay versión de producción anunciada.

## Cómo solicitar ayuda

Abra una incidencia en GitHub con versión, compilación de Windows 10/11, arquitectura, RAM aproximada, modo de análisis, hora y zona horaria, resultado esperado y observado, pasos de reproducción y categoría de error o código de salida. Quite nombres, rutas personales, correos, tokens, identificadores, contenido y hashes de archivos privados. No publique malware, registros sin depurar, cuarentena ni credenciales. No se garantiza atención las 24 horas.

---

- [GitHub](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows) · [Issues / support](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/issues)
- [Windows 3.0.6 LAB R1 CI evidence](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45479946405)
- [Historical memorandum](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/main/MEMORANDO_DESENVOLVIMENTO.md)
- Updated: 2026-11-24.