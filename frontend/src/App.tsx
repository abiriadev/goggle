import { useEffect, useState } from 'react'
import './App.css'
import { loadWasm } from './wasm'
import { Autocomplete, TextField } from '@mui/material'
import './syntaxck.d.ts'
import { ResultItem, query } from './goggle.ts'

function App() {
	const [wasmLoad, setWasmLoad] = useState(false)
	const [resultSet, setResultSet] = useState<
		Array<ResultItem>
	>([])
	const [inp, setInp] = useState('')

	useEffect(() => {
		;(async () => {
			await loadWasm()
			setWasmLoad(true)
		})()
	}, [])

	useEffect(() => {
		if (!(wasmLoad && window.syntaxck(inp))) return
		;(async () => {
			console.log('input change:', inp)

			const rs = await query(inp)

			setResultSet(rs.items)
		})()
	}, [inp])

	return (
		<main>
			<h1>Goggle</h1>
			<Autocomplete
				options={resultSet}
				renderInput={p => <TextField {...p} />}
				onInputChange={(_, i) => setInp(i)}
			/>
		</main>
	)
}

export default App
