import { AppBar, Toolbar, Typography } from '@mui/material'
import GitHubIcon from '@mui/icons-material/GitHub'

export const Bar = () => (
	<AppBar
		position="static"
		enableColorOnDark
		color="primary"
	>
		<Toolbar>
			<Typography flexGrow={1}>News</Typography>
			<GitHubIcon />
		</Toolbar>
	</AppBar>
)
