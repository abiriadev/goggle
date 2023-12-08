import { useEffect, useState } from 'react'
import './App.css'
import { loadWasm } from './wasm'
import {
	Autocomplete,
	Link,
	TextField,
	Typography,
	useTheme,
} from '@mui/material'
import './syntaxck.d.ts'
import { ResultItem, query } from './goggle.ts'
import { Bar } from './Bar.tsx'

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
		<>
			<Bar />
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
					sx={{
						width: 600,
					}}
					options={resultSet}
					filterOptions={_ => _}
					getOptionLabel={o =>
						typeof o === 'string' ? o : o.sig
					}
					renderInput={p => (
						<TextField
							sx={{
								'& .MuiOutlinedInput-root':
									{
										'& > fieldset': {
											borderColor:
												theme
													.palette
													.primary
													.main,
										},
									},
								'& .MuiOutlinedInput-root:hover':
									{
										'& > fieldset': {
											borderColor:
												theme
													.palette
													.primary
													.main,
										},
									},
							}}
							{...p}
						/>
					)}
					renderOption={(p, o) => (
						<li {...p}>
							<Link
								href={o.link}
								underline="none"
								target="_blank"
							>
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
							</Link>
						</li>
					)}
					onInputChange={(_, i) => setInp(i)}
					onChange={(_, o, r) => {
						if (r === 'selectOption') {
							window.open(
								typeof o === 'string'
									? o
									: o?.link,
								'_blank',
							)
						}
						console.log(
							'value:',
							o,
							'reason:',
							r,
						)
					}}
				/>
			</main>
		</>
	)
}

export default App
