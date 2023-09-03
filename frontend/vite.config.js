import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import { checker } from "vite-plugin-checker";
import tsconfigPaths from "vite-tsconfig-paths";
import svgr from "vite-plugin-svgr";
import handlebars from "vite-plugin-handlebars";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// https://vitejs.dev/config/
export default ({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return defineConfig({
    plugins: [
      react(),
      checker({
        overlay: { initialIsOpen: false },
        typescript: true,
        eslint: {
          lintCommand: 'eslint "./src/**/*.{ts,tsx}"',
        },
      }),
      tsconfigPaths(),
      svgr(),
      handlebars({
        partialDirectory: resolve(__dirname, "src/partials"),
      }),
    ],
    server: {
      port: 3000,
      proxy: {
        "/api": {
          target: "http://localhost:5000",
          changeOrigin: true,
          secure: false,
          ws: true,
        },
      },
    },
  });
};
