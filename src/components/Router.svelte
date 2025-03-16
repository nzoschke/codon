<script lang="ts">
  export interface Route {
    component: Component<any>;
    hash: string;
    props?: Record<string, any>;
  }

  import NotFound from "~/components/NotFound.svelte";
  import { type Component, mount, onMount, unmount } from "svelte";

  let { routes }: { routes: Route[] } = $props();
  let mounted: Component | undefined;
  let target: HTMLSpanElement;

  const change = (hash: string) => {
    let r = routes.find((r) => r.hash == (hash || "#/")) || {
      component: NotFound,
      hash: "#/404",
      props: {
        hash,
      },
    };

    if (mounted) unmount(mounted);
    mounted = mount<any, any>(r.component, {
      props: r.props,
      target,
    });
  };

  onMount(() => {
    target = document.createElement("span");
    document.body.appendChild(target);

    change(window.location.hash);
  });

  addEventListener("hashchange", () => {
    change(window.location.hash);
  });
</script>
