import { defineConfig, splitVendorChunkPlugin } from "vite";
import { viteStaticCopy } from "vite-plugin-static-copy";

export default defineConfig({
  plugins: [
    splitVendorChunkPlugin(),
    viteStaticCopy({
      targets: [
        {
          src: "./assets/images/favicon.ico",
          dest: "./",
          overwrite: true,
          dereference: true,
        },
      ],
    }),
  ],
  appType: "custom",
  resolve: {
    alias: {
      "@assets": "/assets",
    },
  },
  build: {
    target: "modules",
    sourcemap: true,
    minify: false,
    terserOptions: {
      format: {
        comments: false,
        beautify: false,
      },
    },
    cssCodeSplit: false,
    manifest: true,
    rollupOptions: {
      input: [
        "template/pages/Index.ts",
        "template/pages/AboutUs.ts",
        "template/pages/SupportedUs.ts",
        "template/pages/Calendar.ts",
        // -> MCMAMINA - GENERATE PAGE
      ],
      output: {
        dir: "dist",
        format: "es",
        entryFileNames: "[name].[format]",
      },
    },
  },
});
