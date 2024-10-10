/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './internal/themes/**/*.templ',
    './internal/themes/**/*_templ.go',
    './internal/themes/defaultTheme/**/*.templ',
    './internal/themes/defaultTheme/**/*_templ.go',
    './internal/plugins/templates/**/*.templ',
    './internal/plugins/templates/**/*_templ.go',
    './internal/plugins/core/**/*.go',
  ],
  theme: {
    extend: {},
  },
  darkMode: 'selector',
  plugins: [],
};
