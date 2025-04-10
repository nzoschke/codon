<script lang="ts">
  import { Icon } from "@steeze-ui/svelte-icon";
  import { Instagram, Youtube } from "@steeze-ui/simple-icons";
</script>

<footer
  class="footer footer-horizontal footer-center bg-base-200 text-base-content rounded p-10"
>
  <nav class="grid grid-flow-col gap-4">
    <a class="link link-hover" href="/#about">About us</a>
    <a class="link link-hover" href="/#contact">Contact</a>
  </nav>
  <nav>
    <div class="grid grid-flow-col gap-4">
      <a href="https://youtube.com">
        <Icon src={Youtube} theme="solid" class="h-6" />
      </a>

      <a href="https://instagram.com/nzoschke">
        <Icon src={Instagram} theme="solid" class="h-6" />
      </a>
    </div>
  </nav>
  <aside>
    <p>
      Copyright © {new Date().getFullYear()} - All right reserved by ACME
      Industries Ltd
    </p>
  </aside>
</footer>
