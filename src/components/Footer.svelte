<script lang="ts">
  import { Icon } from "@steeze-ui/svelte-icon";
  import { Cube } from "@steeze-ui/heroicons";
  import { Instagram, Youtube } from "@steeze-ui/simple-icons";
</script>

<footer
  class="footer sm:footer-horizontal bg-neutral text-neutral-content items-center p-4"
>
  <aside class="grid-flow-col items-center">
    <Icon src={Cube} theme="solid" class="h-9" />
    <p>Copyright © {new Date().getFullYear()} - All right reserved</p>
  </aside>
  <nav class="grid-flow-col gap-4 md:place-self-center md:justify-self-end">
    <a href="https://youtube.com">
      <Icon src={Youtube} theme="solid" class="h-6" />
    </a>

    <a href="https://instagram.com/nzoschke">
      <Icon src={Instagram} theme="solid" class="h-6" />
    </a>
  </nav>
</footer>
