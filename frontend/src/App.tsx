import { useEffect, useState } from 'react'
import './App.css'
import { loadWasm } from './wasm'
import { Autocomplete, TextField } from '@mui/material'

function App() {
	const [wasmLoad, setWasmLoad] = useState(false)
	const [resultSet, setResultSet] = useState([])

	useEffect(() => {
		;(async () => {
			await loadWasm()
			setWasmLoad(true)
		})()
	}, [])

	return (
		<main>
			<h1>Goggle</h1>
			<Autocomplete
				options={resultSet}
				renderInput={p => <TextField {...p} />}
			/>
		</main>
	)
}

export default App
