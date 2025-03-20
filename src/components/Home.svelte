<script lang="ts">
  import Layout from "~/src/components/Layout.svelte";
</script>

<Layout>
  {#snippet header()}
    header
  {/snippet}

  {#snippet asideL()}
    aside left
  {/snippet}

  <div class="text-primary-content">
    <p class="h-[512px] bg-primary p-4 border-2">Paragraph 1</p>
    <p class="h-[512px] bg-primary p-4">Paragraph 2</p>
    <p class="h-[512px] bg-primary p-4">Paragraph 3</p>
  </div>

  {#snippet asideR()}
    aside right
  {/snippet}

  {#snippet footer()}
    footer
  {/snippet}
</Layout>
