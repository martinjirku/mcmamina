import { defineConfig, splitVendorChunkPlugin } from "vite";

export default defineConfig((config) => {
  return {
    plugins: [splitVendorChunkPlugin()],
    appType: "custom",
    build: {
      target: "modules",
      sourcemap: true,
      minify: true,
      cssCodeSplit: false,
      manifest: true,
      rollupOptions: {
        input: ["components/pages/Index.ts"],
        output: {
          dir: "dist",
          format: "es",
          entryFileNames: "[name].[format]",
        },
      },
    },
  };
});
