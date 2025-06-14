/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				theme: {
					base: 'var(--color-base)',
					primary: 'var(--color-primary)',
					secondary: 'var(--color-secondary)',
					accent: 'var(--color-accent)',
					neutral: 'var(--color-neutral)',
					info: 'var(--color-info)',
					success: 'var(--color-success)',
					warning: 'var(--color-warning)',
					error: 'var(--color-error)'
				}
			}
		}
	},
	plugins: []
};
