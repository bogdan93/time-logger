import api from "../../api";
import { eifActions } from "../../store";
import { appDestroy, appInit } from "../app.init";
import Login from "./Login.svelte";
import Register from "./Register.svelte";

let counterId = 0;
export function goToLogin() {
  appDestroy();

  eifActions.replaceLayoutBlock({
    type: "page",
    title: "Login",
    dataPath: `loginForm${++counterId}`,
    component: Login,
    actions: [
      {
        title: 'login',
        action: async (block: EifLayoutBlock) => {
          const data = eifActions.getDataAtPath<{ email: string, password: string }>(block.dataPath);

          try {
            await api.users.authViaEmail(data.email, data.password);

            await appInit();

          } catch (error: any) {
            if (error?.data?.data?.password?.message) {
              eifActions.setError(
                `${block.dataPath}.password`,
                error?.data?.data?.password?.message
              );
            } else if (error?.data?.data?.email?.message) {
              eifActions.setError(
                `${block.dataPath}.email`,
                error?.data?.data?.email?.message
              );
            } else {
              eifActions.setError(
                `${block.dataPath}.password`,
                'Something went wrong. Please try again',
              );
            }
          }
        }
      },
    ],
  }, 'deleteDataAtPath');
}

export function goToRegister() {
  appDestroy();

  eifActions.replaceLayoutBlock({
    type: "page",
    title: "Register",
    dataPath: `registerForm${++counterId}`,
    component: Register,
    actions: [
      {
        title: 'register',
        action: async (block: EifLayoutBlock) => {
          const data = eifActions.getDataAtPath<{ email: string, password: string }>(block.dataPath);

          try {
            await api.users.create({
              email: data.email, password: data.password, passwordConfirm: data.password
            });
            await api.users.authViaEmail(data.email, data.password);

            await appInit();

          } catch (error: any) {
            if (error?.data?.data?.password?.message) {
              eifActions.setError(
                `${block.dataPath}.password`,
                error?.data?.data?.password?.message
              );
            } else if (error?.data?.data?.email?.message) {
              eifActions.setError(
                `${block.dataPath}.email`,
                error?.data?.data?.email?.message
              );
            } else {
              eifActions.setError(
                `${block.dataPath}.password`,
                'Something went wrong. Please try again',
              );
            }
          }
        }
      },
    ],
  }, 'deleteDataAtPath');
}
