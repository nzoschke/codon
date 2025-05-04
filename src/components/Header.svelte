<script lang="ts">
  import { Icon } from "@steeze-ui/svelte-icon";
  import {
    Bars3,
    BookOpen,
    Document,
    Home,
    Photo,
    User,
  } from "@steeze-ui/heroicons";
  import { Github } from "@steeze-ui/simple-icons";
</script>

<div class="navbar shadow-sm">
  <div class="navbar-start">
    <div class="dropdown">
      <div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
        <Icon src={Bars3} class="h-5" />
      </div>
      <ul class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
        <li>
          <a href="#/">
            <Icon src={Home} theme="solid" class="h-5" />
            Home
          </a>
        </li>
        <li>
          <a href="#/contacts">
            <Icon src={User} theme="solid" class="h-5" />
            Contacts
          </a>
        </li>
        <li>
          <a href="#/gallery">
            <Icon src={Photo} theme="solid" class="h-5" />
            Gallery
          </a>
        </li>
        <li>
          <a href="#/doc">
            <Icon src={BookOpen} theme="solid" class="h-5" />
            Docs
          </a>
        </li>
        <li>
          <a href="http://localhost:1234/spec">
            <Icon src={Document} theme="solid" class="h-5" />
            Spec
          </a>
        </li>
      </ul>
    </div>
    <a class="btn btn-ghost text-xl" href="/#">Codon</a>
  </div>
  <div class="navbar-center hidden lg:flex">
    <ul class="menu menu-horizontal px-1">
      <li>
        <a href="#/">
          <Icon src={Home} theme="solid" class="h-5" />
          Home
        </a>
      </li>
      <li>
        <a href="#/contacts">
          <Icon src={User} theme="solid" class="h-5" />
          Contacts
        </a>
      </li>
      <li>
        <a href="#/gallery">
          <Icon src={Photo} theme="solid" class="h-5" />
          Gallery
        </a>
      </li>
      <li>
        <a href="#/doc">
          <Icon src={BookOpen} theme="solid" class="h-5" />
          Docs
        </a>
      </li>
      <li>
        <a href="http://localhost:1234/spec">
          <Icon src={Document} theme="solid" class="h-5" />
          Spec
        </a>
      </li>
    </ul>
  </div>
  <div class="navbar-end">
    <a class="btn" href="https://github.com/nzoschke/codon">
      <Icon src={Github} theme="solid" class="h-5" />
      GitHub
    </a>
  </div>
</div>
