import { Typography } from '@mui/material'
import { Bar } from './Bar.tsx'
import { Search } from './Search.tsx'

// #root {
// 	max-width: 1280px;
// 	margin: 0 auto;
// 	padding: 2rem;
// 	text-align: center;
// }

function App() {
	return (
		<>
			<Bar />
			<main>
				<Typography
					variant="h1"
					sx={{
						color: 'primary.main',
						fontFamily:
							'Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;',
						fontWeight: 'bold',
						fontSize: '3.2em',
					}}
				>
					Goggle
				</Typography>
				<Search />
			</main>
		</>
	)
}

export default App
