import markdown from "~/src/plugins/markdown";
import svelte from "bun-plugin-svelte";

await Bun.build({
  entrypoints: ["./src/index.html"],
  naming: {
    entry: "[name].[ext]",
  },
  outdir: "./build/dist",
  plugins: [markdown, svelte],
  root: ".",
});
