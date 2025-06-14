/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				// 🎮 Classic Pokémon Theme
				classic: {
					red: '#CC0000',
					yellow: '#FFCC00',
					white: '#F8F8F8',
					navy: '#003366',
					green: '#3FA34D',
					gray: '#444444'
				},

				// 🔬 Modern Pokédex UI Theme
				modern: {
					blue: '#0075FF',
					navy: '#0A0F2F',
					gray: '#A3A8B0',
					cyan: '#E6F7FF',
					red: '#FF4D4D'
				},

				// 🌿 Nature & Elemental Theme
				elemental: {
					grass: '#78C850',
					water: '#6890F0',
					fire: '#F08030',
					normal: '#A8A878',
					psychic: '#A040A0'
				},

				// 📘 Minimal & Clean Theme
				clean: {
					blue: '#D6EFFF',
					white: '#FFFFFF',
					gray: '#3E4A59',
					coral: '#FF6B6B',
					mint: '#A8E6CF'
				}
			}
		}
	},
	plugins: []
};
