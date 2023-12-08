import { Container, Stack, Typography } from '@mui/material'
import { Bar } from './Bar.tsx'
import { Search } from './Search.tsx'

function App() {
	return (
		<>
			<Bar />
			<Stack
				direction="column"
				alignItems="center"
				spacing={3}
				marginTop={20}
			>
				<Typography
					variant="h1"
					color="primary.main"
					fontFamily="Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;"
					fontWeight="bold"
					fontSize="3.2em"
				>
					Goggle
				</Typography>
				<Search />
			</Stack>
		</>
	)
}

export default App
