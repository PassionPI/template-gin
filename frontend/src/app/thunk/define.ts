import { definer } from "@/app/internal/creator";
import { AppSlice } from "@/app/slice";

export const {
  defineAppSelector,
  defineAppThunk,
  defineAppMiddleware,
  defineAppAsyncThunkCollection,
  defineAppAsyncThunks,
} = definer<AppSlice>();
