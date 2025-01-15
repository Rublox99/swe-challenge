/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      maxHeight: {
        'calc-78.5': 'calc(78.5%)',
        'calc-86': 'calc(86%)'
      },
      fontFamily: {
        roboto: [
          "Roboto"
        ]
      },
      colors: {
        primary: {
          100: "#d1dfdf",
          200: "#a3bfbf",
          300: "#749e9e",
          400: "#467e7e",
          500: "#185e5e",
          600: "#134b4b",
          700: "#0e3838",
          800: "#0a2626",
          900: "#051313"
        },
        gray: {
          100: "#f6f6f6",
          200: "#eeeeee",
          300: "#e5e5e5",
          400: "#dddddd",
          500: "#d4d4d4",
          600: "#aaaaaa",
          700: "#7f7f7f",
          800: "#555555",
          900: "#2a2a2a"
        },
        light: {
          contrast: "#202030",
          inv_contrast: "#FDFDFD",
        },
        dark: {
          contrast: "#FDFDFD",
          inv_contrast: "#202030",
        }
      },
      animation: {
        'gradient-bounce': 'gradientBounce 3s ease infinite',
        'fade': 'fadeIn 0.5s ease-out'
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0%' },
          '100%': { opacity: '100%' }
        },
        gradientBounce: {
          '0%': {
            'background-position': '0% 50%',
          },
          '50%': {
            'background-position': '100% 50%',
          },
          '100%': {
            'background-position': '0% 50%',
          },
        }
      },
      backgroundSize: {
        '400%': '400%',
      },
    },
  },
  plugins: [],
}

