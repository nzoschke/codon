import type { BunPlugin } from "bun";
import { compile } from "svelte/compiler";
import { marked } from "marked";

const plugin: BunPlugin = {
  name: "markdown",

  async setup(build) {
    const out = compile(
      `
      <script>
        let content = "<i>content</i>";
      </script>

      <article class="prose">{@html content}</article>`,
      {
        filename: "Markdown.svelte",
        generate: "client",
      },
    );

    build.onLoad({ filter: /\.md/ }, async ({ path }) => {
      let md = await Bun.file(path).text();
      let html = marked.parse(md);

      return {
        contents: out.js.code.replace('"<i>content</i>"', JSON.stringify(html)),
        loader: "js",
      };
    });
  },
};

export default plugin;
