import type { BunPlugin } from "bun";
import { parse } from "yaml";
import { marked } from "marked";

const plugin: BunPlugin = {
  name: "markdown",

  async setup(build) {
    build.onLoad({ filter: /\.md/ }, async ({ path }) => {
      let md = await Bun.file(path).text();
      let yaml = "";

      const matches = md.match(/^---\n(.*?)\n---\n/ms);
      if (matches) {
        yaml = matches[1]!;
        md = md.replace(matches[0], "");
      }

      const attrs = parse(yaml);
      // console.log(attrs[0].)
      const html = marked.parse(md, {
        gfm: true,
      });

      return {
        contents: `export default {
          attrs: ${JSON.stringify(attrs)},
          html: ${JSON.stringify(html)},
        }`,
        loader: "js",
      };
    });
  },
};

export default plugin;
