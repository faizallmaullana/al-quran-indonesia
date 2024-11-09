module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false,
  theme: {
    extend: {
      backgroundImage: {
        pattern: "url('@/components/imgs/vector.svg')",
      },
    },
    variants: {
      extend: {},
    },
    plugins: [],
  },
}
