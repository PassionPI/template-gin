import { defineSlices } from "@/app/internal/creator";
import { history } from "@/app/slice/history";
import { loading } from "@/app/slice/loading";

export const slice = defineSlices({
  history,
  loading,
});

export type AppSlice = typeof slice;
