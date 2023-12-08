import { Typography, useTheme } from '@mui/material'
import { Bar } from './Bar.tsx'
import { Search } from './Search.tsx'

// #root {
// 	max-width: 1280px;
// 	margin: 0 auto;
// 	padding: 2rem;
// 	text-align: center;
// }

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
