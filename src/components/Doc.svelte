<script lang="ts">
  import Layout from "~/components/Layout.svelte";
  import Deploy from "@/doc/deploy.md";
</script>

<Layout>
  {#snippet header()}
    header
  {/snippet}

  {#snippet asideL()}
    aside left
  {/snippet}

  <Deploy />

  {#snippet asideR()}
    aside right
  {/snippet}

  {#snippet footer()}
    footer
  {/snippet}
</Layout>
