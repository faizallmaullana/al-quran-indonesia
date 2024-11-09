module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false,
  theme: {
    extend: {
      fontFamily: {
        rubik: ['Rubik', 'sans serif'],
      },
      colors: {
        myorange: '#243c5a',
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
