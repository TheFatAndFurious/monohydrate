/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*.html", "./templates/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
