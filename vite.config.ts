import { defineConfig } from 'vite';
import glob from 'fast-glob';

export default defineConfig((config) => {
    return {
        plugins: [],
        build: {
            outDir: 'dist',
            ssr: true,
            target: 'esnext',
            lib: {
                entry: glob.sync(['components/**/*.ts']),
                formats: ['es'],
            },
            sourcemap: true,
            minify: true,
        },
        assetsInclude: ['assets/**/*', 'assets/*'],
    };
});