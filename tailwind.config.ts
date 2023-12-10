import type { Config } from "tailwindcss";

export default {
  content: ["./components/**/*.{ts,templ}", "./components/*.{ts,go,templ}"],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;
