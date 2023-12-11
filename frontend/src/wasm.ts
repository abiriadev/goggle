import './wasm_exec.js'
import syntaxckPath from './assets/syntaxck.pack.wasm'

export const loadWasm = async (): Promise<void> => {
	const goWasm = new window.Go()
	const result = await WebAssembly.instantiateStreaming(
		fetch(syntaxckPath),
		goWasm.importObject,
	)
	goWasm.run(result.instance)
}
