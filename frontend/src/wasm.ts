import './wasm_exec.js'
import './syntaxck.d.ts'

export const loadWasm = async (): Promise<void> => {
	const goWasm = new window.Go()
	const result = await WebAssembly.instantiateStreaming(
		fetch('/syntaxck.pack.wasm'),
		goWasm.importObject,
	)
	goWasm.run(result.instance)
}
