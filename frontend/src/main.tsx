import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {
	ThemeProvider,
	createTheme,
} from '@mui/material/styles'

const theme = createTheme({
	palette: {
		mode: 'dark',
	},
})

ReactDOM.createRoot(
	document.getElementById('root')!,
).render(
	<React.StrictMode>
		<ThemeProvider theme={theme}>
			<App />
		</ThemeProvider>
	</React.StrictMode>,
)
