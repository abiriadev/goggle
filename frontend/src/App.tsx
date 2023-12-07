import { useEffect, useState } from 'react'
import './App.css'
import { loadWasm } from './wasm'
import {
	Autocomplete,
	TextField,
	useTheme,
} from '@mui/material'
import './syntaxck.d.ts'
import { ResultItem, query } from './goggle.ts'

function App() {
	const [wasmLoad, setWasmLoad] = useState(false)
	const [resultSet, setResultSet] = useState<
		Array<ResultItem>
	>([])
	const [inp, setInp] = useState('')
	const theme = useTheme()

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

			try {
				const rs = await query(inp)

				setResultSet(rs.items)
			} catch (e) {
				console.error(e)
			}
		})()
	}, [inp])

	return (
		<main>
			<h1
				style={{
					color: theme.palette.primary.main,
				}}
			>
				Goggle
			</h1>
			<Autocomplete
				sx={{ width: 600 }}
				options={resultSet}
				renderInput={p => <TextField {...p} />}
				onInputChange={(_, i) => setInp(i)}
				getOptionLabel={o =>
					typeof o === 'string' ? o : o.sig
				}
			/>
		</main>
	)
}

export default App
