import { eifCreateSlice } from "../../my-redux-toolkit/eif.createStore";

export interface LoggedInUser {
  email: string;
}

const loggedInUserSlice = eifCreateSlice('loggedInUser', {}, {
  reducers: {
    set: (_state: LoggedInUser, action: PayloadAction<{ email: string }>) => {
      return { email: action.payload.email };
    },
    reset: (_state: LoggedInUser) => {
      return {};
    }
  }
});

export default loggedInUserSlice;
