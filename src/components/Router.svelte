<script lang="ts">
  import { type Component, mount, unmount } from "svelte";
  import NotFound from "~/components/NotFound.svelte";

  export interface Route {
    component: Component<any>;
    hash: string;
    props?: Record<string, any>;
  }

  let { routes }: { routes: Route[] } = $props();
  let mounted: Component | undefined;

  let target = document.createElement("span");
  document.body.appendChild(target);

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

  addEventListener("hashchange", () => {
    change(window.location.hash);
  });

  change(window.location.hash);
</script>
