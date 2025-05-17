<script lang="ts">
  import Layout from "~/src/components/Layout.svelte";
  import Deploy from "~/doc/deploy.md";
  import Dev from "~/doc/dev.md";
  import Header from "~/src/components/Header.svelte";
  import Footer from "~/src/components/Footer.svelte";

  const sections: Markdown[] = [Deploy, Dev].sort(
    (a, b) => a.attrs.order - b.attrs.order,
  );

  const slug = new URLSearchParams(window.location.search).get("slug") ||
    sections[0]!.attrs.slug;
  const section = sections.find((s) => s.attrs.slug == slug) || Deploy;
</script>

<Layout>
  {#snippet header()}
    <Header />
  {/snippet}

  {#snippet asideL()}
    <ul>
      {#each sections as s}
        <li>
          <a href={`?slug=${s.attrs.slug}${window.location.hash}`}>{
            s.attrs.title
          }</a>
        </li>
      {/each}
    </ul>
  {/snippet}

  <article class="prose w-full">{@html section.html}</article>

  {#snippet asideR()}
    aside right
  {/snippet}

  {#snippet footer()}
    <Footer />
  {/snippet}
</Layout>
