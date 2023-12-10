/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ['./templates/**/*.html'],
  theme: {
    extend: {
      colors: {
        "earth": "#040c33",
        "earth-text": "#00afdd",
      }
    }
  },
  daisyui: {
    themes: [
      {
        base: {
          "color-scheme": "dark",
          "primary": "#5B75F0",
          "secondary": "#f05b75",
          "accent": "#A6B4F7",
          "neutral": "#1c212b",
          "neutral-content": "#B2CCD6",
          "base-100": "#2A303C",
          "base-200": "#242933",
          "base-300": "#20252E",
          "base-content": "#B2CCD6",
          "info": "#28ebff",
          "success": "#62efbd",
          "warning": "#efd057",
          "error": "#ffae9b",
        },
      },
    ],
  },
  plugins: [
    require('@tailwindcss/typography'),
    require("daisyui"),
  ],
}

