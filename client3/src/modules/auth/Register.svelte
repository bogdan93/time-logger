<script lang="ts">
  import { eifActions } from "../../store";
  import Input from "../common/Input.svelte";
  import DataFieldValidator from "../DataFieldValidator.svelte";
  import LayoutPageBlockButton from "../layout/LayoutPageBlockButton.svelte";
  import { goToLogin } from "./auth.actions";

  interface RegisterForm {
    email?: string;
    password?: string;
  }
  export let block: EifLayoutBlock;
  export let data: RegisterForm | undefined;
  export let dataPath: string;
  export let errors: Record<string, string> = {};

  function setEmail(email: string) {
    eifActions.setData(`${dataPath}.email`, email);
  }
  function setPassword(password: string) {
    eifActions.setData(`${dataPath}.password`, password);
  }
</script>

<form autocomplete="off" on:submit={e => e.preventDefault()}>
  <div class="max-w-xl flex flex-col gap-2 mx-auto">
    <h1 class="text-4xl">Register</h1>
    <DataFieldValidator
      dataPath={`${dataPath}.email`}
      validationRules={{ required: "Please input your Email" }}
    >
      <Input
        name="email"
        placeholder="john-snow@example.com"
        label="Email"
        value={data?.email || ""}
        error={errors?.email}
        autocomplete="off"
        on:change={(e) => setEmail(e.detail.value)}
      />
    </DataFieldValidator>

    <DataFieldValidator
      dataPath={`${dataPath}.password`}
      validationRules={{ required: "Please input your Password!" }}
    >
      <Input
        name="password"
        type="password"
        placeholder="type your password"
        label="Password"
        value={data?.password || ""}
        autocomplete="off"
        error={errors?.password}
        on:change={(e) => setPassword(e.detail.value)}
      />
    </DataFieldValidator>

    <button
      class="text-blue-500 hover:underline"
      on:click={goToLogin}>
      Go to login
    </button>

    <LayoutPageBlockButton
      className="self-end bg-white p-2 rounded-md w-20 text-black hover:text-blue-500 inline-flex justify-center"
      {block}
      buttonAction={block.actions[0]}
    />
  </div>
</form>
