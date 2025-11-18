<script lang="ts">
  import { ToastContainer } from 'svelte-toastify'
  import CheckingCurrentlyLoggedInUser from "./modules/auth/CheckingCurrentlyLoggedInUser.svelte";
  import SimpleLayout from "./modules/layout/SimpleLayout.svelte";
</script>

<CheckingCurrentlyLoggedInUser>
  <SimpleLayout />
</CheckingCurrentlyLoggedInUser>
<ToastContainer />
