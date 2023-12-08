import {
	AppBar,
	Link,
	Toolbar,
	Typography,
} from '@mui/material'
import GitHubIcon from '@mui/icons-material/GitHub'

export const Bar = () => (
	<AppBar
		position="static"
		enableColorOnDark
		color="primary"
	>
		<Toolbar variant="dense">
			<Typography flexGrow={1}>Goggle</Typography>
			<Link
				href="https://github.com/abiriadev/goggle"
				color="inherit"
				underline="none"
			>
				<GitHubIcon />
			</Link>
		</Toolbar>
	</AppBar>
)
