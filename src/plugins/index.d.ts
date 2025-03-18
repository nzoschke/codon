interface Markdown {
  attrs: Record<string, any> & {
    order: number;
    slug: string;
    title: string;
  };

  html: string;
}

declare module "*.md" {
  const md: Markdown;
  export default md;
}
