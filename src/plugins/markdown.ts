import type { BunPlugin } from "bun";
import fm from "front-matter";
import { marked } from "marked";

const plugin: BunPlugin = {
  name: "markdown",

  async setup(build) {
    build.onLoad({ filter: /\.md/ }, async ({ path }) => {
      const src = await Bun.file(path).text();
      const res = fm(src);
      const html = marked.parse(res.body, {
        gfm: true,
      });

      return {
        contents: `export default {
          attrs: ${JSON.stringify(res.attributes)},
          html: ${JSON.stringify(html)},
        }`,
        loader: "js",
      };
    });
  },
};

export default plugin;
