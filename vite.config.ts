import { defineConfig, splitVendorChunkPlugin } from "vite";

export default defineConfig({
  plugins: [splitVendorChunkPlugin()],
  appType: "custom",
  resolve: {
    alias: {
      "@assets": "/assets",
    },
  },
  build: {
    target: "modules",
    sourcemap: true,
    minify: true,
    cssCodeSplit: false,
    manifest: true,
    rollupOptions: {
      input: ["template/pages/Index.ts"],
      output: {
        dir: "dist",
        format: "es",
        entryFileNames: "[name].[format]",
      },
    },
  },
});
