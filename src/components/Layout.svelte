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
</script>

<div>
  {#if header}
    <header
      class="sticky top-0 z-10 bg-base-300/80 text-base-content backdrop-blur-sm"
    >
      {@render header()}
    </header>
  {/if}

  <div class="flex flex-grow justify-end overflow-scroll flex-col lg:flex-row">
    {#if asideL}
      <aside class="bg-base-200 text-base-content p-4">
        {@render asideL()}
      </aside>
    {/if}

    <main class="bg-base-100 p-4 space-y-4 w-full min-h-[calc(100vh-8rem)]">
      {@render children?.()}
    </main>

    {#if asideR}
      <aside class="bg-base-200 text-base-content p-4">
        {@render asideR()}
      </aside>
    {/if}
  </div>

  {#if footer}
    <footer class="bg-base-300 text-base-content">
      {@render footer()}
    </footer>
  {/if}
</div>
