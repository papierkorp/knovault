/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./internal/templates/**/*.templ', './internal/templates/**/*_templ.go'],
  theme: {
    extend: {},
  },
  darkMode: 'selector',
  plugins: [],
};
