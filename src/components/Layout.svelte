<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    asideL?: Snippet;
    asideR?: Snippet;
    header?: Snippet;
    children?: Snippet;
    footer?: Snippet;
  }

  let { asideL, asideR, children, header, footer }: Props = $props();

  let cols = [asideL && "auto", "1fr", asideR && "auto"]
    .filter(Boolean)
    .join("_");
</script>

<div>
  {#if header}
    <header
      class="sticky top-0 z-10 bg-base-300/80 text-base-content backdrop-blur-sm"
    >
      {@render header()}
    </header>
  {/if}

  <div class={`grid grid-cols-1 md:grid-cols-[${cols}]`}>
    {#if asideL}
      <aside class="bg-base-200 text-base-content p-4">
        {@render asideL()}
      </aside>
    {/if}

    <main class="bg-base-100 p-4 space-y-4">
      {@render children?.()}
    </main>

    {#if asideR}
      <aside class="bg-base-200 text-base-content p-4">
        {@render asideR()}
      </aside>
    {/if}
  </div>

  {#if footer}
    <footer class="bg-base-300 text-base-content p-4">
      {@render footer()}
    </footer>
  {/if}
</div>
