import type { Config } from "tailwindcss";

export default {
  content: ["./components/**/*.{js,templ}", "./components/*.{js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;
