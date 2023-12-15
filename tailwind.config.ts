import type { Config } from "tailwindcss";

export default {
  content: ["./template/**/*.{ts,templ}", "./template/*.{ts,go,templ}"],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;
