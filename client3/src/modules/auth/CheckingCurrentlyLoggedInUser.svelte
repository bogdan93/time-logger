<script lang="ts">
  import GearSolid from "svelte-awesome-icons/GearSolid.svelte";
  import api from "../../api";
  import { appInit } from "../app.init";
  import { goToLogin } from "./auth.actions";
  import loggedInUserSlice from "./loggedInUser.slice";

  async function checkIfLoggedIn() {
    if (api.authStore.model && api.authStore.isValid) {
      try {
        loggedInUserSlice.actions.set({ email: api.authStore.model.email });
        await appInit();
        return;
      } catch {
        goToLogin();
      }
    }

    goToLogin();
  }
</script>

{#await checkIfLoggedIn()}
  <div class="container mx-auto min-h-screen flex items-center justify-center">
    <div class="text-white animate-ping">
      <GearSolid class="animate-spin" />
    </div>
  </div>
{:then}
  <slot />
{/await}
