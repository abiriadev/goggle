import { useEffect, useState } from 'react'
import './App.css'
import { loadWasm } from './wasm'

function App() {
	const [wasmLoad, setWasmLoad] = useState(false)

	useEffect(() => {
		;(async () => {
			await loadWasm()
			setWasmLoad(true)
		})()
	}, [])

	return <>hello</>
}

export default App
