const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    colors: {
      indigo: colors.indigo,
    },
  },
  plugins: [],
};
