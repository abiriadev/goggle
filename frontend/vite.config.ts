import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => ({
	plugins: [react()],

	base: mode === 'development' ? '/' : '/goggle',

	server: {
		proxy: {
			'/api': {
				target: 'http://localhost:6099',
				changeOrigin: true,
				rewrite: p => p.replace(/^\/api/, ''),
			},
		},
	},

	assetsInclude: ['**/*.wasm'],
}))
