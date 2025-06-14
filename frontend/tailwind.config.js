/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				// 🎮 Classic Pokémon Theme
				classic: {
					primary: '#CC0000', // Pokéball red
					secondary: '#FFCC00', // Pikachu yellow
					accent: '#3FA34D', // Grass green
					neutral: '#444444', // UI text
					base: '#F8F8F8', // Background
					info: '#003366', // Navy info blocks
					success: '#3FA34D',
					warning: '#FFCC00',
					error: '#CC0000'
				},

				// 🔬 Modern Pokédex UI
				modern: {
					primary: '#0075FF', // Electric blue
					secondary: '#A3A8B0', // Cool gray
					accent: '#FF4D4D', // Alert red
					neutral: '#0A0F2F', // Dark background
					base: '#E6F7FF', // Light cyan bg
					info: '#0075FF',
					success: '#3FA34D',
					warning: '#FFB020',
					error: '#FF4D4D'
				},

				// 🌿 Elemental Theme
				elemental: {
					primary: '#78C850', // Grass
					secondary: '#6890F0', // Water
					accent: '#F08030', // Fire
					neutral: '#A8A878', // Normal type
					base: '#FFFFFF',
					info: '#A040A0', // Psychic
					success: '#78C850',
					warning: '#F08030',
					error: '#A040A0'
				},

				// 📘 Minimal Clean Theme
				clean: {
					primary: '#3E4A59', // Steel gray text
					secondary: '#D6EFFF', // Soft blue bg
					accent: '#FF6B6B', // Coral action
					neutral: '#FFFFFF',
					base: '#FFFFFF',
					info: '#3E4A59',
					success: '#A8E6CF',
					warning: '#FFCC00',
					error: '#FF6B6B'
				}
			}
		}
	},
	plugins: []
};
