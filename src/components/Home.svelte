<script lang="ts">
  import Layout from "~/src/components/Layout.svelte";
  import Header from "~/src/components/Header.svelte";
  import Footer from "~/src/components/Footer.svelte";
  import { Icon } from "@steeze-ui/svelte-icon";
  import { Cog6Tooth } from "@steeze-ui/heroicons";
</script>

<Layout>
  {#snippet header()}
    <Header />
  {/snippet}

  <!-- {#snippet asideL()}
    <div class="">aside left</div>
  {/snippet} -->

  <div class="text-base-content space-y-4 flex flex-col items-center">
    <div class="hero bg-base-200 min-h-[60vh]">
      <div class="hero-content text-center space-y-6 flex flex-col">
        <div
          class="text-5xl font-bold bg-clip-text text-transparent bg-linear-to-r pb-2 from-purple-200 to-emerald-200"
        >
          Codon
        </div>
        <h1 class="text-5xl font-bold">
          A lean way to dev, build and deploy a web app
        </h1>
        <p class="text-2xl">
          Built with <a class="underline font-bold" href="https://bun.sh">Bun</a
          >,
          <a class="underline font-bold" href="https://go.dev">Go</a>,
          <a class="underline font-bold" href="https://svelte.dev">Svelte</a>,
          <a class="underline font-bold" href="https://sqlite.org">Sqlite</a>,
          <a class="underline font-bold" href="https://tailwindcss.com"
            >Tailwind</a
          >,
          <a class="underline font-bold" href="https://daisyui.com">DaisyUI</a>
        </p>
        <button class="btn btn-primary">Get Started</button>
      </div>
    </div>

    <div class="max-w-lg mx-auto text-center">
      <div
        class="text-3xl md:text-5xl font-bold bg-clip-text text-transparent bg-linear-to-r pb-2 from-purple-200 to-emerald-200"
      >
        Explore the Features
      </div>
      <div class="mt-4 text-xl font-bold">
        And try them on this <span class="underline">fully functional demo</span
        >
      </div>
    </div>

    <div class="flex flex-wrap justify-center gap-4 max-w-4xl">
      {#each { length: 6 }, i}
        <div class="card bg-base-100 w-64 shadow-sm">
          <figure class="pt-6">
            <Icon src={Cog6Tooth} class="h-16" />
          </figure>
          <div class="card-body items-center">
            <h2 class="card-title">Feature {i + 1}</h2>
            <p>
              A card component has a figure, a body part, and inside body there
              are title and actions parts
            </p>
            <div class="card-actions justify-center">
              <button class="btn btn-primary">Learn More</button>
            </div>
          </div>
        </div>
      {/each}
    </div>

    <div class="hero bg-base-200 max-w-5xl min-h-[60vh]">
      <div class="hero-content flex-col lg:flex-row">
        <img
          src="https://img.daisyui.com/images/stock/photo-1635805737707-575885ab0820.webp"
          class="max-w-sm rounded-lg shadow-2xl"
        />
        <div>
          <h1 class="text-5xl font-bold">See it in action</h1>
          <p class="py-6">
            Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
            excepturi exercitationem quasi. In deleniti eaque aut repudiandae et
            a id nisi.
          </p>
          <button class="btn btn-primary">Get Started</button>
        </div>
      </div>
    </div>
  </div>

  <!-- {#snippet asideR()}
    <div class="">aside right</div>
  {/snippet} -->

  {#snippet footer()}
    <Footer />
  {/snippet}
</Layout>
