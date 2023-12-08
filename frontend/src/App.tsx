import { Typography } from '@mui/material'
import { Bar } from './Bar.tsx'
import { Search } from './Search.tsx'

function App() {
	return (
		<>
			<Bar />
			<main
				style={{
					display: 'flex',
					alignItems: 'center',
					flexDirection: 'column',
				}}
			>
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
