/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pkg/templates/**/*.html"],
  theme: {
    extend: {},
  },
  variants: {
    extend: {
      margin: ["first", "last"],
    },
  },
  plugins: [],
}
