import { defineSliceItem } from "@/app/internal/creator";

export const loading = defineSliceItem({
  initialState: (): Record<string, boolean> => ({}),
});
