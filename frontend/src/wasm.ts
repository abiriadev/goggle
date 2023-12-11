Using Node v20.10.0
import './wasm_exec.js'

const syntaxckPath = './syntaxck.pack.wasm'

export const loadWasm = async (): Promise<void> => {
	const goWasm = new window.Go()
	const result = await WebAssembly.instantiateStreaming(
		fetch(syntaxckPath),
		goWasm.importObject,
	)
	goWasm.run(result.instance)
}
