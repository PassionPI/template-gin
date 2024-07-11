import { defineAppAsyncThunks } from "./define";
import { task } from "./task";

export const thunks = defineAppAsyncThunks({
  task,
});
