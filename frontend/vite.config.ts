import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [],
	test: {
		environment: 'jsdom',
		exclude: ['**/node_modules/**', '**/e2e/**'],
	},
});
