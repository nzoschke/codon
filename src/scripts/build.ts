import svelte from "bun-plugin-svelte";
import tailwind from "bun-plugin-tailwind";

await Bun.build({
  entrypoints: ["./src/index.html"],
  naming: {
    entry: "[name].[ext]",
  },
  outdir: "./build/dist",
  plugins: [svelte, tailwind],
  root: ".",
});
