import { creator, defineAppHooks } from "@/app/internal/creator";
import { autoLoading } from "@/app/middleware/autoLoading";
import { AppSlice, slice } from "@/app/slice";

export const { store } = creator({
  slice,
  middleware: [autoLoading()],
});
export const { dispatch, getState } = store;
export const { useAppDispatch, useAppSelector } = defineAppHooks<AppSlice>();
