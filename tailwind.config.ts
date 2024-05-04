import type { Config } from "tailwindcss";

export default {
  content: [
    "./template/**/*.{ts,templ}",
    "./template/*.{ts,go,templ}",
    "./templates/*.{ts,go,tmpl}",
    "./templates/**/*.{ts,go,tmpl}",
    "./handlers/*.go"
  ],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;
