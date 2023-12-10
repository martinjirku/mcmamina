import type { Config } from "tailwindcss";

export default {
  content: ["./components/**/*.{ts,templ}", "./components/*.{ts,templ}"],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;
