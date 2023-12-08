import './App.css'
import { useTheme } from '@mui/material'
import { Bar } from './Bar.tsx'
import { Search } from './Search.tsx'

function App() {
	const theme = useTheme()

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
				<Search />
			</main>
		</>
	)
}

export default App
