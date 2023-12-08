import { useEffect, useState } from 'react'
import './App.css'
import { loadWasm } from './wasm'
import {
	Autocomplete,
	Box,
	Grid,
	TextField,
	Typography,
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
				freeSolo
				sx={{ width: 600 }}
				options={resultSet}
				renderInput={p => <TextField {...p} />}
				renderOption={(p, o) => (
					<li {...p}>
						<a href={o.link} target="_blank">
							{o.sig}
							<Typography
								paragraph
								noWrap
								variant="caption"
								color="#bbbbbb"
								sx={{
									fontWeight: 300,
								}}
							>
								{o.summary}
							</Typography>
						</a>
					</li>
				)}
				onInputChange={(_, i) => setInp(i)}
				getOptionLabel={o =>
					typeof o === 'string' ? o : o.sig
				}
				onChange={(_, v, r) => {
					if (r === 'selectOption') {
						// r.
					}
					console.log('value:', v, 'reason:', r)
				}}
			/>
		</main>
	)
}

export default App
