import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import eslintPLugin from "vite-plugin-eslint";
import svgLoader from "vite-svg-loader";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 8080,
    host: "127.0.0.1",
  },
  preview: {
    port: 8080,
    host: "127.0.0.1",
  },
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
      "~bootstrap": resolve(__dirname, "node_modules/bootstrap"),
      "~nouislider": resolve(__dirname, "node_modules/nouislider"),
      "~tom-select": resolve(__dirname, "node_modules/tom-select"),
      "~jsvectormap": resolve(__dirname, "node_modules/jsvectormap"),
      //
    },
  },
  plugins: [vue(), svgLoader(), eslintPLugin()],
});
