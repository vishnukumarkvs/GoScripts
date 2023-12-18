/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js,templ,go}"],
  theme: {
    extend: {},
    fontFamily: {
      cb: ["Caveat Brush"],
    },
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
