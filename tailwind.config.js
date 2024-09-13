/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './internal/themes/**/*.templ',
    './internal/themes/**/*_templ.go',
    './internal/templates/**/*.templ',
    './internal/templates/**/*_templ.go',
  ],
  theme: {
    extend: {},
  },
  darkMode: 'selector',
  plugins: [],
};
