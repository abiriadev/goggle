import { AppBar, Toolbar, Typography } from '@mui/material'

export const Bar = () => (
	<AppBar
		position="static"
		enableColorOnDark
		color="primary"
	>
		<Toolbar>
			<Typography>News</Typography>
		</Toolbar>
	</AppBar>
)
