import { thunks } from "@/app/thunk";
import { defineAppAsyncThunkCollection } from "./define";

export const task = defineAppAsyncThunkCollection({
  async task1(payload: string, { getState }) {
    const l = getState().loading;
    console.log("????", l, payload);
  },
  async task2(payload: string, { dispatch }) {
    dispatch(thunks.task.task1(payload + "1"));
  },
});
