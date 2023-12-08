import { AppBar, Toolbar, Typography } from '@mui/material'

export const Bar = () => (
	<AppBar
		position="static"
		enableColorOnDark
		color="primary"
	>
		<Toolbar>
			<Typography flexGrow={1}>News</Typography>
			<Typography>GitHub</Typography>
		</Toolbar>
	</AppBar>
)
